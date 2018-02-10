package main

import (
	"github.com/peter-mueller/guenztal-wasserspender/app/driver"
	"github.com/peter-mueller/guenztal-wasserspender/timer"
	"log"
	"github.com/peter-mueller/guenztal-wasserspender/money"
	moneyDriver "github.com/peter-mueller/guenztal-wasserspender/money/driver"
	"github.com/peter-mueller/guenztal-wasserspender/app/rest"
	"github.com/peter-mueller/guenztal-wasserspender/valve/control"
	"github.com/peter-mueller/guenztal-wasserspender/role"
)

var (


	coinTimer = &timer.Timer{}


	payLog = moneyDriver.NewFilePayLog()
	payer   = money.NewPayer(coinTimer, payLog)
	roleProvider = role.NewProvider(coinTimer)

	valves  = driver.NewValveStorage()
	coinAcceptor = driver.NewCoinAcceptor(payer)

	manager = control.NewController(valves, roleProvider)



	server = rest.NewServer(manager, coinTimer,payLog)
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
