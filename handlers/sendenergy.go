package handlers

import (
	"fmt"
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/metainf"
	"github.com/b612lpp/goprj001/utils"
)

type EnergyHandler struct {
	Edc *application.EnergyDataCase
}

func NewEnergyHandlerFunc(edc *application.EnergyDataCase) *EnergyHandler {
	return &EnergyHandler{Edc: edc}
}

func (eh *EnergyHandler) ParseEnergyData(w http.ResponseWriter, r *http.Request) {
	var v metainf.DataEnergy
	if err := utils.ParseUserData(w, r, &v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)

	}
	if v.Day+v.Night != v.Summ || v.Day < 0 || v.Night < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "некорректные данные")
		return

	}
	eh.Edc.EnergyDataProcessor(v)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "данные приняты")
}
