package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/b612lpp/goprj001/application"
)

type EnergyHistoryHandlerFunc struct {
	UseCase *application.EnergyDataCase
}

func NewEnergyHistoryHandlerFunc(edc *application.EnergyDataCase) *EnergyHistoryHandlerFunc {
	return &EnergyHistoryHandlerFunc{UseCase: edc}
}

func (eh *EnergyHistoryHandlerFunc) EnergyHistory(w http.ResponseWriter, r *http.Request) {

	z, err := eh.UseCase.EnergyHistory()
	slog.Info("из базы данные получили ")
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(z)

}
