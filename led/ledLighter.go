package led

import (
	"github.com/stianeikeland/go-rpio"
	"fmt"
)

type BedLighter struct {
	pin rpio.Pin
	isEnabled bool
}

func NewBedLighter(gpio int) (*BedLighter, error) {
	fmt.Println("Initializing the bed light")
	err := rpio.Open()
	if err != nil {
		return nil, err
	}
	light := new(BedLighter)
	light.pin = rpio.Pin(gpio)
	light.TurnOff()
	return light, nil
}

func (l *BedLighter) Close() error {
	fmt.Println("Closing the bed light")
	return rpio.Close()
}

// GetStatus returns a status of light, where true stands for lights being on and false for light being off.
func (l *BedLighter) IsOn() bool {
	return l.isEnabled
}

func (l *BedLighter) TurnOn() {
	fmt.Println("Turning lights on")
	l.pin.Output()
	l.pin.Low()
	l.isEnabled = true
}

func (l *BedLighter) TurnOff() {
	fmt.Println("Turning lights off")
	l.pin.Output()
	l.pin.High()
	l.isEnabled = false
}