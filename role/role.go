package role

import "errors"

type(
	Role struct {	}

	Provider struct {
		timer Timer
	}

	Timer interface {
		IsSet() bool
	}
)

var (
	ErrPremiumRequired = errors.New("Premium status is required")
)

func NewProvider(timer Timer) *Provider {
	return &Provider{timer}
}

func (r *Provider) IsPremium() bool {
	return r.timer.IsSet()
}