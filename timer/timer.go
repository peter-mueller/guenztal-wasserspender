package timer

import "time"

type (
	Timer struct {
		End   time.Time
		OnEnd func()
		timer *time.Timer
	}
)

func (t *Timer) AddDuration(duration time.Duration) {
	if !t.IsSet() {
		t.End = time.Now()
	}
	t.End = t.End.Add(duration)

	if t.timer != nil {
		t.timer.Stop()
	}
	if t.OnEnd != nil {
		t.timer = time.AfterFunc(t.Remaining(), t.OnEnd)
	}
}

func (t *Timer) Remaining() time.Duration {
	if !t.IsSet() {
		return 0
	}
	return time.Until(t.End)
}

func (t *Timer) IsSet() bool {
	if t.End.IsZero() {
		return false
	}
	return t.End.After(time.Now())
}
