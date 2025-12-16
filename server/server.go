// Параметры уровня сервера, глобальные сущности БД и use case для данных
package server

import (
	"log/slog"
	"os"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/storage"
)

type Server struct {
	Port   string
	Logger *slog.Logger
	DB     *storage.IMDB
	Uc     application.UseCases
}

func NewServerConf() Server {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	port := ":8081"
	db := storage.NewDB()
	//самая важная часть. определяется связь юз кейса и интерфейса хранилища. важно db реализует интерфейс writer
	x := application.UseCases{GasUc: application.NewGasDataCase(db), EnergyUc: application.NewEnergyDataCase(db)}
	return Server{port, logger, db, x}
}
