// +build !rpi

package driver

import (
	"github.com/peter-mueller/guenztal-wasserspender/valve"
	"github.com/peter-mueller/guenztal-wasserspender/money"
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
	return &valve.Storage{
		Cold:   valve.NewValve("cold", Memory{}),
		Warm:   valve.NewValve("warm", Memory{}),
		Osmose: valve.NewValve("osmose", Memory{}),
	}
}

func (m Memory) HIGH() error {
	m.High = true
	return nil
}

func (m Memory) LOW() error {
	m.High = false
	return nil
}
