package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/peter-mueller/guenztal-wasserspender/valve/control"
)

type (
	Valve struct {
		Open bool
	}

	ValveResource struct {
		vc *control.Controller
	}
)

func NewValveResource(vc *control.Controller) *ValveResource {
	return &ValveResource{
		vc: vc,
	}
}

func (b *ValveResource) FindAll(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	err := json.NewEncoder(w).Encode(b.vc.Valves)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (b *ValveResource) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	name := params.ByName("name")
	var v Valve
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	switch name {
	case "cold":
		request := control.UpdateColdValveRequest{Open: v.Open}
		b.vc.UpdateColdValve(request)
		break
	case "warm":
		request := control.UpdateWarmValveRequest{Open: v.Open}
		b.vc.UpdateWarmValve(request)
		break
	case "osmose":
		request := control.UpdateOsmoseValveRequest{Open: v.Open}
		b.vc.UpdateOsmoseValve(request)
		break
	default:
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
