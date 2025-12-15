// Создаём роутер, прописываем пути
package server

import (
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/handlers"
)

func NewRouter(uc *application.UseCases) *http.ServeMux {

	m := http.NewServeMux()
	m.HandleFunc("/sendgas", handlers.NewGasHandlerFunc(uc.GasUc).ParseGasData) //создаём экземпляр. передаём созданную структуру бизнеслогики
	m.HandleFunc("/sendenergy", handlers.NewEnergyHandlerFunc(uc.EnergyUc).ParseEnergyData)

	return m
}
