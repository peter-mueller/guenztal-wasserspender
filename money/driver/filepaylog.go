package driver

import (
	"github.com/peter-mueller/guenztal-wasserspender/money"
	"os"
	"encoding/json"
	"time"
	"log"
	"bufio"
)

type (
	FilePayLog struct {
		logFile *os.File
	}
)

func (f *FilePayLog) LogPay(m money.Money) {
	log := money.PayLog{
		Time:    time.Now(),
		Payment: m,
	}
	json.NewEncoder(f.logFile).Encode(log)
}

func (f *FilePayLog) FindAllLogs() <-chan money.PayLog {
	c := make(chan money.PayLog)

	go func() {
		defer close(c)

		f, err := os.Open("paylog.log")
		if err != nil {
			log.Println(err)
			close(c)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var l money.PayLog
			err := json.Unmarshal(scanner.Bytes(), &l)
			if err != nil {
				log.Println(err)
			}
			c <- l
		}
		if err := scanner.Err(); err != nil {
			log.Println(err)
		}
	}()

	return c
}

func NewFilePayLog() *FilePayLog {
	f, err := os.OpenFile("paylog.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	return &FilePayLog{logFile: f}
}
func (f *FilePayLog) Close() {
	f.logFile.Close()
}
