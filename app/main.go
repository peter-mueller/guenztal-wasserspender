package main

import (
	"github.com/peter-mueller/guenztal-wasserspender/app/driver"
	"github.com/peter-mueller/guenztal-wasserspender/timer"
	"log"
	"github.com/peter-mueller/guenztal-wasserspender/money"
	"github.com/peter-mueller/guenztal-wasserspender/app/rest"
	"github.com/peter-mueller/guenztal-wasserspender/valve/control"
	"github.com/peter-mueller/guenztal-wasserspender/role"
)

var (
	valves  = driver.NewValveStorage()

	coinTimer = &timer.Timer{}
	payer   = money.NewPayer(coinTimer)
	roleProvider = role.NewProvider(coinTimer)
	manager = control.NewController(valves, roleProvider)

	server = rest.NewServer(manager, coinTimer)
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	payer.Pay(money.Money{money.Cent * 20})
	server.Start()
}
