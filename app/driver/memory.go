// +build !rpi

package driver

import "github.com/peter-mueller/guenztal-wasserspender/valve"

type (
	Memory struct {
		High bool
	}
)

func NewValveStorage() *valve.Storage {
	return &valve.Storage{
		Cold: valve.NewValve("cold", Memory{}),
		Warm: valve.NewValve("warm", Memory{}),
		Osmose: valve.NewValve("osmose", Memory{}),
	}
}

func (m Memory) HIGH() error {
	m.High = true
	return nil
}

func (m Memory) LOW()error {
	m.High = false
	return nil
}
