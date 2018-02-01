package valve

import (
	"log"
	"time"
)

type (
	Valve struct {
		Name string
		OpenDuration time.Duration

		output Output
		timer *time.Timer
	}

	Storage struct {
		Warm Valve
		Cold Valve
		Osmose Valve
	}

	Output interface {
		HIGH() error
		LOW() error
	}
)

const (
	defaultOpenTime = time.Second * 20
)

func NewValve(name string, output Output) Valve {
	return Valve{output: output, Name: name, OpenDuration: defaultOpenTime}
}

func (v* Valve) Open() error {
	if v.timer != nil  {
		v.timer.Stop()
	}

	v.timer = time.AfterFunc(v.OpenDuration, func () {
		err := v.Close()
		if err != nil {
			log.Println(err)
		}
	})

	log.Printf("Openened Valve %s", v.Name)
	return v.output.HIGH()
}

func (v* Valve) Close() error {
	if v.timer != nil  {
		v.timer.Stop()
	}

	log.Printf("Closed Valve %s", v.Name)
	return v.output.LOW()
}