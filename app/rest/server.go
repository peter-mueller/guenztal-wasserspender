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
		money *AccountingResource
	}
)

func NewServer(vc *control.Controller, timer *timer.Timer, logger PayLogger) *Server {

	return &Server{
		valves: NewValveResource(vc),
		timer:  NewTimerResource(timer),
		money:  NewAccountingResource(logger),
	}
}

func (s Server) Start() error {
	router := httprouter.New()
	router.ServeFiles("/app/*filepath", http.Dir("web"))
	router.PUT("/api/v1/buttons/:name", s.valves.Update)
	router.GET("/api/v1/timer/", s.timer.Query)
	router.GET("/api/v1/accounting/", s.money.Query)
	return http.ListenAndServe(":8080", router)
}
