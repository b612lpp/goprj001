// Хэндлер. Стык логики и инфраструктуры. Здесь мы обрабатываем запрос от клиента, проводим десереализацию. Заполняем структуру и отдаём блок бизнеслогики.
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

type GasHandler struct {
	UseCase *application.GasDataCase
}

// создаём новый экземпляр на этапе инициализации сервера
func NewGasHandlerFunc(gdc *application.GasDataCase) *GasHandler {
	return &GasHandler{UseCase: gdc}
}

// хэндлер. получает от клиента данные по http, адаптирует(десереализует) и отдает в бизнес логику
func (gh *GasHandler) ParseGasData(w http.ResponseWriter, r *http.Request) {
	var v domain.DataGas
	//парсим json
	if err := utils.ParseUserData(r, &v); err != nil {
		w.WriteHeader(400)
		return
	}

	//скармливаем структуру в бизнес логику юзкейса и возвращаем ошибки в канал
	err := gh.UseCase.GasDataCase(v)
	if errors.Is(err, domain.ErrDBConn) {
		w.WriteHeader(500)
		slog.Error("ошибка записи в БД")
		return
	}

	if errors.Is(err, domain.ErrWrongData) {
		w.WriteHeader(400)
		slog.Error("ошибка данных по газу")
		return
	}
	if err == nil {
		w.WriteHeader(http.StatusOK)
		slog.Info("данные газа успешно сохранены")
		fmt.Fprint(w, "данные приняты")

	}

}
