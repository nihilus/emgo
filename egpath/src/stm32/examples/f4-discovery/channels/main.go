package main

import (
	"delay"

	"stm32/hal/gpio"
	"stm32/hal/setup"
)

var LED *gpio.Port

const (
	Green  = 12
	Orange = 13
	Red    = 14
	Blue   = 15
)

func init() {
	setup.Performance168(8)

	gpio.D.EnableClock(false)

	LED = gpio.D
	LED.SetMode(Green, gpio.Out)
	LED.SetMode(Orange, gpio.Out)
	LED.SetMode(Red, gpio.Out)
	LED.SetMode(Blue, gpio.Out)
}

func toggle(led, d int) {
	LED.SetPin(led)
	delay.Millisec(d)
	LED.ClearPin(led)
	delay.Millisec(d)
}

const dly = 100

func blink(color <-chan int, end chan<- struct{}) {
	for {
		led, ok := <-color
		if !ok {
			end <- struct{}{}
			return
		}
		toggle(led, dly*10)
	}
}

func main() {
	color := make(chan int, 10)
	end := make(chan struct{}, 3)

	// Consumers
	go blink(color, end)
	go blink(color, end)
	go blink(color, end)

	// Producer
	for i := 0; i < 10; i++ {
		color <- Red
		toggle(Orange, dly)
		color <- Blue
		toggle(Orange, dly)
		color <- Green
		toggle(Orange, dly)
	}
	close(color)

	// Wait for consumers.
	<-end
	<-end
	<-end
	for {
		toggle(Orange, dly)
	}
}
