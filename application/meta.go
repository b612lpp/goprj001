package application

import (
	"github.com/b612lpp/goprj001/metainf"
)

type UseCases struct {
	GasUc    *GasDataCase
	EnergyUc *EnergyDataCase
}

// Описываем интерфейс записи данных в хранилище. За интерефейсом может быть что угодно созданное на уровне создания сервера
type GasDataWriter interface {
	AddGas(metainf.DataGas) error
}

type EnergyDataWriter interface {
	AddEnergy(metainf.DataEnergy) error
}

type GasDataCase struct {
	Gdw GasDataWriter
}

type EnergyDataCase struct {
	Edw EnergyDataWriter
}
