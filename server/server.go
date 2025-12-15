// Параметры уровня сервера, глобальные сущности БД и use case для данных
package server

import (
	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/storage"
)

type Server struct {
	Port string
	DB   *storage.IMDB
	Uc   application.UseCases
}

func NewServerConf() Server {

	port := ":8081"
	db := storage.NewDB()
	x := application.UseCases{GasUc: application.NewGasDataCase(db), EnergyUc: application.NewEnergyDataCase(db)}
	return Server{port, db, x}
}
