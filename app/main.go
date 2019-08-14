package main

import (
	"log"

	"github.com/peter-mueller/guenztal-wasserspender/app/driver"
	"github.com/peter-mueller/guenztal-wasserspender/app/rest"
	"github.com/peter-mueller/guenztal-wasserspender/money"
	moneyDriver "github.com/peter-mueller/guenztal-wasserspender/money/driver"
	"github.com/peter-mueller/guenztal-wasserspender/role"
	"github.com/peter-mueller/guenztal-wasserspender/timer"
	"github.com/peter-mueller/guenztal-wasserspender/valve/control"
)

var (
	valves = driver.NewValveStorage()

	coinTimer = &timer.Timer{
		OnEnd: func() {
			valves.Osmose.Close()
			valves.Warm.Close()
		},
	}

	payLog       = moneyDriver.NewFilePayLog()
	payer        = money.NewPayer(coinTimer, payLog)
	roleProvider = role.NewProvider(coinTimer)

	coinAcceptor = driver.NewCoinAcceptor(payer)

	manager = control.NewController(valves, roleProvider)

	server = rest.NewServer(manager, coinTimer, payLog)
)

func main() {
	defer payLog.Close()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	payer.Pay(money.Euro)
	err := server.Start()
	if err != nil {
		panic(err)
	}
}
