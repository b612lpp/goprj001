package server

import (
	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/storage"
)

type Server struct {
	port  string
	DB    *storage.IMDB
	Gasuc *application.GasDataCase
	Dw    application.DataWriter
}

func NewServerConf() Server {
	port := ":8081"
	db := storage.NewDB()
	gasuc := application.NewGasDataCase()

	return Server{port, db, gasuc}
}
