package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	ti := &Timer{}

	if ti.IsSet() {
		t.Error("timer is set")
	}
	if ti.Remaining() > 0 {
		t.Error("some time is remaining")
	}

	ti.AddDuration(time.Millisecond)
	if !ti.IsSet() {
		t.Error("timer is not set")
	}
	if ti.Remaining() == 0 {
		t.Error("no time remaining")
	}

	if ti.Remaining() == 0 {
		t.Error("no time remaining")
	}

	time.Sleep(ti.Remaining())

	if ti.IsSet() {
		t.Error("timer is set")
	}
}
