// Создаём роутер, прописываем пути
package server

import (
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/handlers"
)

func NewRouter(s *application.GasDataCase) *http.ServeMux {
	q := handlers.NewGasHandler(s) //создаём экземпляр. передаём созданную структуру бизнеслогики
	m := http.NewServeMux()
	m.HandleFunc("/sendgas", q.SendGas)

	return m
}
