package main

import (
	"delay"
	"mmio"

	"stm32/hal/gpio"
	"stm32/hal/irq"
	"stm32/hal/system"
	"stm32/hal/system/timer/rtc"

	"stm32/hal/raw/rcc"
	"stm32/hal/raw/tim"
)

const PWMmax = 1e4

var led1, led2 *mmio.U16

func init() {
	system.Setup(8, 72/8, false)
	rtc.Setup(32768)

	gpio.B.EnableClock(true)
	leds := gpio.B

	cfg := gpio.Config{Mode: gpio.Alt, Speed: gpio.Low}
	leds.Setup(gpio.Pin7|gpio.Pin6, &cfg)
	rcc.RCC.TIM4EN().Set()
	t := tim.TIM4
	const (
		pwmfreq = 100 // Hz
		pwmmode = 6   // Mode 1
	)
	pclk := system.APB1.Clock()
	if pclk != system.AHB.Clock() {
		pclk *= 2
	}
	t.PSC.U16.Store(uint16(pclk/(PWMmax*pwmfreq) - 1))
	t.ARR.Store(PWMmax - 1)
	t.OC1M().Store(pwmmode << tim.OC1Mn)
	t.OC2M().Store(pwmmode << tim.OC2Mn)
	t.OC1PE().Set()
	t.OC2PE().Set()
	t.CCER.SetBits(tim.CC1E | tim.CC2E)
	t.ARPE().Set()
	t.UG().Set()
	t.CEN().Set()

	led1 = &t.CCR2.U16
	led2 = &t.CCR1.U16
}

func pwmblink(led *mmio.U16, dly int) {
	var x uint
	var back bool
	for {
		led.Store(uint16(x))
		if x <= 1 {
			x = 1
			back = false
		}
		if x >= PWMmax {
			x = PWMmax
			back = true

		}
		if back {
			x /= 2
		} else {
			x *= 2
		}
		delay.Millisec(dly)
	}
}

func main() {
	go pwmblink(led1, 53)
	pwmblink(led2, 89)
}

//emgo:const
//c:__attribute__((section(".ISRs")))
var ISRs = [...]func(){
	irq.RTCAlarm: rtc.ISR,
}
