package led

import "fmt"

type DumbLighter struct {
	status bool
}

func NewDumbLighter() (*DumbLighter, error) {
	light := new(DumbLighter)
	return light, nil
}

func (d *DumbLighter) TurnOn() {
	fmt.Println("Turning on!")
	d.status = true
}

func (d *DumbLighter) TurnOff() {
	fmt.Println("Turning off!")
	d.status = false
}

func (d *DumbLighter) Close() error {
	fmt.Println("Closing a dumb lighter.")
	return nil
}

func (d *DumbLighter) IsOn() bool {
	return d.status
}
