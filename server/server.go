// Параметры уровня сервера, глобальные сущности БД и use case для данных
package server

import (
	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/storage"
)

type Server struct {
	Port string
	DB   *storage.IMDB
	//Gasuc *application.GasDataCase
	Guc *application.GasDataCase
	Euc *application.EnergyDataCase
}

func NewServerConf() Server {

	port := ":8081"
	db := storage.NewDB()
	guc := application.NewGasDataCase(db)
	euc := application.NewEnergyDataCase(db)

	return Server{port, db, guc, euc}
}
