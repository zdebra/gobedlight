package led

import "fmt"

type DumbLighter struct {
	status bool
}

func (d DumbLighter) Initialize(gpio int) error {
	fmt.Println("Initializing a dumb lighter.")
	d.status = false
	return nil
}

func (d DumbLighter) TurnOn() {
	fmt.Println("Turning on!")
	d.status = true
}

func (d DumbLighter) TurnOff() {
	fmt.Println("Turning off!")
	d.status = false
}

func (d DumbLighter) Close() error {
	fmt.Println("Closing a dumb lighter.")
	return nil
}

func (d DumbLighter) IsOn() bool {
	return d.status
}
