package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/metainf"
)

type EnergyHandler struct {
	Edc *application.EnergyDataCase
}

func NewEnergyHandlerFunc(edc *application.EnergyDataCase) *EnergyHandler {
	return &EnergyHandler{Edc: edc}
}

func (eh *EnergyHandler) ParseEnergyData(w http.ResponseWriter, r *http.Request) {
	var v metainf.DataEnergy
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	eh.Edc.EnergyDataProcessor(v)
	w.WriteHeader(http.StatusOK)
}
