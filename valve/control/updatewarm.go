package control

import "github.com/peter-mueller/guenztal-wasserspender/role"

type (
	UpdateWarmValveRequest struct {
		Open bool
	}
)

func (c Controller) UpdateWarmValve(request UpdateWarmValveRequest) error {
	if !request.Open {

		return c.Valves.Warm.Close()
	}

	if !c.RoleProvider.IsPremium() {
		return role.ErrPremiumRequired
	}
	return c.Valves.Warm.Open()
}
