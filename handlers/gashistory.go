package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/b612lpp/goprj001/application"
)

type GasHistoryHandlerFunc struct {
	UseCase *application.GasDataCase
}

func NewGasHistoryHandlerFunc(gdc *application.GasDataCase) *GasHistoryHandlerFunc {
	return &GasHistoryHandlerFunc{UseCase: gdc}
}

func (hf *GasHistoryHandlerFunc) GasHistory(w http.ResponseWriter, r *http.Request) {
	gd, err := hf.UseCase.GasHistory()
	if err != nil {
		slog.Error("ошибка получения истории газа")
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(gd)

}

/* с чем должен работать хэндлер:
на вход получает запрос от сервера по http дай мне историю, он не знает ничего про историю, но знает у кого попросить данные.
Идёт в юз кейс просит бизнес данные, десереализует их и возвращает в сервер.



*/
