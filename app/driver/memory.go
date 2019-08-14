// +build !rpi

package driver

import (
	"time"

	"github.com/peter-mueller/guenztal-wasserspender/money"
	"github.com/peter-mueller/guenztal-wasserspender/valve"
)

type (
	CoinAcceptor struct {
	}

	Payer interface {
		Pay(Money money.Money)
	}

	Memory struct {
		High bool
	}
)

func NewCoinAcceptor(payer Payer) *CoinAcceptor {
	return &CoinAcceptor{}
}

func NewValveStorage() *valve.Storage {
	m := &valve.Storage{
		Cold:   valve.NewValve("cold", Memory{}),
		Warm:   valve.NewValve("warm", Memory{}),
		Osmose: valve.NewValve("osmose", Memory{}),
	}

	m.Warm.OpenDuration = time.Minute * 20
	m.Osmose.OpenDuration = time.Minute * 20
	return m
}

func (m Memory) HIGH() error {
	m.High = true
	return nil
}

func (m Memory) LOW() error {
	m.High = false
	return nil
}
