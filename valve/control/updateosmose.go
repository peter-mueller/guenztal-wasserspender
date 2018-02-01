package control

import "github.com/peter-mueller/guenztal-wasserspender/role"

type(
	UpdateOsmoseValveRequest struct {
		Open bool
	}
)

func (c Controller) UpdateOsmoseValve(request UpdateOsmoseValveRequest) error {
	if !request.Open {
		return c.Valves.Osmose.Close()
	}

	if !c.RoleProvider.IsPremium() {
		return role.ErrPremiumRequired
	}
	return c.Valves.Osmose.Open()
}
