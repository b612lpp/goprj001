package handlers

import (
	"errors"
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

// хэндлер. обработчик запроса по http. нет бизнес логики, только транспорт и передача данных в юз кейс
func (eh *EnergyHandler) ParseEnergyData(w http.ResponseWriter, r *http.Request) {
	var v metainf.DataEnergy
	if err := utils.ParseUserData(r, &v); err != nil {
		w.WriteHeader(400)
		return
	}

	err := eh.Edc.EnergyDataProcessor(v)
	if errors.Is(err, metainf.ErrDBConn) {
		w.WriteHeader(500)
		return
	}

	if errors.Is(err, metainf.ErrWrongData) {
		w.WriteHeader(400)
		return
	}
	if err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "данные приняты")
	}

}
