package main

import (
	"net/http"

	"github.com/b612lpp/goprj001/server"
)

func main() {
	h := server.NewMySrv(":8082")
	http.ListenAndServe(h.Port, &h.Mux)
}

/**
хэндлер фанк это функциональный тип который на вход принимает путь(/blabla) и функцию и реализует интерфейс Handler с методом ServeHTTP
listenandserv принимает как обработчик объект соответствующий интерфейсу Handler

*/
