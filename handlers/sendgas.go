package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/metainf"
)

type GasHandler struct {
	Gdc *application.GasDataCase
}

func NewGasHandler(gdc *application.GasDataCase) *GasHandler {
	return &GasHandler{Gdc: gdc}
}

func (gh *GasHandler) SendGas(w http.ResponseWriter, r *http.Request) {
	var v metainf.DataGas
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if gh.Gdc.AddGasRow(v) != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)

}
