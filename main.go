package main

import (
	"net/http"

	"github.com/b612lpp/goprj001/server"
)

func main() {

	s := server.NewServerConf() //
	r := server.NewRouter(s.Gasuc)
	http.ListenAndServe(s.Port, r)
}

/**
хэндлер фанк это функциональный тип который на вход принимает путь(/blabla) и функцию и реализует интерфейс Handler с методом ServeHTTP
listenandserv принимает как обработчик объект соответствующий интерфейсу Handler

*/
