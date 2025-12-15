// Параметры уровня сервера, глобальные сущности БД и use case для данных
package server

import (
	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/storage"
)

type Server struct {
	Port  string
	DB    *storage.IMDB
	Gasuc *application.GasDataCase
}

func NewServerConf() Server {

	port := ":8081"
	db := storage.NewDB()
	gasuc := application.NewGasDataCase(db)

	return Server{port, db, gasuc}
}
