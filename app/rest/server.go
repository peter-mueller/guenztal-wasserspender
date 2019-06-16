package rest

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/julienschmidt/httprouter"
	"github.com/peter-mueller/guenztal-wasserspender/timer"
	"github.com/peter-mueller/guenztal-wasserspender/valve/control"
)

type (
	Server struct {
		valves *ValveResource
		timer  *TimerResource
		money  *AccountingResource
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
	router.PUT("/api/v1/valves/:name", s.valves.Update)
	router.GET("/api/v1/valves/", s.valves.FindAll)
	router.GET("/api/v1/timer/", s.timer.Query)
	router.GET("/api/v1/accounting/", s.money.Query)

	router.NotFound = http.FileServer(http.Dir("web"))

	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return http.ListenAndServe(":8080", handlers.CORS(headersOk, originsOk, methodsOk)(router))
}
