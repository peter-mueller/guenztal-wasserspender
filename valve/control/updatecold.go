package control

type (
	UpdateColdValveRequest struct {
		Open bool
	}
)

func (c Controller) UpdateColdValve(request UpdateColdValveRequest) error {
	if !request.Open {
		return c.Valves.Cold.Close()
	}

	return c.Valves.Cold.Open()
}
