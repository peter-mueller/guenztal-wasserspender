package valve

import (
	"testing"
	"time"
)

type mockOutput struct{}

func (mockOutput) HIGH() error {
	return nil
}

func (mockOutput) LOW() error {
	return nil
}

func TestValve_Open(t *testing.T) {
	v1 := NewValve("warm", mockOutput{})
	v1.OpenDuration = time.Millisecond * 2

	v1.Open()
	time.Sleep(time.Millisecond)
	v1.Close()
}

func TestValve_Open2(t *testing.T) {
	v1 := NewValve("warm", mockOutput{})
	v1.OpenDuration = time.Millisecond

	v1.Open()
	time.Sleep(time.Millisecond * 5)
	v1.Close()
}

func TestValve_Open3(t *testing.T) {
	v1 := NewValve("warm", mockOutput{})
	v1.OpenDuration = time.Millisecond

	v1.Open()
	time.Sleep(time.Millisecond * 5)
}

func TestValve_Close(t *testing.T) {
	v1 := NewValve("warm", mockOutput{})
	v1.Close()
}
