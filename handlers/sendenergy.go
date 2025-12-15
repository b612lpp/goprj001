package handlers

import (
	"errors"
	"fmt"
	"log/slog"
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
		slog.Error("некорректные данные по электричеству")
		return
	}

	err := eh.Edc.EnergyDataProcessor(v)
	if errors.Is(err, metainf.ErrDBConn) {
		w.WriteHeader(500)
		slog.Error("ошибка записи в БД")
		return
	}

	if errors.Is(err, metainf.ErrWrongData) {
		w.WriteHeader(400)
		slog.Error("некорректные данные по электричеству")
		return
	}
	if err == nil {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "данные приняты")
		slog.Info("данные по электричеству приняты")
	}

}
