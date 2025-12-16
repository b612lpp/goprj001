package handlers

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/domain"
	"github.com/b612lpp/goprj001/utils"
)

type EnergyHandler struct {
	UseCase *application.EnergyDataCase
}

func NewEnergyHandlerFunc(edc *application.EnergyDataCase) *EnergyHandler {
	return &EnergyHandler{UseCase: edc}
}

// хэндлер. обработчик запроса по http. нет бизнес логики, только транспорт и передача данных в юз кейс
func (eh *EnergyHandler) ParseEnergyData(w http.ResponseWriter, r *http.Request) {
	var v domain.DataEnergy
	//запускаем парсер json
	if err := utils.ParseUserData(r, &v); err != nil {
		w.WriteHeader(400)
		slog.Error("некорректные данные по электричеству")
		return
	}
	//скармливаем структурку в бизнесс логику и возвращаем ошибки в канал
	err := eh.UseCase.EnergyDataCase(v)
	if errors.Is(err, domain.ErrDBConn) {
		w.WriteHeader(500)
		slog.Error("ошибка записи в БД")
		return
	}

	if errors.Is(err, domain.ErrWrongData) {
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
