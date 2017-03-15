package led

type Lighter interface {
	Initialize(gpio int) error
	TurnOn()
	TurnOff()
	Close() error
	IsOn() bool
}