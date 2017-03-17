package main

import (
	"github.com/zdebra/gobedlight/led"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hc"
	"log"
)

const GPIO int = 18

func main()  {

	log.Println("Initializing the light.")

	light, err := led.NewBedLighter(GPIO)
	if err != nil {
		return
	}
	defer light.Close()

	log.Println("Initializing homekit server.")

	info := accessory.Info{
		Name: "Lamp",
	}
	acc := accessory.NewSwitch(info)
	config := hc.Config{Pin: "00102003"}
	t, err := hc.NewIPTransport(config,acc.Accessory)
	if err != nil {
		log.Panic(err)
	}


	// Log to console when client (e.g. iOS app) changes the value of the on characteristic
	acc.Switch.On.OnValueRemoteUpdate(func(on bool) {
		if on == true {
			log.Println("Client changed switch to on")
			light.TurnOn()
		} else {
			log.Println("Client changed switch to off")
			light.TurnOff()
		}
	})

	hc.OnTermination(func() {
		t.Stop()
	})

	t.Start()


}