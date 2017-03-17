package main

import (
	"time"
	"github.com/zdebra/gobedlight/led"
	"fmt"
)

const GPIO int = 18

func main()  {

	light, err := led.NewBedLighter(GPIO)
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