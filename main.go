package main

import (
	"log/slog"
	"net/http"

	"github.com/b612lpp/goprj001/server"
	"github.com/b612lpp/goprj001/server/midlware"
)

func main() {

	s := server.NewServerConf() //
	r := midlware.MWCors(server.NewRouter(&s.Uc).ServeHTTP)
	slog.SetDefault(s.Logger)
	slog.Info("Сервер запущен. Порт сервера " + s.Port + " подключени к экземпляру БД  " + s.DB.Title)
	http.ListenAndServe(s.Port, r)

}

/**
хэндлер фанк это функциональный тип который на вход принимает путь(/blabla) и функцию и реализует интерфейс Handler с методом ServeHTTP
listenandserv принимает как обработчик объект соответствующий интерфейсу Handler


Middleware должен оборачивать router, а не отдельные хэндлеры.
*/
