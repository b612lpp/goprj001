// выглядит так что этот файл избыточен. уместен здесь только UseCases. Надо думать о разносе по application
package application

import (
	"github.com/b612lpp/goprj001/domain"
)

// Структура хранения бизнес сценариев через которые проходят данные.
type UseCases struct {
	GasUc    *GasDataCase
	EnergyUc *EnergyDataCase
}

// Описываем интерфейс записи данных в хранилище. За интерефейсом может быть что угодно созданное на уровне создания сервера
type GasDataProvider interface {
	AddGas(domain.DataGas) error
	ReadGas() ([]domain.DataGas, error)
}

// Интерфейс для записи данных по электричеству
type EnergyDataProvider interface {
	AddEnergy(domain.DataEnergy) error
	ReadEnergy() ([]domain.DataEnergy, error)
}

type GasDataCase struct {
	Provider GasDataProvider
}

type EnergyDataCase struct {
	Provider EnergyDataProvider
}
