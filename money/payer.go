package money

import (
	"time"
)

type (
	Payer struct {
		DurationAdder
	}

	DurationAdder interface {
		AddDuration(duration time.Duration)
	}
)

func NewPayer(adder DurationAdder) *Payer {
	return &Payer{
		DurationAdder: adder,
	}
}

func (p *Payer) Pay(m Money) {
	if m.Cents == 0 {
		return
	}

	d := time.Duration(m.Cents) * time.Second
	p.AddDuration(d)
}
