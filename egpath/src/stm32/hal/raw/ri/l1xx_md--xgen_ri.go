// +build l1xx_md

package ri

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"bits"
	"mmio"
	"unsafe"

	"stm32/o/l1xx_md/mmap"
)

type RI_Periph struct {
	ICR    ICR
	ASCR1  ASCR1
	ASCR2  ASCR2
	HYSCR1 HYSCR1
	HYSCR2 HYSCR2
	HYSCR3 HYSCR3
	HYSCR4 HYSCR4
	ASMR1  ASMR1
	CMR1   CMR1
	CICR1  CICR1
	ASMR2  ASMR2
	CMR2   CMR2
	CICR2  CICR2
	ASMR3  ASMR3
	CMR3   CMR3
	CICR3  CICR3
	ASMR4  ASMR4
	CMR4   CMR4
	CICR4  CICR4
	ASMR5  ASMR5
	CMR5   CMR5
	CICR5  CICR5
}

func (p *RI_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var RI = (*RI_Periph)(unsafe.Pointer(uintptr(mmap.RI_BASE)))

type ICR_Bits uint32

func (b ICR_Bits) Field(mask ICR_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ICR_Bits) J(v int) ICR_Bits {
	return ICR_Bits(bits.Make32(v, uint32(mask)))
}

type ICR struct{ mmio.U32 }

func (r *ICR) Bits(mask ICR_Bits) ICR_Bits { return ICR_Bits(r.U32.Bits(uint32(mask))) }
func (r *ICR) StoreBits(mask, b ICR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ICR) SetBits(mask ICR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *ICR) ClearBits(mask ICR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *ICR) Load() ICR_Bits              { return ICR_Bits(r.U32.Load()) }
func (r *ICR) Store(b ICR_Bits)            { r.U32.Store(uint32(b)) }

type ICR_Mask struct{ mmio.UM32 }

func (rm ICR_Mask) Load() ICR_Bits   { return ICR_Bits(rm.UM32.Load()) }
func (rm ICR_Mask) Store(b ICR_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) IC1Z() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IC1Z)}}
}

func (p *RI_Periph) IC2Z() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IC2Z)}}
}

func (p *RI_Periph) IC3Z() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IC3Z)}}
}

func (p *RI_Periph) IC4Z() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IC4Z)}}
}

func (p *RI_Periph) TIM() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(TIM)}}
}

func (p *RI_Periph) IC1() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IC1)}}
}

func (p *RI_Periph) IC2() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IC2)}}
}

func (p *RI_Periph) IC3() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IC3)}}
}

func (p *RI_Periph) IC4() ICR_Mask {
	return ICR_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 0)), uint32(IC4)}}
}

type ASCR1_Bits uint32

func (b ASCR1_Bits) Field(mask ASCR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASCR1_Bits) J(v int) ASCR1_Bits {
	return ASCR1_Bits(bits.Make32(v, uint32(mask)))
}

type ASCR1 struct{ mmio.U32 }

func (r *ASCR1) Bits(mask ASCR1_Bits) ASCR1_Bits { return ASCR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *ASCR1) StoreBits(mask, b ASCR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ASCR1) SetBits(mask ASCR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ASCR1) ClearBits(mask ASCR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ASCR1) Load() ASCR1_Bits                { return ASCR1_Bits(r.U32.Load()) }
func (r *ASCR1) Store(b ASCR1_Bits)              { r.U32.Store(uint32(b)) }

type ASCR1_Mask struct{ mmio.UM32 }

func (rm ASCR1_Mask) Load() ASCR1_Bits   { return ASCR1_Bits(rm.UM32.Load()) }
func (rm ASCR1_Mask) Store(b ASCR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) CH() ASCR1_Mask {
	return ASCR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(CH)}}
}

func (p *RI_Periph) CH_31() ASCR1_Mask {
	return ASCR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(CH_31)}}
}

func (p *RI_Periph) VCOMP() ASCR1_Mask {
	return ASCR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(VCOMP)}}
}

func (p *RI_Periph) SCM() ASCR1_Mask {
	return ASCR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 4)), uint32(SCM)}}
}

type ASCR2_Bits uint32

func (b ASCR2_Bits) Field(mask ASCR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASCR2_Bits) J(v int) ASCR2_Bits {
	return ASCR2_Bits(bits.Make32(v, uint32(mask)))
}

type ASCR2 struct{ mmio.U32 }

func (r *ASCR2) Bits(mask ASCR2_Bits) ASCR2_Bits { return ASCR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *ASCR2) StoreBits(mask, b ASCR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ASCR2) SetBits(mask ASCR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ASCR2) ClearBits(mask ASCR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ASCR2) Load() ASCR2_Bits                { return ASCR2_Bits(r.U32.Load()) }
func (r *ASCR2) Store(b ASCR2_Bits)              { r.U32.Store(uint32(b)) }

type ASCR2_Mask struct{ mmio.UM32 }

func (rm ASCR2_Mask) Load() ASCR2_Bits   { return ASCR2_Bits(rm.UM32.Load()) }
func (rm ASCR2_Mask) Store(b ASCR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) GR10_1() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR10_1)}}
}

func (p *RI_Periph) GR10_2() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR10_2)}}
}

func (p *RI_Periph) GR10_3() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR10_3)}}
}

func (p *RI_Periph) GR10_4() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR10_4)}}
}

func (p *RI_Periph) GR6_1() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR6_1)}}
}

func (p *RI_Periph) GR6_2() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR6_2)}}
}

func (p *RI_Periph) GR5_1() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR5_1)}}
}

func (p *RI_Periph) GR5_2() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR5_2)}}
}

func (p *RI_Periph) GR5_3() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR5_3)}}
}

func (p *RI_Periph) GR4_1() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR4_1)}}
}

func (p *RI_Periph) GR4_2() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR4_2)}}
}

func (p *RI_Periph) GR4_3() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR4_3)}}
}

func (p *RI_Periph) GR4_4() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR4_4)}}
}

func (p *RI_Periph) CH0b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH0b)}}
}

func (p *RI_Periph) CH1b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH1b)}}
}

func (p *RI_Periph) CH2b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH2b)}}
}

func (p *RI_Periph) CH3b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH3b)}}
}

func (p *RI_Periph) CH6b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH6b)}}
}

func (p *RI_Periph) CH7b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH7b)}}
}

func (p *RI_Periph) CH8b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH8b)}}
}

func (p *RI_Periph) CH9b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH9b)}}
}

func (p *RI_Periph) CH10b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH10b)}}
}

func (p *RI_Periph) CH11b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH11b)}}
}

func (p *RI_Periph) CH12b() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(CH12b)}}
}

func (p *RI_Periph) GR6_3() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR6_3)}}
}

func (p *RI_Periph) GR6_4() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR6_4)}}
}

func (p *RI_Periph) GR5_4() ASCR2_Mask {
	return ASCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 8)), uint32(GR5_4)}}
}

type HYSCR1_Bits uint32

func (b HYSCR1_Bits) Field(mask HYSCR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HYSCR1_Bits) J(v int) HYSCR1_Bits {
	return HYSCR1_Bits(bits.Make32(v, uint32(mask)))
}

type HYSCR1 struct{ mmio.U32 }

func (r *HYSCR1) Bits(mask HYSCR1_Bits) HYSCR1_Bits { return HYSCR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *HYSCR1) StoreBits(mask, b HYSCR1_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HYSCR1) SetBits(mask HYSCR1_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *HYSCR1) ClearBits(mask HYSCR1_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *HYSCR1) Load() HYSCR1_Bits                 { return HYSCR1_Bits(r.U32.Load()) }
func (r *HYSCR1) Store(b HYSCR1_Bits)               { r.U32.Store(uint32(b)) }

type HYSCR1_Mask struct{ mmio.UM32 }

func (rm HYSCR1_Mask) Load() HYSCR1_Bits   { return HYSCR1_Bits(rm.UM32.Load()) }
func (rm HYSCR1_Mask) Store(b HYSCR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PA() HYSCR1_Mask {
	return HYSCR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PA)}}
}

func (p *RI_Periph) PB() HYSCR1_Mask {
	return HYSCR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 12)), uint32(PB)}}
}

type HYSCR2_Bits uint32

func (b HYSCR2_Bits) Field(mask HYSCR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HYSCR2_Bits) J(v int) HYSCR2_Bits {
	return HYSCR2_Bits(bits.Make32(v, uint32(mask)))
}

type HYSCR2 struct{ mmio.U32 }

func (r *HYSCR2) Bits(mask HYSCR2_Bits) HYSCR2_Bits { return HYSCR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *HYSCR2) StoreBits(mask, b HYSCR2_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HYSCR2) SetBits(mask HYSCR2_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *HYSCR2) ClearBits(mask HYSCR2_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *HYSCR2) Load() HYSCR2_Bits                 { return HYSCR2_Bits(r.U32.Load()) }
func (r *HYSCR2) Store(b HYSCR2_Bits)               { r.U32.Store(uint32(b)) }

type HYSCR2_Mask struct{ mmio.UM32 }

func (rm HYSCR2_Mask) Load() HYSCR2_Bits   { return HYSCR2_Bits(rm.UM32.Load()) }
func (rm HYSCR2_Mask) Store(b HYSCR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PC() HYSCR2_Mask {
	return HYSCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(PC)}}
}

func (p *RI_Periph) PD() HYSCR2_Mask {
	return HYSCR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 16)), uint32(PD)}}
}

type HYSCR3_Bits uint32

func (b HYSCR3_Bits) Field(mask HYSCR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HYSCR3_Bits) J(v int) HYSCR3_Bits {
	return HYSCR3_Bits(bits.Make32(v, uint32(mask)))
}

type HYSCR3 struct{ mmio.U32 }

func (r *HYSCR3) Bits(mask HYSCR3_Bits) HYSCR3_Bits { return HYSCR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *HYSCR3) StoreBits(mask, b HYSCR3_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HYSCR3) SetBits(mask HYSCR3_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *HYSCR3) ClearBits(mask HYSCR3_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *HYSCR3) Load() HYSCR3_Bits                 { return HYSCR3_Bits(r.U32.Load()) }
func (r *HYSCR3) Store(b HYSCR3_Bits)               { r.U32.Store(uint32(b)) }

type HYSCR3_Mask struct{ mmio.UM32 }

func (rm HYSCR3_Mask) Load() HYSCR3_Bits   { return HYSCR3_Bits(rm.UM32.Load()) }
func (rm HYSCR3_Mask) Store(b HYSCR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PF() HYSCR3_Mask {
	return HYSCR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 20)), uint32(PF)}}
}

type HYSCR4_Bits uint32

func (b HYSCR4_Bits) Field(mask HYSCR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask HYSCR4_Bits) J(v int) HYSCR4_Bits {
	return HYSCR4_Bits(bits.Make32(v, uint32(mask)))
}

type HYSCR4 struct{ mmio.U32 }

func (r *HYSCR4) Bits(mask HYSCR4_Bits) HYSCR4_Bits { return HYSCR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *HYSCR4) StoreBits(mask, b HYSCR4_Bits)     { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *HYSCR4) SetBits(mask HYSCR4_Bits)          { r.U32.SetBits(uint32(mask)) }
func (r *HYSCR4) ClearBits(mask HYSCR4_Bits)        { r.U32.ClearBits(uint32(mask)) }
func (r *HYSCR4) Load() HYSCR4_Bits                 { return HYSCR4_Bits(r.U32.Load()) }
func (r *HYSCR4) Store(b HYSCR4_Bits)               { r.U32.Store(uint32(b)) }

type HYSCR4_Mask struct{ mmio.UM32 }

func (rm HYSCR4_Mask) Load() HYSCR4_Bits   { return HYSCR4_Bits(rm.UM32.Load()) }
func (rm HYSCR4_Mask) Store(b HYSCR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PG() HYSCR4_Mask {
	return HYSCR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 24)), uint32(PG)}}
}

type ASMR1_Bits uint32

func (b ASMR1_Bits) Field(mask ASMR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASMR1_Bits) J(v int) ASMR1_Bits {
	return ASMR1_Bits(bits.Make32(v, uint32(mask)))
}

type ASMR1 struct{ mmio.U32 }

func (r *ASMR1) Bits(mask ASMR1_Bits) ASMR1_Bits { return ASMR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *ASMR1) StoreBits(mask, b ASMR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ASMR1) SetBits(mask ASMR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ASMR1) ClearBits(mask ASMR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ASMR1) Load() ASMR1_Bits                { return ASMR1_Bits(r.U32.Load()) }
func (r *ASMR1) Store(b ASMR1_Bits)              { r.U32.Store(uint32(b)) }

type ASMR1_Mask struct{ mmio.UM32 }

func (rm ASMR1_Mask) Load() ASMR1_Bits   { return ASMR1_Bits(rm.UM32.Load()) }
func (rm ASMR1_Mask) Store(b ASMR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PA() ASMR1_Mask {
	return ASMR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 28)), uint32(PA)}}
}

type CMR1_Bits uint32

func (b CMR1_Bits) Field(mask CMR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMR1_Bits) J(v int) CMR1_Bits {
	return CMR1_Bits(bits.Make32(v, uint32(mask)))
}

type CMR1 struct{ mmio.U32 }

func (r *CMR1) Bits(mask CMR1_Bits) CMR1_Bits { return CMR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMR1) StoreBits(mask, b CMR1_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMR1) SetBits(mask CMR1_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CMR1) ClearBits(mask CMR1_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CMR1) Load() CMR1_Bits               { return CMR1_Bits(r.U32.Load()) }
func (r *CMR1) Store(b CMR1_Bits)             { r.U32.Store(uint32(b)) }

type CMR1_Mask struct{ mmio.UM32 }

func (rm CMR1_Mask) Load() CMR1_Bits   { return CMR1_Bits(rm.UM32.Load()) }
func (rm CMR1_Mask) Store(b CMR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PA() CMR1_Mask {
	return CMR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 32)), uint32(PA)}}
}

type CICR1_Bits uint32

func (b CICR1_Bits) Field(mask CICR1_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CICR1_Bits) J(v int) CICR1_Bits {
	return CICR1_Bits(bits.Make32(v, uint32(mask)))
}

type CICR1 struct{ mmio.U32 }

func (r *CICR1) Bits(mask CICR1_Bits) CICR1_Bits { return CICR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CICR1) StoreBits(mask, b CICR1_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CICR1) SetBits(mask CICR1_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CICR1) ClearBits(mask CICR1_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CICR1) Load() CICR1_Bits                { return CICR1_Bits(r.U32.Load()) }
func (r *CICR1) Store(b CICR1_Bits)              { r.U32.Store(uint32(b)) }

type CICR1_Mask struct{ mmio.UM32 }

func (rm CICR1_Mask) Load() CICR1_Bits   { return CICR1_Bits(rm.UM32.Load()) }
func (rm CICR1_Mask) Store(b CICR1_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PA() CICR1_Mask {
	return CICR1_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 36)), uint32(PA)}}
}

type ASMR2_Bits uint32

func (b ASMR2_Bits) Field(mask ASMR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASMR2_Bits) J(v int) ASMR2_Bits {
	return ASMR2_Bits(bits.Make32(v, uint32(mask)))
}

type ASMR2 struct{ mmio.U32 }

func (r *ASMR2) Bits(mask ASMR2_Bits) ASMR2_Bits { return ASMR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *ASMR2) StoreBits(mask, b ASMR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ASMR2) SetBits(mask ASMR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ASMR2) ClearBits(mask ASMR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ASMR2) Load() ASMR2_Bits                { return ASMR2_Bits(r.U32.Load()) }
func (r *ASMR2) Store(b ASMR2_Bits)              { r.U32.Store(uint32(b)) }

type ASMR2_Mask struct{ mmio.UM32 }

func (rm ASMR2_Mask) Load() ASMR2_Bits   { return ASMR2_Bits(rm.UM32.Load()) }
func (rm ASMR2_Mask) Store(b ASMR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PB() ASMR2_Mask {
	return ASMR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 40)), uint32(PB)}}
}

type CMR2_Bits uint32

func (b CMR2_Bits) Field(mask CMR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMR2_Bits) J(v int) CMR2_Bits {
	return CMR2_Bits(bits.Make32(v, uint32(mask)))
}

type CMR2 struct{ mmio.U32 }

func (r *CMR2) Bits(mask CMR2_Bits) CMR2_Bits { return CMR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMR2) StoreBits(mask, b CMR2_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMR2) SetBits(mask CMR2_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CMR2) ClearBits(mask CMR2_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CMR2) Load() CMR2_Bits               { return CMR2_Bits(r.U32.Load()) }
func (r *CMR2) Store(b CMR2_Bits)             { r.U32.Store(uint32(b)) }

type CMR2_Mask struct{ mmio.UM32 }

func (rm CMR2_Mask) Load() CMR2_Bits   { return CMR2_Bits(rm.UM32.Load()) }
func (rm CMR2_Mask) Store(b CMR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PB() CMR2_Mask {
	return CMR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 44)), uint32(PB)}}
}

type CICR2_Bits uint32

func (b CICR2_Bits) Field(mask CICR2_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CICR2_Bits) J(v int) CICR2_Bits {
	return CICR2_Bits(bits.Make32(v, uint32(mask)))
}

type CICR2 struct{ mmio.U32 }

func (r *CICR2) Bits(mask CICR2_Bits) CICR2_Bits { return CICR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CICR2) StoreBits(mask, b CICR2_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CICR2) SetBits(mask CICR2_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CICR2) ClearBits(mask CICR2_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CICR2) Load() CICR2_Bits                { return CICR2_Bits(r.U32.Load()) }
func (r *CICR2) Store(b CICR2_Bits)              { r.U32.Store(uint32(b)) }

type CICR2_Mask struct{ mmio.UM32 }

func (rm CICR2_Mask) Load() CICR2_Bits   { return CICR2_Bits(rm.UM32.Load()) }
func (rm CICR2_Mask) Store(b CICR2_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PB() CICR2_Mask {
	return CICR2_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 48)), uint32(PB)}}
}

type ASMR3_Bits uint32

func (b ASMR3_Bits) Field(mask ASMR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASMR3_Bits) J(v int) ASMR3_Bits {
	return ASMR3_Bits(bits.Make32(v, uint32(mask)))
}

type ASMR3 struct{ mmio.U32 }

func (r *ASMR3) Bits(mask ASMR3_Bits) ASMR3_Bits { return ASMR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *ASMR3) StoreBits(mask, b ASMR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ASMR3) SetBits(mask ASMR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ASMR3) ClearBits(mask ASMR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ASMR3) Load() ASMR3_Bits                { return ASMR3_Bits(r.U32.Load()) }
func (r *ASMR3) Store(b ASMR3_Bits)              { r.U32.Store(uint32(b)) }

type ASMR3_Mask struct{ mmio.UM32 }

func (rm ASMR3_Mask) Load() ASMR3_Bits   { return ASMR3_Bits(rm.UM32.Load()) }
func (rm ASMR3_Mask) Store(b ASMR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PC() ASMR3_Mask {
	return ASMR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 52)), uint32(PC)}}
}

type CMR3_Bits uint32

func (b CMR3_Bits) Field(mask CMR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMR3_Bits) J(v int) CMR3_Bits {
	return CMR3_Bits(bits.Make32(v, uint32(mask)))
}

type CMR3 struct{ mmio.U32 }

func (r *CMR3) Bits(mask CMR3_Bits) CMR3_Bits { return CMR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMR3) StoreBits(mask, b CMR3_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMR3) SetBits(mask CMR3_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CMR3) ClearBits(mask CMR3_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CMR3) Load() CMR3_Bits               { return CMR3_Bits(r.U32.Load()) }
func (r *CMR3) Store(b CMR3_Bits)             { r.U32.Store(uint32(b)) }

type CMR3_Mask struct{ mmio.UM32 }

func (rm CMR3_Mask) Load() CMR3_Bits   { return CMR3_Bits(rm.UM32.Load()) }
func (rm CMR3_Mask) Store(b CMR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PC() CMR3_Mask {
	return CMR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 56)), uint32(PC)}}
}

type CICR3_Bits uint32

func (b CICR3_Bits) Field(mask CICR3_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CICR3_Bits) J(v int) CICR3_Bits {
	return CICR3_Bits(bits.Make32(v, uint32(mask)))
}

type CICR3 struct{ mmio.U32 }

func (r *CICR3) Bits(mask CICR3_Bits) CICR3_Bits { return CICR3_Bits(r.U32.Bits(uint32(mask))) }
func (r *CICR3) StoreBits(mask, b CICR3_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CICR3) SetBits(mask CICR3_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CICR3) ClearBits(mask CICR3_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CICR3) Load() CICR3_Bits                { return CICR3_Bits(r.U32.Load()) }
func (r *CICR3) Store(b CICR3_Bits)              { r.U32.Store(uint32(b)) }

type CICR3_Mask struct{ mmio.UM32 }

func (rm CICR3_Mask) Load() CICR3_Bits   { return CICR3_Bits(rm.UM32.Load()) }
func (rm CICR3_Mask) Store(b CICR3_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PC() CICR3_Mask {
	return CICR3_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 60)), uint32(PC)}}
}

type ASMR4_Bits uint32

func (b ASMR4_Bits) Field(mask ASMR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASMR4_Bits) J(v int) ASMR4_Bits {
	return ASMR4_Bits(bits.Make32(v, uint32(mask)))
}

type ASMR4 struct{ mmio.U32 }

func (r *ASMR4) Bits(mask ASMR4_Bits) ASMR4_Bits { return ASMR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *ASMR4) StoreBits(mask, b ASMR4_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ASMR4) SetBits(mask ASMR4_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ASMR4) ClearBits(mask ASMR4_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ASMR4) Load() ASMR4_Bits                { return ASMR4_Bits(r.U32.Load()) }
func (r *ASMR4) Store(b ASMR4_Bits)              { r.U32.Store(uint32(b)) }

type ASMR4_Mask struct{ mmio.UM32 }

func (rm ASMR4_Mask) Load() ASMR4_Bits   { return ASMR4_Bits(rm.UM32.Load()) }
func (rm ASMR4_Mask) Store(b ASMR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PF() ASMR4_Mask {
	return ASMR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 64)), uint32(PF)}}
}

type CMR4_Bits uint32

func (b CMR4_Bits) Field(mask CMR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMR4_Bits) J(v int) CMR4_Bits {
	return CMR4_Bits(bits.Make32(v, uint32(mask)))
}

type CMR4 struct{ mmio.U32 }

func (r *CMR4) Bits(mask CMR4_Bits) CMR4_Bits { return CMR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMR4) StoreBits(mask, b CMR4_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMR4) SetBits(mask CMR4_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CMR4) ClearBits(mask CMR4_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CMR4) Load() CMR4_Bits               { return CMR4_Bits(r.U32.Load()) }
func (r *CMR4) Store(b CMR4_Bits)             { r.U32.Store(uint32(b)) }

type CMR4_Mask struct{ mmio.UM32 }

func (rm CMR4_Mask) Load() CMR4_Bits   { return CMR4_Bits(rm.UM32.Load()) }
func (rm CMR4_Mask) Store(b CMR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PF() CMR4_Mask {
	return CMR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 68)), uint32(PF)}}
}

type CICR4_Bits uint32

func (b CICR4_Bits) Field(mask CICR4_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CICR4_Bits) J(v int) CICR4_Bits {
	return CICR4_Bits(bits.Make32(v, uint32(mask)))
}

type CICR4 struct{ mmio.U32 }

func (r *CICR4) Bits(mask CICR4_Bits) CICR4_Bits { return CICR4_Bits(r.U32.Bits(uint32(mask))) }
func (r *CICR4) StoreBits(mask, b CICR4_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CICR4) SetBits(mask CICR4_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CICR4) ClearBits(mask CICR4_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CICR4) Load() CICR4_Bits                { return CICR4_Bits(r.U32.Load()) }
func (r *CICR4) Store(b CICR4_Bits)              { r.U32.Store(uint32(b)) }

type CICR4_Mask struct{ mmio.UM32 }

func (rm CICR4_Mask) Load() CICR4_Bits   { return CICR4_Bits(rm.UM32.Load()) }
func (rm CICR4_Mask) Store(b CICR4_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PF() CICR4_Mask {
	return CICR4_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 72)), uint32(PF)}}
}

type ASMR5_Bits uint32

func (b ASMR5_Bits) Field(mask ASMR5_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask ASMR5_Bits) J(v int) ASMR5_Bits {
	return ASMR5_Bits(bits.Make32(v, uint32(mask)))
}

type ASMR5 struct{ mmio.U32 }

func (r *ASMR5) Bits(mask ASMR5_Bits) ASMR5_Bits { return ASMR5_Bits(r.U32.Bits(uint32(mask))) }
func (r *ASMR5) StoreBits(mask, b ASMR5_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *ASMR5) SetBits(mask ASMR5_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *ASMR5) ClearBits(mask ASMR5_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *ASMR5) Load() ASMR5_Bits                { return ASMR5_Bits(r.U32.Load()) }
func (r *ASMR5) Store(b ASMR5_Bits)              { r.U32.Store(uint32(b)) }

type ASMR5_Mask struct{ mmio.UM32 }

func (rm ASMR5_Mask) Load() ASMR5_Bits   { return ASMR5_Bits(rm.UM32.Load()) }
func (rm ASMR5_Mask) Store(b ASMR5_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PG() ASMR5_Mask {
	return ASMR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 76)), uint32(PG)}}
}

type CMR5_Bits uint32

func (b CMR5_Bits) Field(mask CMR5_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CMR5_Bits) J(v int) CMR5_Bits {
	return CMR5_Bits(bits.Make32(v, uint32(mask)))
}

type CMR5 struct{ mmio.U32 }

func (r *CMR5) Bits(mask CMR5_Bits) CMR5_Bits { return CMR5_Bits(r.U32.Bits(uint32(mask))) }
func (r *CMR5) StoreBits(mask, b CMR5_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CMR5) SetBits(mask CMR5_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *CMR5) ClearBits(mask CMR5_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *CMR5) Load() CMR5_Bits               { return CMR5_Bits(r.U32.Load()) }
func (r *CMR5) Store(b CMR5_Bits)             { r.U32.Store(uint32(b)) }

type CMR5_Mask struct{ mmio.UM32 }

func (rm CMR5_Mask) Load() CMR5_Bits   { return CMR5_Bits(rm.UM32.Load()) }
func (rm CMR5_Mask) Store(b CMR5_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PG() CMR5_Mask {
	return CMR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 80)), uint32(PG)}}
}

type CICR5_Bits uint32

func (b CICR5_Bits) Field(mask CICR5_Bits) int {
	return bits.Field32(uint32(b), uint32(mask))
}
func (mask CICR5_Bits) J(v int) CICR5_Bits {
	return CICR5_Bits(bits.Make32(v, uint32(mask)))
}

type CICR5 struct{ mmio.U32 }

func (r *CICR5) Bits(mask CICR5_Bits) CICR5_Bits { return CICR5_Bits(r.U32.Bits(uint32(mask))) }
func (r *CICR5) StoreBits(mask, b CICR5_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CICR5) SetBits(mask CICR5_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CICR5) ClearBits(mask CICR5_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CICR5) Load() CICR5_Bits                { return CICR5_Bits(r.U32.Load()) }
func (r *CICR5) Store(b CICR5_Bits)              { r.U32.Store(uint32(b)) }

type CICR5_Mask struct{ mmio.UM32 }

func (rm CICR5_Mask) Load() CICR5_Bits   { return CICR5_Bits(rm.UM32.Load()) }
func (rm CICR5_Mask) Store(b CICR5_Bits) { rm.UM32.Store(uint32(b)) }

func (p *RI_Periph) PG() CICR5_Mask {
	return CICR5_Mask{mmio.UM32{(*mmio.U32)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 84)), uint32(PG)}}
}
