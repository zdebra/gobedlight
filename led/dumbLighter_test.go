package led

import "testing"

func TestDumbLighter_Close(t *testing.T) {
	light,_ := NewDumbLighter()
	err := light.Close()
	if err != nil {
		t.Error("Closing the light gives an error.")
	}

}

func TestDumbLighter_IsOn(t *testing.T) {
	light,_ := NewDumbLighter()
	if light.IsOn() {
		t.Error("Lighter should be initialized as turned off.")
	}

	light.status = true
	if !light.IsOn() {
		t.Error("Light should be on")
	}
}

func TestDumbLighter_TurnOff(t *testing.T) {
	light,_ := NewDumbLighter()
	light.TurnOff()
	t.Log("something nice to log")
	if light.status {
		t.Fail()
	}
}

func TestDumbLighter_TurnOn(t *testing.T) {
	light,_ := NewDumbLighter()
	light.TurnOn()
	if !light.status {
		t.Fail()
	}
}


