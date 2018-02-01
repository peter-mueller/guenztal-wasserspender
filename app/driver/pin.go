// +build rpi

package driver

import(
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi"
	"github.com/peter-mueller/guenztal-wasserspender/valve"
)
type (
	Pin struct {
		Number uint
	}
)
func init() {
	err := embd.InitGPIO()
	if err != nil {
		panic(err)
	}
}

func NewValveStorage() *valve.Storage {
	return &ValveFactory{
		Cold: valve.NewValve("cold", NewPin(1)),
		Warm: valve.NewValve("warm", NewPin(1)),
		Osmose: valve.NewValve("osmose", NewPin(1)),
	}
}

func NewPin(pin uint) Pin {
	err := embd.SetDirection(pin, embd.Out)
	if err != nil {
		panic(err)
	}
	return Pin{Number: pin}
}
func (p Pin) HIGH() error {
	return embd.DigitalWrite(p.Number, embd.High)
}

func (p Pin) LOW() error {
	return embd.DigitalWrite(p.Number, embd.Low)
}