// +build rpi

package driver

import (
	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/host"
	"periph.io/x/periph/host/rpi"
	"github.com/peter-mueller/guenztal-wasserspender/valve"
	"github.com/peter-mueller/guenztal-wasserspender/money"
	"log"
)

type (
	Pin struct {
		Number uint
		e      gpio.PinOut
	}

	CoinAcceptor struct {
		e gpio.PinIn
	}

	Payer interface {
		Pay(Money money.Money)
	}
)

func init() {
	_, err := host.Init()
	if err != nil {
		panic(err)
	}
}

func NewCoinAcceptor(payer Payer) *CoinAcceptor {

	pin := rpi.P1_38
	pin.In(gpio.PullUp, gpio.FallingEdge)

	go func() {
		for {
			log.Println("start wait")
			pin.WaitForEdge(-1)
			log.Println("paying")
			payer.Pay(money.Cent)
		}
	}();
	return &CoinAcceptor{e: pin}
}

func NewValveStorage() *valve.Storage {
	return &valve.Storage{
		Cold:   valve.NewValve("cold", NewPin(rpi.P1_35)),
		Warm:   valve.NewValve("warm", NewPin(rpi.P1_36)),
		Osmose: valve.NewValve("osmose", NewPin(rpi.P1_37)),
	}
}

func NewPin(pin gpio.PinOut) Pin {
	err := pin.Out(gpio.Low)
	if err != nil {
		panic(err)
	}
	return Pin{e: pin}
}
func (p Pin) HIGH() error {
	return p.e.Out(gpio.High)
}

func (p Pin) LOW() error {
	return p.e.Out(gpio.Low)
}
