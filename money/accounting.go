package money

import (
	"time"
	"fmt"
)

type (
	PayLog struct {
		Time    time.Time
		Payment Money
	}

	Month struct {
		Month time.Month
		Year int
	}
)

func Sum(logs <-chan PayLog) (sum Money) {
	for log := range logs {
		sum += log.Payment
	}
	return sum
}

func SumPerMonth(logs <-chan PayLog) (map[Month] Money) {
	sums := make(map[Month] Money)
	for log := range logs {
		m := Month{Month: log.Time.Month(), Year: log.Time.Year()}
		sums[m] += log.Payment
	}
	return sums
}

func (m Month)  MarshalText() ([]byte, error) {
	date := fmt.Sprintf("%04d-%02d",m.Year , m.Month)
	return []byte(date),nil
}