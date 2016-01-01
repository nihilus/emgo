package can

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/f10x_hd/mmap"
)

type CAN_Periph struct {
	MCR   MCR
	MSR   MSR
	TSR   TSR
	RF0R  RF0R
	RF1R  RF1R
	IER   IER
	ESR   ESR
	BTR   BTR
	_     [100]uint32
	FMR   FMR
	FM1R  FM1R
	_     uint32
	FS1R  FS1R
	_     uint32
	FFA1R FFA1R
	_     uint32
	FA1R  FA1R
}

func (p *CAN_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var CAN1 = (*CAN_Periph)(unsafe.Pointer(uintptr(mmap.CAN1_BASE)))

var CAN2 = (*CAN_Periph)(unsafe.Pointer(uintptr(mmap.CAN2_BASE)))

type MCR_Bits uint32

func (b MCR_Bits) Field(mask MCR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MCR_Bits) J(v int) MCR_Bits {
	return MCR_Bits(bits.Make32(v, uint32(mask)))
}

type MCR struct{ mmio.U32 }

func (r *MCR) Bits(mask MCR_Bits) MCR_Bits { return MCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *MCR) StoreBits(mask, b MCR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MCR) SetBits(mask MCR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *MCR) ClearBits(mask MCR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *MCR) Load() MCR_Bits              { return MCR_Bits(r.U32.Load()) }
func (r *MCR) Store(b MCR_Bits)            { r.U32.Store(uint32(b)) }

type MCR_Mask struct{ mmio.UM32 }

func (rm MCR_Mask) Load() MCR_Bits   { return MCR_Bits(rm.UM32.Load()) }
func (rm MCR_Mask) Store(b MCR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) INRQ() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(INRQ)}}
}

func (p *CAN_Periph) SLEEP() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(SLEEP)}}
}

func (p *CAN_Periph) TXFP() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(TXFP)}}
}

func (p *CAN_Periph) RFLM() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(RFLM)}}
}

func (p *CAN_Periph) NART() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(NART)}}
}

func (p *CAN_Periph) AWUM() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(AWUM)}}
}

func (p *CAN_Periph) ABOM() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(ABOM)}}
}

func (p *CAN_Periph) TTCM() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(TTCM)}}
}

func (p *CAN_Periph) RESET() MCR_Mask {
	return MCR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(RESET)}}
}

type MSR_Bits uint32

func (b MSR_Bits) Field(mask MSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask MSR_Bits) J(v int) MSR_Bits {
	return MSR_Bits(bits.Make32(v, uint32(mask)))
}

type MSR struct{ mmio.U32 }

func (r *MSR) Bits(mask MSR_Bits) MSR_Bits { return MSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *MSR) StoreBits(mask, b MSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *MSR) SetBits(mask MSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *MSR) ClearBits(mask MSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *MSR) Load() MSR_Bits              { return MSR_Bits(r.U32.Load()) }
func (r *MSR) Store(b MSR_Bits)            { r.U32.Store(uint32(b)) }

type MSR_Mask struct{ mmio.UM32 }

func (rm MSR_Mask) Load() MSR_Bits   { return MSR_Bits(rm.UM32.Load()) }
func (rm MSR_Mask) Store(b MSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) INAK() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(INAK)}}
}

func (p *CAN_Periph) SLAK() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SLAK)}}
}

func (p *CAN_Periph) ERRI() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(ERRI)}}
}

func (p *CAN_Periph) WKUI() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(WKUI)}}
}

func (p *CAN_Periph) SLAKI() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SLAKI)}}
}

func (p *CAN_Periph) TXM() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(TXM)}}
}

func (p *CAN_Periph) RXM() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(RXM)}}
}

func (p *CAN_Periph) SAMP() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SAMP)}}
}

func (p *CAN_Periph) RX() MSR_Mask {
	return MSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(RX)}}
}

type TSR_Bits uint32

func (b TSR_Bits) Field(mask TSR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask TSR_Bits) J(v int) TSR_Bits {
	return TSR_Bits(bits.Make32(v, uint32(mask)))
}

type TSR struct{ mmio.U32 }

func (r *TSR) Bits(mask TSR_Bits) TSR_Bits { return TSR_Bits(r.U32.Bits(uint32(mask))) }
func (r *TSR) StoreBits(mask, b TSR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *TSR) SetBits(mask TSR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *TSR) ClearBits(mask TSR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *TSR) Load() TSR_Bits              { return TSR_Bits(r.U32.Load()) }
func (r *TSR) Store(b TSR_Bits)            { r.U32.Store(uint32(b)) }

type TSR_Mask struct{ mmio.UM32 }

func (rm TSR_Mask) Load() TSR_Bits   { return TSR_Bits(rm.UM32.Load()) }
func (rm TSR_Mask) Store(b TSR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) RQCP0() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(RQCP0)}}
}

func (p *CAN_Periph) TXOK0() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TXOK0)}}
}

func (p *CAN_Periph) ALST0() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ALST0)}}
}

func (p *CAN_Periph) TERR0() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TERR0)}}
}

func (p *CAN_Periph) ABRQ0() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ABRQ0)}}
}

func (p *CAN_Periph) RQCP1() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(RQCP1)}}
}

func (p *CAN_Periph) TXOK1() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TXOK1)}}
}

func (p *CAN_Periph) ALST1() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ALST1)}}
}

func (p *CAN_Periph) TERR1() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TERR1)}}
}

func (p *CAN_Periph) ABRQ1() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ABRQ1)}}
}

func (p *CAN_Periph) RQCP2() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(RQCP2)}}
}

func (p *CAN_Periph) TXOK2() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TXOK2)}}
}

func (p *CAN_Periph) ALST2() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ALST2)}}
}

func (p *CAN_Periph) TERR2() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TERR2)}}
}

func (p *CAN_Periph) ABRQ2() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(ABRQ2)}}
}

func (p *CAN_Periph) CODE() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CODE)}}
}

func (p *CAN_Periph) TME() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(TME)}}
}

func (p *CAN_Periph) LOW() TSR_Mask {
	return TSR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(LOW)}}
}

type RF0R_Bits uint32

func (b RF0R_Bits) Field(mask RF0R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RF0R_Bits) J(v int) RF0R_Bits {
	return RF0R_Bits(bits.Make32(v, uint32(mask)))
}

type RF0R struct{ mmio.U32 }

func (r *RF0R) Bits(mask RF0R_Bits) RF0R_Bits { return RF0R_Bits(r.U32.Bits(uint32(mask))) }
func (r *RF0R) StoreBits(mask, b RF0R_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RF0R) SetBits(mask RF0R_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RF0R) ClearBits(mask RF0R_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RF0R) Load() RF0R_Bits               { return RF0R_Bits(r.U32.Load()) }
func (r *RF0R) Store(b RF0R_Bits)             { r.U32.Store(uint32(b)) }

type RF0R_Mask struct{ mmio.UM32 }

func (rm RF0R_Mask) Load() RF0R_Bits   { return RF0R_Bits(rm.UM32.Load()) }
func (rm RF0R_Mask) Store(b RF0R_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FMP0() RF0R_Mask {
	return RF0R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(FMP0)}}
}

func (p *CAN_Periph) FULL0() RF0R_Mask {
	return RF0R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(FULL0)}}
}

func (p *CAN_Periph) FOVR0() RF0R_Mask {
	return RF0R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(FOVR0)}}
}

func (p *CAN_Periph) RFOM0() RF0R_Mask {
	return RF0R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(RFOM0)}}
}

type RF1R_Bits uint32

func (b RF1R_Bits) Field(mask RF1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask RF1R_Bits) J(v int) RF1R_Bits {
	return RF1R_Bits(bits.Make32(v, uint32(mask)))
}

type RF1R struct{ mmio.U32 }

func (r *RF1R) Bits(mask RF1R_Bits) RF1R_Bits { return RF1R_Bits(r.U32.Bits(uint32(mask))) }
func (r *RF1R) StoreBits(mask, b RF1R_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *RF1R) SetBits(mask RF1R_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *RF1R) ClearBits(mask RF1R_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *RF1R) Load() RF1R_Bits               { return RF1R_Bits(r.U32.Load()) }
func (r *RF1R) Store(b RF1R_Bits)             { r.U32.Store(uint32(b)) }

type RF1R_Mask struct{ mmio.UM32 }

func (rm RF1R_Mask) Load() RF1R_Bits   { return RF1R_Bits(rm.UM32.Load()) }
func (rm RF1R_Mask) Store(b RF1R_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FMP1() RF1R_Mask {
	return RF1R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(FMP1)}}
}

func (p *CAN_Periph) FULL1() RF1R_Mask {
	return RF1R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(FULL1)}}
}

func (p *CAN_Periph) FOVR1() RF1R_Mask {
	return RF1R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(FOVR1)}}
}

func (p *CAN_Periph) RFOM1() RF1R_Mask {
	return RF1R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(RFOM1)}}
}

type IER_Bits uint32

func (b IER_Bits) Field(mask IER_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask IER_Bits) J(v int) IER_Bits {
	return IER_Bits(bits.Make32(v, uint32(mask)))
}

type IER struct{ mmio.U32 }

func (r *IER) Bits(mask IER_Bits) IER_Bits { return IER_Bits(r.U32.Bits(uint32(mask))) }
func (r *IER) StoreBits(mask, b IER_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IER) SetBits(mask IER_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IER) ClearBits(mask IER_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IER) Load() IER_Bits              { return IER_Bits(r.U32.Load()) }
func (r *IER) Store(b IER_Bits)            { r.U32.Store(uint32(b)) }

type IER_Mask struct{ mmio.UM32 }

func (rm IER_Mask) Load() IER_Bits   { return IER_Bits(rm.UM32.Load()) }
func (rm IER_Mask) Store(b IER_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) TMEIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(TMEIE)}}
}

func (p *CAN_Periph) FMPIE0() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(FMPIE0)}}
}

func (p *CAN_Periph) FFIE0() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(FFIE0)}}
}

func (p *CAN_Periph) FOVIE0() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(FOVIE0)}}
}

func (p *CAN_Periph) FMPIE1() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(FMPIE1)}}
}

func (p *CAN_Periph) FFIE1() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(FFIE1)}}
}

func (p *CAN_Periph) FOVIE1() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(FOVIE1)}}
}

func (p *CAN_Periph) EWGIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(EWGIE)}}
}

func (p *CAN_Periph) EPVIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(EPVIE)}}
}

func (p *CAN_Periph) BOFIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(BOFIE)}}
}

func (p *CAN_Periph) LECIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(LECIE)}}
}

func (p *CAN_Periph) ERRIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(ERRIE)}}
}

func (p *CAN_Periph) WKUIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(WKUIE)}}
}

func (p *CAN_Periph) SLKIE() IER_Mask {
	return IER_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(SLKIE)}}
}

type ESR_Bits uint32

func (b ESR_Bits) Field(mask ESR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ESR_Bits) J(v int) ESR_Bits {
	return ESR_Bits(bits.Make32(v, uint32(mask)))
}

type ESR struct{ mmio.U32 }

func (r *ESR) Bits(mask ESR_Bits) ESR_Bits { return ESR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ESR) StoreBits(mask, b ESR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ESR) SetBits(mask ESR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ESR) ClearBits(mask ESR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ESR) Load() ESR_Bits              { return ESR_Bits(r.U32.Load()) }
func (r *ESR) Store(b ESR_Bits)            { r.U32.Store(uint32(b)) }

type ESR_Mask struct{ mmio.UM32 }

func (rm ESR_Mask) Load() ESR_Bits   { return ESR_Bits(rm.UM32.Load()) }
func (rm ESR_Mask) Store(b ESR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) EWGF() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(EWGF)}}
}

func (p *CAN_Periph) EPVF() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(EPVF)}}
}

func (p *CAN_Periph) BOFF() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(BOFF)}}
}

func (p *CAN_Periph) LEC() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(LEC)}}
}

func (p *CAN_Periph) TEC() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(TEC)}}
}

func (p *CAN_Periph) REC() ESR_Mask {
	return ESR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(REC)}}
}

type BTR_Bits uint32

func (b BTR_Bits) Field(mask BTR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask BTR_Bits) J(v int) BTR_Bits {
	return BTR_Bits(bits.Make32(v, uint32(mask)))
}

type BTR struct{ mmio.U32 }

func (r *BTR) Bits(mask BTR_Bits) BTR_Bits { return BTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *BTR) StoreBits(mask, b BTR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *BTR) SetBits(mask BTR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *BTR) ClearBits(mask BTR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *BTR) Load() BTR_Bits              { return BTR_Bits(r.U32.Load()) }
func (r *BTR) Store(b BTR_Bits)            { r.U32.Store(uint32(b)) }

type BTR_Mask struct{ mmio.UM32 }

func (rm BTR_Mask) Load() BTR_Bits   { return BTR_Bits(rm.UM32.Load()) }
func (rm BTR_Mask) Store(b BTR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) BRP() BTR_Mask {
	return BTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(BRP)}}
}

func (p *CAN_Periph) TS1() BTR_Mask {
	return BTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(TS1)}}
}

func (p *CAN_Periph) TS2() BTR_Mask {
	return BTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(TS2)}}
}

func (p *CAN_Periph) SJW() BTR_Mask {
	return BTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(SJW)}}
}

func (p *CAN_Periph) LBKM() BTR_Mask {
	return BTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(LBKM)}}
}

func (p *CAN_Periph) SILM() BTR_Mask {
	return BTR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(SILM)}}
}

type FMR_Bits uint32

func (b FMR_Bits) Field(mask FMR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FMR_Bits) J(v int) FMR_Bits {
	return FMR_Bits(bits.Make32(v, uint32(mask)))
}

type FMR struct{ mmio.U32 }

func (r *FMR) Bits(mask FMR_Bits) FMR_Bits { return FMR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FMR) StoreBits(mask, b FMR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FMR) SetBits(mask FMR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *FMR) ClearBits(mask FMR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *FMR) Load() FMR_Bits              { return FMR_Bits(r.U32.Load()) }
func (r *FMR) Store(b FMR_Bits)            { r.U32.Store(uint32(b)) }

type FMR_Mask struct{ mmio.UM32 }

func (rm FMR_Mask) Load() FMR_Bits   { return FMR_Bits(rm.UM32.Load()) }
func (rm FMR_Mask) Store(b FMR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FINIT() FMR_Mask {
	return FMR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 432)), uint32(FINIT)}}
}

type FM1R_Bits uint32

func (b FM1R_Bits) Field(mask FM1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FM1R_Bits) J(v int) FM1R_Bits {
	return FM1R_Bits(bits.Make32(v, uint32(mask)))
}

type FM1R struct{ mmio.U32 }

func (r *FM1R) Bits(mask FM1R_Bits) FM1R_Bits { return FM1R_Bits(r.U32.Bits(uint32(mask))) }
func (r *FM1R) StoreBits(mask, b FM1R_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FM1R) SetBits(mask FM1R_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *FM1R) ClearBits(mask FM1R_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *FM1R) Load() FM1R_Bits               { return FM1R_Bits(r.U32.Load()) }
func (r *FM1R) Store(b FM1R_Bits)             { r.U32.Store(uint32(b)) }

type FM1R_Mask struct{ mmio.UM32 }

func (rm FM1R_Mask) Load() FM1R_Bits   { return FM1R_Bits(rm.UM32.Load()) }
func (rm FM1R_Mask) Store(b FM1R_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FBM() FM1R_Mask {
	return FM1R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 436)), uint32(FBM)}}
}

type FS1R_Bits uint32

func (b FS1R_Bits) Field(mask FS1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FS1R_Bits) J(v int) FS1R_Bits {
	return FS1R_Bits(bits.Make32(v, uint32(mask)))
}

type FS1R struct{ mmio.U32 }

func (r *FS1R) Bits(mask FS1R_Bits) FS1R_Bits { return FS1R_Bits(r.U32.Bits(uint32(mask))) }
func (r *FS1R) StoreBits(mask, b FS1R_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FS1R) SetBits(mask FS1R_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *FS1R) ClearBits(mask FS1R_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *FS1R) Load() FS1R_Bits               { return FS1R_Bits(r.U32.Load()) }
func (r *FS1R) Store(b FS1R_Bits)             { r.U32.Store(uint32(b)) }

type FS1R_Mask struct{ mmio.UM32 }

func (rm FS1R_Mask) Load() FS1R_Bits   { return FS1R_Bits(rm.UM32.Load()) }
func (rm FS1R_Mask) Store(b FS1R_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FSC() FS1R_Mask {
	return FS1R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 444)), uint32(FSC)}}
}

type FFA1R_Bits uint32

func (b FFA1R_Bits) Field(mask FFA1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FFA1R_Bits) J(v int) FFA1R_Bits {
	return FFA1R_Bits(bits.Make32(v, uint32(mask)))
}

type FFA1R struct{ mmio.U32 }

func (r *FFA1R) Bits(mask FFA1R_Bits) FFA1R_Bits { return FFA1R_Bits(r.U32.Bits(uint32(mask))) }
func (r *FFA1R) StoreBits(mask, b FFA1R_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FFA1R) SetBits(mask FFA1R_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *FFA1R) ClearBits(mask FFA1R_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *FFA1R) Load() FFA1R_Bits                { return FFA1R_Bits(r.U32.Load()) }
func (r *FFA1R) Store(b FFA1R_Bits)              { r.U32.Store(uint32(b)) }

type FFA1R_Mask struct{ mmio.UM32 }

func (rm FFA1R_Mask) Load() FFA1R_Bits   { return FFA1R_Bits(rm.UM32.Load()) }
func (rm FFA1R_Mask) Store(b FFA1R_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FFA() FFA1R_Mask {
	return FFA1R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 452)), uint32(FFA)}}
}

type FA1R_Bits uint32

func (b FA1R_Bits) Field(mask FA1R_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask FA1R_Bits) J(v int) FA1R_Bits {
	return FA1R_Bits(bits.Make32(v, uint32(mask)))
}

type FA1R struct{ mmio.U32 }

func (r *FA1R) Bits(mask FA1R_Bits) FA1R_Bits { return FA1R_Bits(r.U32.Bits(uint32(mask))) }
func (r *FA1R) StoreBits(mask, b FA1R_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FA1R) SetBits(mask FA1R_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *FA1R) ClearBits(mask FA1R_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *FA1R) Load() FA1R_Bits               { return FA1R_Bits(r.U32.Load()) }
func (r *FA1R) Store(b FA1R_Bits)             { r.U32.Store(uint32(b)) }

type FA1R_Mask struct{ mmio.UM32 }

func (rm FA1R_Mask) Load() FA1R_Bits   { return FA1R_Bits(rm.UM32.Load()) }
func (rm FA1R_Mask) Store(b FA1R_Bits) { rm.UM32.Store(uint32(b)) }

func (p *CAN_Periph) FACT() FA1R_Mask {
	return FA1R_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 460)), uint32(FACT)}}
}
