package led

type Lighter interface {
	TurnOn()
	TurnOff()
	Close() error
	IsOn() bool
}