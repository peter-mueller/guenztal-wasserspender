package control

import (
	"github.com/peter-mueller/guenztal-wasserspender/valve"
)

type (
	Controller struct {
		Valves *valve.Storage

		RoleProvider
	}

	RoleProvider interface {
		IsPremium() bool
	}
)

func NewController(valves *valve.Storage, r RoleProvider) *Controller {
	return &Controller{
		Valves: valves,
		RoleProvider:  r,
	}
}