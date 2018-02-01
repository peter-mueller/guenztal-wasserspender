package rest

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/peter-mueller/guenztal-wasserspender/valve/control"
	"github.com/peter-mueller/guenztal-wasserspender/timer"
)

type (
	Server struct {
		valves *ValveResource
		timer *TimerResource
	}
)

func NewServer(vc *control.Controller, timer *timer.Timer) *Server {

	return &Server{
		valves: NewValveResource(vc),
		timer: NewTimerResource(timer),
	}
}

func (s Server) Start() {
	router := httprouter.New()
	router.ServeFiles("/app/*filepath", http.Dir("web"))
	router.PUT("/api/v1/buttons/:name", s.valves.Update)
	router.GET("/api/v1/timer/", s.timer.Query)
	http.ListenAndServe(":8080", router)
}
