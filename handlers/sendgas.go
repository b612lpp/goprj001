// Хэндлер. Стык логики и инфраструктуры. Здесь мы обрабатываем запрос от клиента, проводим десереализацию. Заполняем структуру и отдаём блок бизнеслогики.
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/metainf"
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
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gh.Gdc.GasDataProcessor(v) //передача куска десереализованной структуры
	fmt.Println(v)

}
