package money

import (
	"time"
)

type (
	Payer struct {
		DurationAdder
		PayLogger
	}

	PayLogger interface {
		LogPay(money Money)
	}

	DurationAdder interface {
		AddDuration(duration time.Duration)
	}
)

func NewPayer(adder DurationAdder, logger PayLogger) *Payer {
	return &Payer{
		DurationAdder: adder,
		PayLogger: logger,
	}
}

func (p *Payer) Pay(m Money) {
	if m == 0 {
		return
	}

	p.LogPay(m)

	d := time.Duration(m) * time.Second
	p.AddDuration(d)
}
