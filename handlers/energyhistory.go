package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/b612lpp/goprj001/application"
)

type EnergyHistoryHandlerFunc struct {
	Edc *application.EnergyDataCase
}

func NewEnergyHistoryHandlerFunc(edc *application.EnergyDataCase) *EnergyHistoryHandlerFunc {
	return &EnergyHistoryHandlerFunc{Edc: edc}
}

func (eh *EnergyHistoryHandlerFunc) GetAndFormJsonEn(w http.ResponseWriter, r *http.Request) {

	z, err := eh.Edc.EnergyDataGetter()
	slog.Info("из базы данные получили ")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(z)

}
