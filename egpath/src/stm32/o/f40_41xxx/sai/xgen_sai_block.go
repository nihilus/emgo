package sai

// DO NOT EDIT THIS FILE. GENERATED BY xgen.

import (
	"mmio"
	"unsafe"

	"stm32/o/f40_41xxx/mmap"
)

type SAI_Block_Periph struct {
	CR1   CR1
	CR2   CR2
	FRCR  FRCR
	SLOTR SLOTR
	IMR   IMR
	SR    SR
	CLRFR CLRFR
	DR    DR
}

func (p *SAI_Block_Periph) BaseAddr() uintptr {
	return uintptr(unsafe.Pointer(p))
}

var SAI1_Block_A = (*SAI_Block_Periph)(unsafe.Pointer(uintptr(mmap.SAI1_Block_A_BASE)))

var SAI1_Block_B = (*SAI_Block_Periph)(unsafe.Pointer(uintptr(mmap.SAI1_Block_B_BASE)))

type CR1_Bits uint32

type CR1 struct{ mmio.U32 }

func (r *CR1) Bits(mask CR1_Bits) CR1_Bits { return CR1_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR1) StoreBits(mask, b CR1_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR1) SetBits(mask CR1_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR1) ClearBits(mask CR1_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR1) Load() CR1_Bits              { return CR1_Bits(r.U32.Load()) }
func (r *CR1) Store(b CR1_Bits)            { r.U32.Store(uint32(b)) }

type CR1_Mask struct{ mmio.UM32 }

func (rm CR1_Mask) Load() CR1_Bits   { return CR1_Bits(rm.UM32.Load()) }
func (rm CR1_Mask) Store(b CR1_Bits) { rm.UM32.Store(uint32(b)) }

type CR2_Bits uint32

type CR2 struct{ mmio.U32 }

func (r *CR2) Bits(mask CR2_Bits) CR2_Bits { return CR2_Bits(r.U32.Bits(uint32(mask))) }
func (r *CR2) StoreBits(mask, b CR2_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CR2) SetBits(mask CR2_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *CR2) ClearBits(mask CR2_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *CR2) Load() CR2_Bits              { return CR2_Bits(r.U32.Load()) }
func (r *CR2) Store(b CR2_Bits)            { r.U32.Store(uint32(b)) }

type CR2_Mask struct{ mmio.UM32 }

func (rm CR2_Mask) Load() CR2_Bits   { return CR2_Bits(rm.UM32.Load()) }
func (rm CR2_Mask) Store(b CR2_Bits) { rm.UM32.Store(uint32(b)) }

type FRCR_Bits uint32

type FRCR struct{ mmio.U32 }

func (r *FRCR) Bits(mask FRCR_Bits) FRCR_Bits { return FRCR_Bits(r.U32.Bits(uint32(mask))) }
func (r *FRCR) StoreBits(mask, b FRCR_Bits)   { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *FRCR) SetBits(mask FRCR_Bits)        { r.U32.SetBits(uint32(mask)) }
func (r *FRCR) ClearBits(mask FRCR_Bits)      { r.U32.ClearBits(uint32(mask)) }
func (r *FRCR) Load() FRCR_Bits               { return FRCR_Bits(r.U32.Load()) }
func (r *FRCR) Store(b FRCR_Bits)             { r.U32.Store(uint32(b)) }

type FRCR_Mask struct{ mmio.UM32 }

func (rm FRCR_Mask) Load() FRCR_Bits   { return FRCR_Bits(rm.UM32.Load()) }
func (rm FRCR_Mask) Store(b FRCR_Bits) { rm.UM32.Store(uint32(b)) }

type SLOTR_Bits uint32

type SLOTR struct{ mmio.U32 }

func (r *SLOTR) Bits(mask SLOTR_Bits) SLOTR_Bits { return SLOTR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SLOTR) StoreBits(mask, b SLOTR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SLOTR) SetBits(mask SLOTR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *SLOTR) ClearBits(mask SLOTR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *SLOTR) Load() SLOTR_Bits                { return SLOTR_Bits(r.U32.Load()) }
func (r *SLOTR) Store(b SLOTR_Bits)              { r.U32.Store(uint32(b)) }

type SLOTR_Mask struct{ mmio.UM32 }

func (rm SLOTR_Mask) Load() SLOTR_Bits   { return SLOTR_Bits(rm.UM32.Load()) }
func (rm SLOTR_Mask) Store(b SLOTR_Bits) { rm.UM32.Store(uint32(b)) }

type IMR_Bits uint32

type IMR struct{ mmio.U32 }

func (r *IMR) Bits(mask IMR_Bits) IMR_Bits { return IMR_Bits(r.U32.Bits(uint32(mask))) }
func (r *IMR) StoreBits(mask, b IMR_Bits)  { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *IMR) SetBits(mask IMR_Bits)       { r.U32.SetBits(uint32(mask)) }
func (r *IMR) ClearBits(mask IMR_Bits)     { r.U32.ClearBits(uint32(mask)) }
func (r *IMR) Load() IMR_Bits              { return IMR_Bits(r.U32.Load()) }
func (r *IMR) Store(b IMR_Bits)            { r.U32.Store(uint32(b)) }

type IMR_Mask struct{ mmio.UM32 }

func (rm IMR_Mask) Load() IMR_Bits   { return IMR_Bits(rm.UM32.Load()) }
func (rm IMR_Mask) Store(b IMR_Bits) { rm.UM32.Store(uint32(b)) }

type SR_Bits uint32

type SR struct{ mmio.U32 }

func (r *SR) Bits(mask SR_Bits) SR_Bits { return SR_Bits(r.U32.Bits(uint32(mask))) }
func (r *SR) StoreBits(mask, b SR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *SR) SetBits(mask SR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *SR) ClearBits(mask SR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *SR) Load() SR_Bits             { return SR_Bits(r.U32.Load()) }
func (r *SR) Store(b SR_Bits)           { r.U32.Store(uint32(b)) }

type SR_Mask struct{ mmio.UM32 }

func (rm SR_Mask) Load() SR_Bits   { return SR_Bits(rm.UM32.Load()) }
func (rm SR_Mask) Store(b SR_Bits) { rm.UM32.Store(uint32(b)) }

type CLRFR_Bits uint32

type CLRFR struct{ mmio.U32 }

func (r *CLRFR) Bits(mask CLRFR_Bits) CLRFR_Bits { return CLRFR_Bits(r.U32.Bits(uint32(mask))) }
func (r *CLRFR) StoreBits(mask, b CLRFR_Bits)    { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *CLRFR) SetBits(mask CLRFR_Bits)         { r.U32.SetBits(uint32(mask)) }
func (r *CLRFR) ClearBits(mask CLRFR_Bits)       { r.U32.ClearBits(uint32(mask)) }
func (r *CLRFR) Load() CLRFR_Bits                { return CLRFR_Bits(r.U32.Load()) }
func (r *CLRFR) Store(b CLRFR_Bits)              { r.U32.Store(uint32(b)) }

type CLRFR_Mask struct{ mmio.UM32 }

func (rm CLRFR_Mask) Load() CLRFR_Bits   { return CLRFR_Bits(rm.UM32.Load()) }
func (rm CLRFR_Mask) Store(b CLRFR_Bits) { rm.UM32.Store(uint32(b)) }

type DR_Bits uint32

type DR struct{ mmio.U32 }

func (r *DR) Bits(mask DR_Bits) DR_Bits { return DR_Bits(r.U32.Bits(uint32(mask))) }
func (r *DR) StoreBits(mask, b DR_Bits) { r.U32.StoreBits(uint32(mask), uint32(b)) }
func (r *DR) SetBits(mask DR_Bits)      { r.U32.SetBits(uint32(mask)) }
func (r *DR) ClearBits(mask DR_Bits)    { r.U32.ClearBits(uint32(mask)) }
func (r *DR) Load() DR_Bits             { return DR_Bits(r.U32.Load()) }
func (r *DR) Store(b DR_Bits)           { r.U32.Store(uint32(b)) }

type DR_Mask struct{ mmio.UM32 }

func (rm DR_Mask) Load() DR_Bits   { return DR_Bits(rm.UM32.Load()) }
func (rm DR_Mask) Store(b DR_Bits) { rm.UM32.Store(uint32(b)) }
