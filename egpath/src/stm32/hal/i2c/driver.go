package i2c

import (
	"rtos"
	"sync"
	"unsafe"

	"arch/cortexm"

	"stm32/hal/dma"

	"stm32/hal/raw/i2c"
)

// Driver implements interrupt driven driver to I2C peripheral. To use the
// Driver the Periph field must be set and the I2CISR method must be setup as
// I2C interrupt handler for both (event and error) Periph's IRQs.
//
// Setting the RxDMA or/and TxDMA fields enables using DMA for Rx or/and Tx data
// transfer. If DMA is enabled for some direction the DMAISR method must be
// setup as DMA interrupt handler for this direction.
type Driver struct {
	*Periph
	RxDMA *dma.Channel
	TxDMA *dma.Channel

	mutex sync.Mutex
	done  rtos.EventFlag
	buf   []byte
	n     int
	addr  int16
	stop  bool
	state state
}

// NewDriver provides convenient way to create heap allocated Driver struct.
func NewDriver(p *Periph, rxdma, txdma *dma.Channel) *Driver {
	d := new(Driver)
	d.Periph = p
	d.RxDMA = rxdma
	d.TxDMA = txdma
	return d
}

// MasterConn returns initialized MasterConn struct that can be used to
// communicate with the slave device. Addr is the I2C address of the slave.
func (d *Driver) MasterConn(addr int16, stopm StopMode) MasterConn {
	return MasterConn{d: d, addr: addr << 1, stopm: stopm}
	// TODO: Add support for 10-bit addr.
}

// NewMasterConn is like MasterConn but returns pointer to heap allocated
// MasterConn struct.
func (d *Driver) NewMasterConn(addr int16, stopm StopMode) *MasterConn {
	mc := new(MasterConn)
	*mc = d.MasterConn(addr, stopm)
	return mc
}

// Unlock must be used after recovering from error.
func (d *Driver) Unlock() {
	d.mutex.Unlock()
}

// I2CISR is I2C (event and error) interrupt handler.
func (d *Driver) I2CISR() {
	sr1 := d.Periph.raw.SR1.Load()
	if e := getError(sr1); e != 0 {
		d.handleErrors()
		return
	}
	eventHandlersDMA[d.state](d, sr1)
}

type state byte

const (
	stateIdle state = iota
	stateStart
	stateAddr

	stateWriteDMA
	stateWriteWait
	stateWrite

	stateReadDMA
	stateReadWait
	stateRead
	stateRead1
	stateRead2
	stateReadN

	stateError
	stateBelatedStop
	stateBadEvent
	stateTimeout
)

//emgo:const
var eventHandlersDMA = [...]func(d *Driver, sr1 i2c.SR1_Bits){
	stateIdle:  (*Driver).idleEH,
	stateStart: (*Driver).startISR,
	stateAddr:  (*Driver).addrISR,

	stateWriteDMA:  nil,
	stateWrite:     (*Driver).writeISR,
	stateWriteWait: (*Driver).writeWaitEH,

	stateReadDMA:  nil,
	stateReadWait: (*Driver).readWaitEH,
	stateRead:     (*Driver).readISR,
	stateRead1:    (*Driver).read1ISR,
	stateRead2:    (*Driver).read2ISR,
	stateReadN:    (*Driver).readNISR,
}

func (d *Driver) idleEH(sr1 i2c.SR1_Bits) {
	p := &d.Periph.raw
	if d.addr < 0 {
		// Slave - not implemented.
	} else {
		// Master
		bits := i2c.START
		if d.addr&1 != 0 {
			bits |= i2c.ACK
		}
		p.CR1.SetBits(bits)
		d.state = stateStart
	}
	if sr1&i2c.BTF != 0 {
		// Repeated start (most likely).
		// Eensure that BTF was cleared before enable interrupts.
		maxrep := (d.Periph.Bus().Clock() + 16) / 32 // Timeout: 1/32 s.
		for {
			if getError(sr1) != 0 {
				d.handleErrors()
				return
			}
			sr1 = p.SR1.Load()
			if sr1&i2c.BTF == 0 {
				break
			}
			if maxrep == 0 {
				d.state = stateTimeout
				d.handleErrors()
				return
			}
			maxrep--
		}
	}
	d.enableIntI2C(i2c.ITEVTEN | i2c.ITERREN)
}

func (d *Driver) startISR(sr1 i2c.SR1_Bits) {
	if sr1&i2c.SB == 0 {
		d.badEvent(sr1)
		return
	}
	d.state = stateAddr
	p := &d.Periph.raw
	p.DR.Store(i2c.DR_Bits(d.addr))
	p.SR1.Load() // Ensure that SB was cleared before return.
}

func (d *Driver) addrISR(sr1 i2c.SR1_Bits) {
	if sr1&i2c.ADDR == 0 {
		d.badEvent(sr1)
		return
	}
	p := &d.Periph.raw
	if d.addr&1 == 0 {
		// Write.
		p.SR2.Load()
		d.write()
		return
	}
	// Read
	if d.RxDMA != nil && len(d.buf) > 1 {
		d.state = stateReadDMA
		p.SR2.Load()
		d.setupDMA(d.RxDMA, dma.PTM|dma.IncM|dma.FIFO_1_4)
		d.startDMA(d.RxDMA)
		return
	}
	if !d.stop {
		d.state = stateRead
		p.SR2.Load()
		return
	}
	switch len(d.buf) {
	case 1:
		d.state = stateRead1
		p.ACK().Clear()
		cortexm.SetPRIMASK()
		p.SR2.Load()
		p.STOP().Set()
		cortexm.ClearPRIMASK()
		d.enableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
	case 2:
		d.state = stateRead2
		p.POS().Set()
		cortexm.SetPRIMASK()
		p.SR2.Load()
		p.ACK().Clear()
		cortexm.ClearPRIMASK()
	default:
		d.state = stateReadN
		p.SR2.Load()
	}
}

func (d *Driver) readISR(sr1 i2c.SR1_Bits) {
	if sr1&i2c.BTF == 0 {
		d.badEvent(sr1)
		return
	}
	d.buf[d.n] = byte(d.Periph.raw.DR.Load())
	d.n++
	if d.n == len(d.buf) {
		d.disableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
		d.state = stateReadWait
		d.done.Set()
	}
}

func (d *Driver) read1ISR(sr1 i2c.SR1_Bits) {
	if sr1&i2c.RXNE == 0 {
		d.badEvent(sr1)
		return
	}
	d.disableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
	n := d.n
	d.buf[n] = byte(d.Periph.raw.DR.Load())
	d.n = n + 1
	d.state = stateIdle
	d.done.Set()
}

func (d *Driver) read2ISR(sr1 i2c.SR1_Bits) {
	if sr1&i2c.BTF == 0 {
		d.badEvent(sr1)
		return
	}
	p := &d.Periph.raw
	cortexm.SetPRIMASK()
	p.STOP().Set()
	dr := p.DR.Load()
	cortexm.ClearPRIMASK()
	p.POS().Clear()
	d.disableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
	d.buf[0] = byte(dr)
	d.buf[1] = byte(p.DR.Load())
	d.n = 2
	d.state = stateIdle
	d.done.Set()
}

func (d *Driver) readNISR(sr1 i2c.SR1_Bits) {
	if sr1&i2c.BTF == 0 {
		d.badEvent(sr1)
		return
	}
	p := &d.Periph.raw
	n := d.n
	m := len(d.buf) - n
	if m < 3 {
		d.disableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
		d.state = stateBelatedStop
		d.done.Set()
		return
	}
	var dr i2c.DR_Bits
	if m == 3 {
		p.ACK().Clear()
		cortexm.SetPRIMASK()
		dr = p.DR.Load()
		p.STOP().Set()
		cortexm.ClearPRIMASK()
		d.buf[n] = byte(dr)
		n++
		dr = p.DR.Load()
		d.state = stateRead1
		// ITBUFEN must be set after second DR read.
		d.enableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
	} else {
		dr = p.DR.Load()
	}
	d.buf[n] = byte(dr)
	d.n = n + 1
}

func (d *Driver) readWaitEH(_ i2c.SR1_Bits) {
	if d.RxDMA != nil && len(d.buf) > 1 {
		d.state = stateReadDMA
		if d.stop {
			d.Periph.raw.LAST().Set()
		}
		d.setupDMA(d.RxDMA, dma.PTM|dma.IncM|dma.FIFO_1_4)
		d.startDMA(d.RxDMA)
		return
	}
	if d.stop {
		d.state = stateReadN
	} else {
		d.state = stateRead
	}
	d.enableIntI2C(i2c.ITEVTEN | i2c.ITERREN)
}

func (d *Driver) writeISR(sr1 i2c.SR1_Bits) {
	if sr1&i2c.BTF == 0 {
		d.badEvent(sr1)
		return
	}
	n := d.n
	if n == len(d.buf) {
		d.disableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
		d.state = stateWriteWait
		d.done.Set()
		return
	}
	d.Periph.raw.DR.Store(i2c.DR_Bits(d.buf[n]))
	d.n = n + 1
}

func (d *Driver) writeWaitEH(_ i2c.SR1_Bits) {
	d.write()
	if d.state == stateWrite {
		d.enableIntI2C(i2c.ITEVTEN | i2c.ITERREN)
	}
}

func (d *Driver) write() {
	if d.TxDMA != nil && len(d.buf) > 1 {
		d.state = stateWriteDMA
		d.n = 0
		d.setupDMA(d.TxDMA, dma.MTP|dma.IncM|dma.FIFO_4_4)
		d.startDMA(d.TxDMA)
	} else {
		d.state = stateWrite
		d.n = 1
		d.Periph.raw.DR.Store(i2c.DR_Bits(d.buf[0]))
	}
}

func (d *Driver) setupDMA(ch *dma.Channel, mode dma.Mode) {
	d.disableDMA(ch)
	ch.ClearEvents(dma.EV | dma.ERR)
	ch.Setup(mode)
	ch.SetWordSize(1, 1)
	ch.SetAddrP(unsafe.Pointer(d.Periph.raw.DR.U16.Addr()))
}

func (d *Driver) startDMA(ch *dma.Channel) {
	d.disableIntI2C(i2c.ITEVTEN | i2c.ITBUFEN)
	n := d.n
	m := (len(d.buf) - n) & 0xffff
	if len(d.buf)-(n+m) == 1 {
		m-- // Avoid last transfer size 1.
	}
	d.n = n + m
	dmabits := i2c.DMAEN
	if d.stop && d.n == len(d.buf) {
		dmabits |= i2c.LAST
	}
	d.Periph.raw.CR2.SetBits(dmabits)
	ch.SetAddrM(unsafe.Pointer(&d.buf[n]))
	ch.SetLen(m)
	ch.Enable()
	ch.EnableInt(dma.TRCE | dmaErrMask)
}

func (d *Driver) disableDMA(ch *dma.Channel) {
	ch.Disable()
	ch.DisableInt(dma.EV | dma.ERR)
	d.Periph.raw.CR2.ClearBits(i2c.DMAEN | i2c.LAST)
}

func (d *Driver) enableIntI2C(m i2c.CR2_Bits) {
	d.Periph.raw.CR2.SetBits(m)
}

func (d *Driver) disableIntI2C(m i2c.CR2_Bits) {
	d.Periph.raw.CR2.ClearBits(m)
}

func (d *Driver) handleErrors() {
	d.disableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
	if d.RxDMA != nil {
		d.disableDMA(d.RxDMA)
	}
	if d.TxDMA != nil {
		d.disableDMA(d.TxDMA)
	}
	if d.state < stateError {
		d.state = stateError
	}
	d.done.Set()
}

func (d *Driver) badEvent(sr1 i2c.SR1_Bits) {
	d.state = stateBadEvent
	d.handleErrors()
}

func (d *Driver) DMAISR(ch *dma.Channel) {
	d.disableDMA(ch)
	tx := d.addr&1 == 0
	if tx && d.state != stateWriteDMA || !tx && d.state != stateReadDMA {
		d.badEvent(0)
		return
	}
	if ch.Events()&dmaErrMask != 0 {
		d.n -= ch.Len()
		d.handleErrors()
		return
	}
	ch.ClearEvents(dma.EV | dma.ERR)
	if len(d.buf)-d.n != 0 {
		d.startDMA(ch)
		return
	}
	if tx {
		d.state = stateWrite
		d.enableIntI2C(i2c.ITEVTEN | i2c.ITERREN)
		return
	}
	if d.stop {
		d.state = stateIdle
	} else {
		d.state = stateReadWait
	}
	d.done.Set()
}

func (d *Driver) waitDone(ch *dma.Channel) (e Error) {
	timeout := byteTimeout + 2*9e9*int64(len(d.buf)+1)/int64(d.Speed())
	if d.done.Wait(rtos.Nanosec() + timeout) {
		d.done.Clear()
	} else {
		e = SoftTimeout
		d.disableIntI2C(i2c.ITEVTEN | i2c.ITERREN | i2c.ITBUFEN)
		if ch != nil {
			d.disableDMA(ch)
		}
		if d.state < stateError {
			d.state = stateError
		}
	}
	p := &d.Periph.raw
	if d.state >= stateError {
		switch d.state {
		case stateBelatedStop:
			e |= BelatedStop
		case stateBadEvent:
			e |= BadEvent
		case stateTimeout:
			e |= SoftTimeout
		}
		e |= getError(p.SR1.Load())
		if e&Timeout == 0 {
			p.STOP().Set()
		}
		p.SR1.Store(0) // Clear error flags.
		if ch != nil {
			if ch.Events()&dmaErrMask != 0 {
				e |= DMAErr
			}
			ch.ClearEvents(dma.EV | dma.ERR)
		}
		d.state = stateIdle
	}
	return
}

/*
//emgo:const
var stateTXT = [...]string{
	stateIdle:  " stateIdle\n",
	stateStart: " stateStart\n",
	stateAddr:  " stateAddr\n",

	stateWriteDMA:  " stateWriteDMA\n",
	stateWrite:     " stateWrite\n",
	stateWriteWait: " stateWriteWait\n",

	stateRead1:    " stateRead1\n",
	stateReadDMA:  " stateReadDMA\n",
	stateReadWait: " stateReadWait\n",

	stateError:       " stateError\n",
	stateBelatedStop: " stateBelatedStop\n",
	stateBadEvent:    " stateBadEvent\n",
}

func (d *Driver) printState(sr1 i2c.SR1_Bits) {
	strconv.WriteUint32(dbg, uint32(sr1), 16, 4)
	if int(d.state) >= len(stateTXT) {
		dbg.WriteString("unknown\n")
	}
	dbg.WriteString(stateTXT[d.state])
}
*/
