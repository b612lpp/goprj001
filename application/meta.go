// выглядит так что этот файл избыточен. уместен здесь только UseCases. Надо думать о разносе по application
package application

import (
	"github.com/b612lpp/goprj001/metainf"
)

// Структура хранения бизнес сценариев через которые проходят данные.
type UseCases struct {
	GasUc    *GasDataCase
	EnergyUc *EnergyDataCase
}

// Описываем интерфейс записи данных в хранилище. За интерефейсом может быть что угодно созданное на уровне создания сервера
type GasDataProvider interface {
	AddGas(metainf.DataGas) error
	ReadGas() ([]metainf.DataGas, error)
}

// Интерфейс для записи данных по электричеству
type EnergyDataProvider interface {
	AddEnergy(metainf.DataEnergy) error
	ReadEnergy() ([]metainf.DataEnergy, error)
}

type GasDataCase struct {
	Provider GasDataProvider
}

type EnergyDataCase struct {
	Provider EnergyDataProvider
}
