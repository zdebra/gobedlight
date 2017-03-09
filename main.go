package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

func main()  {
	fmt.Println("Start awesome light")
	err := rpio.Open()
	if err != nil {
		return
	}
	pin := rpio.Pin(18)

	//pin.Output()       // Output mode
	//pin.High()         // Set pin High

	pin.PullUp()
	time.Sleep(2000 * time.Millisecond)
	pin.Low()          // Set pin Low
	fmt.Println("turning it off")

	time.Sleep(2000 * time.Millisecond)
	//rpio.Close()
	pin.PullOff()

}