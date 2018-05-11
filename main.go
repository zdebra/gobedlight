package main

import (
	"log"

	"net/http"

	"flag"

	"github.com/zdebra/gobedlight/led"
)

const gpio = 18

func main() {

	pin := flag.Int("pin", gpio, "pin where the light is connected")
	flag.Parse()

	log.Printf("initializing the light, light is expected on pin %d", *pin)

	lighter, err := led.NewBedLighter(*pin)
	if err != nil {
		panic(err)
	}
	defer lighter.Close()

	srv := &server{lighter}
	http.HandleFunc("/toggle", srv.HandleToggle())
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

}
