// Создаём роутер, прописываем пути
package server

import (
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/handlers"
)

// Тут сложно. Через конструктор создаем новый обработчик для урла и обозначаем зависимость с юз кейсом который создается на уровне сервера
func NewRouter(uc *application.UseCases) *http.ServeMux {

	m := http.NewServeMux()
	m.HandleFunc("/sendgas", handlers.NewGasHandlerFunc(uc.GasUc).ParseGasData) //создаём экземпляр. передаём созданную структуру бизнеслогики
	m.HandleFunc("/sendenergy", handlers.NewEnergyHandlerFunc(uc.EnergyUc).ParseEnergyData)
	m.HandleFunc("/history/gas", handlers.NewGasHistoryHandlerFunc(uc.GasUc).GetAndFormJson)
	m.HandleFunc("/history/energy", handlers.NewEnergyHistoryHandlerFunc(uc.EnergyUc).GetAndFormJsonEn)

	return m
}
