package server

import (
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/handlers"
)

type MainHandler struct {
	mux *http.ServeMux
}

func NewMainHandler() MainHandler {

	m := http.NewServeMux()
	m.HandleFunc("/sendgas", handlers.NewGasHandler(application.NewGasDataCase()).SendGas)

	return MainHandler{mux: m}
}
