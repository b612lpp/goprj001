// Хэндлер. Стык логики и инфраструктуры. Здесь мы обрабатываем запрос от клиента, проводим десереализацию. Заполняем структуру и отдаём блок бизнеслогики.
package handlers

import (
	"fmt"
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/metainf"
	"github.com/b612lpp/goprj001/utils"
)

type GasHandler struct {
	Gdc *application.GasDataCase
}

// создаём новый экземпляр на этапе инициализации сервера
func NewGasHandlerFunc(gdc *application.GasDataCase) *GasHandler {
	return &GasHandler{Gdc: gdc}
}

// механика десереализации
func (gh *GasHandler) ParseGasData(w http.ResponseWriter, r *http.Request) {
	var v metainf.DataGas
	if err := utils.ParseUserData(w, r, &v); err != nil {
		fmt.Fprint(w, err)
		return
	}

	if v.Value < 0 {
		w.WriteHeader(400)
		fmt.Fprint(w, "некорректные данные")
		return
	}

	if err := gh.Gdc.GasDataProcessor(v); err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "ошибка записи в базу")
		return
	} //передача куска десереализованной структуры
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "данные приняты")

}
