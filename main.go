package main

import (
	"time"
	"github.com/zdebra/gobedlight/led"
	"fmt"
)

const GPIO int = 18

func main()  {

	/*
	fmt.Println("Start awesome light")
	err := rpio.Open()
	if err != nil {
		return
	}
	defer rpio.Close()

	pin := rpio.Pin(GPIO)

	pin.Output()       // Output mode
	pin.Low()         // turn the lights on
	time.Sleep(2000 * time.Millisecond)
	pin.High()          // turn the light off
	fmt.Println("turning it off")

	time.Sleep(2000 * time.Millisecond)
	rpio.Close()
	*/

	light, err := led.New(GPIO)
	if err != nil {
		return
	}
	defer light.Close()

	fmt.Println(light.IsOn())

	time.Sleep(2000 * time.Millisecond)

	light.TurnOn()
	fmt.Println(light.IsOn())

	time.Sleep(2000 * time.Millisecond)

	light.TurnOff()
	time.Sleep(2000 * time.Millisecond)

	fmt.Println(light.IsOn())

}