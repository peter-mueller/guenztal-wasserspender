package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/peter-mueller/guenztal-wasserspender/money"
)

type (
	AccountingResource struct {
		logger PayLogger
	}

	PayLogger interface {
		FindAllLogs() <-chan money.PayLog
	}
)

func NewAccountingResource(logger PayLogger) *AccountingResource {
	return &AccountingResource{
		logger: logger,
	}
}

func (b *AccountingResource) Query(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	monthly := r.FormValue("groupBy")
	log.Println(monthly)

	switch monthly {
	case "monthly":
		c := b.logger.FindAllLogs()
		err := json.NewEncoder(w).Encode(money.SumPerMonth(c))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		break
	default:
		c := b.logger.FindAllLogs()
		err := json.NewEncoder(w).Encode(money.Sum(c))
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		break
	}
}
