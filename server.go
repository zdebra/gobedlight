package main

import (
	"net/http"

	"log"

	"errors"

	"github.com/zdebra/gobedlight/led"
)

type server struct {
	lighter led.Lighter
}

func (s *server) HandleToggle() http.HandlerFunc {
	log.Print("toggle handler registered")
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				var err error
				switch t := r.(type) {
				case string:
					err = errors.New(t)
				case error:
					err = t
				default:
					err = errors.New("Unknown error")
				}
				log.Printf("recovered from panic: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}()

		log.Print("toggle handler triggered")
		if s.lighter.IsOn() {
			log.Print("turning light off")
			s.lighter.TurnOff()
		} else {
			log.Print("turning light on")
			s.lighter.TurnOn()
		}
	}
}
