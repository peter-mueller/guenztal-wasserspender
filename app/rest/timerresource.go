package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/peter-mueller/guenztal-wasserspender/timer"
)

type (
	Timer struct {
		End time.Time
	}

	TimerResource struct {
		timer *timer.Timer
	}
)

func NewTimerResource(timer *timer.Timer) *TimerResource {
	return &TimerResource{
		timer: timer,
	}
}

func (b *TimerResource) Query(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	var t Timer
	t.End = b.timer.End

	err := json.NewEncoder(w).Encode(t)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
