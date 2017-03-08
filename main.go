package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	"time"
)

func main()  {
	fmt.Println("Start awesome light")
	err := rpio.Open()
	pin := rpio.Pin(12)

	pin.Output()       // Output mode
	pin.High()         // Set pin High
	time.Sleep(1000 * time.Millisecond)
	pin.Low()          // Set pin Low
	pin.Toggle()       // Toggle pin (Low -> High -> Low)

	rpio.Close()
	

}