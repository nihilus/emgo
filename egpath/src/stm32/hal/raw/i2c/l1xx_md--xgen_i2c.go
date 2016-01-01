// +build l1xx_md

package i2c

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type I2C_Periph struct {
	CR1   CR1
	_     uint16
	CR2   CR2
	_     uint16
	OAR1  OAR1
	_     uint16
	OAR2  OAR2
	_     uint16
	DR    DR
	_     uint16
	SR1   SR1
	_     uint16
	SR2   SR2
	_     uint16
	CCR   CCR
	_     uint16
	TRISE TRISE
}

func (p *I2C_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var I2C1 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C1_BASE)))

var I2C2 = (*I2C_Periph)(unsafe.Pointer(uintptr(mmap.I2C2_BASE)))

type CR1_Bits uint16

func (b CR1_Bits) Field(mask CR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR1_Bits) J(v int) CR1_Bits {
	return CR1_Bits(bits.Make32(v, uint32(mask)))
}

type CR1 struct{ mmio.U16 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U16.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U16.Store(uint16(b)) }

type CR1_Mask struct{ mmio.UM16 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM16.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) PE() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(PE)}}
}

func (p *I2C_Periph) SMBUS() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(SMBUS)}}
}

func (p *I2C_Periph) SMBTYPE() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(SMBTYPE)}}
}

func (p *I2C_Periph) ENARP() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(ENARP)}}
}

func (p *I2C_Periph) ENPEC() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(ENPEC)}}
}

func (p *I2C_Periph) ENGC() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(ENGC)}}
}

func (p *I2C_Periph) NOSTRETCH() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(NOSTRETCH)}}
}

func (p *I2C_Periph) START() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(START)}}
}

func (p *I2C_Periph) STOP() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(STOP)}}
}

func (p *I2C_Periph) ACK() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(ACK)}}
}

func (p *I2C_Periph) POS() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(POS)}}
}

func (p *I2C_Periph) PEC() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(PEC)}}
}

func (p *I2C_Periph) ALERT() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(ALERT)}}
}

func (p *I2C_Periph) SWRST() CR1_Mask {
	return CR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint16(SWRST)}}
}

type CR2_Bits uint16

func (b CR2_Bits) Field(mask CR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CR2_Bits) J(v int) CR2_Bits {
	return CR2_Bits(bits.Make32(v, uint32(mask)))
}

type CR2 struct{ mmio.U16 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U16.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U16.Store(uint16(b)) }

type CR2_Mask struct{ mmio.UM16 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM16.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) FREQ() CR2_Mask {
	return CR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(FREQ)}}
}

func (p *I2C_Periph) ITERREN() CR2_Mask {
	return CR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(ITERREN)}}
}

func (p *I2C_Periph) ITEVTEN() CR2_Mask {
	return CR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(ITEVTEN)}}
}

func (p *I2C_Periph) ITBUFEN() CR2_Mask {
	return CR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(ITBUFEN)}}
}

func (p *I2C_Periph) DMAEN() CR2_Mask {
	return CR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(DMAEN)}}
}

func (p *I2C_Periph) LAST() CR2_Mask {
	return CR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint16(LAST)}}
}

type OAR1_Bits uint16

func (b OAR1_Bits) Field(mask OAR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OAR1_Bits) J(v int) OAR1_Bits {
	return OAR1_Bits(bits.Make32(v, uint32(mask)))
}

type OAR1 struct{ mmio.U16 }

func (r *OAR1) Bits(mask OAR1_Bits) OAR1_Bits { return OAR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *OAR1) StoreBits(mask, b OAR1_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *OAR1) SetBits(mask OAR1_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *OAR1) ClearBits(mask OAR1_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *OAR1) Load() OAR1_Bits               { return OAR1_Bits(r.U16.Load()) }
func (r *OAR1) Store(b OAR1_Bits)             { r.U16.Store(uint16(b)) }

type OAR1_Mask struct{ mmio.UM16 }

func (rm OAR1_Mask) Load() OAR1_Bits   { return OAR1_Bits(rm.UM16.Load()) }
func (rm OAR1_Mask) Store(b OAR1_Bits) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) ADD1_7() OAR1_Mask {
	return OAR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint16(ADD1_7)}}
}

func (p *I2C_Periph) ADD8_9() OAR1_Mask {
	return OAR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint16(ADD8_9)}}
}

func (p *I2C_Periph) ADD0() OAR1_Mask {
	return OAR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint16(ADD0)}}
}

func (p *I2C_Periph) ADDMODE() OAR1_Mask {
	return OAR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint16(ADDMODE)}}
}

type OAR2_Bits uint16

func (b OAR2_Bits) Field(mask OAR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask OAR2_Bits) J(v int) OAR2_Bits {
	return OAR2_Bits(bits.Make32(v, uint32(mask)))
}

type OAR2 struct{ mmio.U16 }

func (r *OAR2) Bits(mask OAR2_Bits) OAR2_Bits { return OAR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *OAR2) StoreBits(mask, b OAR2_Bits)   { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *OAR2) SetBits(mask OAR2_Bits)        { r.U16.SetBits(uint16(mask)) }
func (r *OAR2) ClearBits(mask OAR2_Bits)      { r.U16.ClearBits(uint16(mask)) }
func (r *OAR2) Load() OAR2_Bits               { return OAR2_Bits(r.U16.Load()) }
func (r *OAR2) Store(b OAR2_Bits)             { r.U16.Store(uint16(b)) }

type OAR2_Mask struct{ mmio.UM16 }

func (rm OAR2_Mask) Load() OAR2_Bits   { return OAR2_Bits(rm.UM16.Load()) }
func (rm OAR2_Mask) Store(b OAR2_Bits) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) ENDUAL() OAR2_Mask {
	return OAR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint16(ENDUAL)}}
}

func (p *I2C_Periph) ADD2() OAR2_Mask {
	return OAR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint16(ADD2)}}
}

type DR_Bits uint16

func (b DR_Bits) Field(mask DR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask DR_Bits) J(v int) DR_Bits {
	return DR_Bits(bits.Make32(v, uint32(mask)))
}

type DR struct{ mmio.U16 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U16.Bits(uint16(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U16.SetBits(uint16(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U16.ClearBits(uint16(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U16.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U16.Store(uint16(b)) }

type DR_Mask struct{ mmio.UM16 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM16.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM16.Store(uint16(b)) }

type SR1_Bits uint16

func (b SR1_Bits) Field(mask SR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR1_Bits) J(v int) SR1_Bits {
	return SR1_Bits(bits.Make32(v, uint32(mask)))
}

type SR1 struct{ mmio.U16 }

func (r *SR1) Bits(mask SR1_Bits) SR1_Bits { return SR1_Bits(r.U16.Bits(uint16(mask))) }
func (r *SR1) StoreBits(mask, b SR1_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *SR1) SetBits(mask SR1_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *SR1) ClearBits(mask SR1_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *SR1) Load() SR1_Bits              { return SR1_Bits(r.U16.Load()) }
func (r *SR1) Store(b SR1_Bits)            { r.U16.Store(uint16(b)) }

type SR1_Mask struct{ mmio.UM16 }

func (rm SR1_Mask) Load() SR1_Bits   { return SR1_Bits(rm.UM16.Load()) }
func (rm SR1_Mask) Store(b SR1_Bits) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) SB() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(SB)}}
}

func (p *I2C_Periph) ADDR() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(ADDR)}}
}

func (p *I2C_Periph) BTF() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(BTF)}}
}

func (p *I2C_Periph) ADD10() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(ADD10)}}
}

func (p *I2C_Periph) STOPF() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(STOPF)}}
}

func (p *I2C_Periph) RXNE() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(RXNE)}}
}

func (p *I2C_Periph) TXE() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(TXE)}}
}

func (p *I2C_Periph) BERR() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(BERR)}}
}

func (p *I2C_Periph) ARLO() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(ARLO)}}
}

func (p *I2C_Periph) AF() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(AF)}}
}

func (p *I2C_Periph) OVR() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(OVR)}}
}

func (p *I2C_Periph) PECERR() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(PECERR)}}
}

func (p *I2C_Periph) TIMEOUT() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(TIMEOUT)}}
}

func (p *I2C_Periph) SMBALERT() SR1_Mask {
	return SR1_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint16(SMBALERT)}}
}

type SR2_Bits uint16

func (b SR2_Bits) Field(mask SR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask SR2_Bits) J(v int) SR2_Bits {
	return SR2_Bits(bits.Make32(v, uint32(mask)))
}

type SR2 struct{ mmio.U16 }

func (r *SR2) Bits(mask SR2_Bits) SR2_Bits { return SR2_Bits(r.U16.Bits(uint16(mask))) }
func (r *SR2) StoreBits(mask, b SR2_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *SR2) SetBits(mask SR2_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *SR2) ClearBits(mask SR2_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *SR2) Load() SR2_Bits              { return SR2_Bits(r.U16.Load()) }
func (r *SR2) Store(b SR2_Bits)            { r.U16.Store(uint16(b)) }

type SR2_Mask struct{ mmio.UM16 }

func (rm SR2_Mask) Load() SR2_Bits   { return SR2_Bits(rm.UM16.Load()) }
func (rm SR2_Mask) Store(b SR2_Bits) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) MSL() SR2_Mask {
	return SR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint16(MSL)}}
}

func (p *I2C_Periph) BUSY() SR2_Mask {
	return SR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint16(BUSY)}}
}

func (p *I2C_Periph) TRA() SR2_Mask {
	return SR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint16(TRA)}}
}

func (p *I2C_Periph) GENCALL() SR2_Mask {
	return SR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint16(GENCALL)}}
}

func (p *I2C_Periph) SMBDEFAULT() SR2_Mask {
	return SR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint16(SMBDEFAULT)}}
}

func (p *I2C_Periph) SMBHOST() SR2_Mask {
	return SR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint16(SMBHOST)}}
}

func (p *I2C_Periph) DUALF() SR2_Mask {
	return SR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint16(DUALF)}}
}

func (p *I2C_Periph) PEC() SR2_Mask {
	return SR2_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint16(PEC)}}
}

type CCR_Bits uint16

func (b CCR_Bits) Field(mask CCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CCR_Bits) J(v int) CCR_Bits {
	return CCR_Bits(bits.Make32(v, uint32(mask)))
}

type CCR struct{ mmio.U16 }

func (r *CCR) Bits(mask CCR_Bits) CCR_Bits { return CCR_Bits(r.U16.Bits(uint16(mask))) }
func (r *CCR) StoreBits(mask, b CCR_Bits)  { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *CCR) SetBits(mask CCR_Bits)       { r.U16.SetBits(uint16(mask)) }
func (r *CCR) ClearBits(mask CCR_Bits)     { r.U16.ClearBits(uint16(mask)) }
func (r *CCR) Load() CCR_Bits              { return CCR_Bits(r.U16.Load()) }
func (r *CCR) Store(b CCR_Bits)            { r.U16.Store(uint16(b)) }

type CCR_Mask struct{ mmio.UM16 }

func (rm CCR_Mask) Load() CCR_Bits   { return CCR_Bits(rm.UM16.Load()) }
func (rm CCR_Mask) Store(b CCR_Bits) { rm.UM16.Store(uint16(b)) }

func (p *I2C_Periph) CCR() CCR_Mask {
	return CCR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint16(CCR)}}
}

func (p *I2C_Periph) DUTY() CCR_Mask {
	return CCR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint16(DUTY)}}
}

func (p *I2C_Periph) FS() CCR_Mask {
	return CCR_Mask{mmio.UM16{(*mmio.U16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint16(FS)}}
}

type TRISE_Bits uint16

func (b TRISE_Bits) Field(mask TRISE_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TRISE_Bits) J(v int) TRISE_Bits {
	return TRISE_Bits(bits.Make32(v, uint32(mask)))
}

type TRISE struct{ mmio.U16 }

func (r *TRISE) Bits(mask TRISE_Bits) TRISE_Bits { return TRISE_Bits(r.U16.Bits(uint16(mask))) }
func (r *TRISE) StoreBits(mask, b TRISE_Bits)    { r.U16.StoreBits(uint16(mask), uint16(b)) }
func (r *TRISE) SetBits(mask TRISE_Bits)         { r.U16.SetBits(uint16(mask)) }
func (r *TRISE) ClearBits(mask TRISE_Bits)       { r.U16.ClearBits(uint16(mask)) }
func (r *TRISE) Load() TRISE_Bits                { return TRISE_Bits(r.U16.Load()) }
func (r *TRISE) Store(b TRISE_Bits)              { r.U16.Store(uint16(b)) }

type TRISE_Mask struct{ mmio.UM16 }

func (rm TRISE_Mask) Load() TRISE_Bits   { return TRISE_Bits(rm.UM16.Load()) }
func (rm TRISE_Mask) Store(b TRISE_Bits) { rm.UM16.Store(uint16(b)) }
