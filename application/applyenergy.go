package application

import (
	"time"

	"github.com/b612lpp/goprj001/metainf"
)

// Создаёт экземпляр объекта в котором данные по электричеству обрабатываются по бизнесс сценарию и сохраняются в БД
func NewEnergyDataCase(dp EnergyDataProvider) *EnergyDataCase {
	return &EnergyDataCase{Provider: dp}
}

// Бизнес сценарий. В данном случае проверка данных по шаблону бизнеса, а не типов.
func (edc *EnergyDataCase) EnergyDataCase(ed metainf.DataEnergy) error {
	if ed.Day < 0 || ed.Night < 0 || ed.Day+ed.Night != ed.Summ {
		return metainf.ErrWrongData
	}
	ed.Time = time.Now()
	//запуск метода записи в БД через интерфейс
	if err := edc.Provider.AddEnergy(ed); err != nil {
		return metainf.ErrDBConn
	}
	return nil
}

func (edc *EnergyDataCase) EnergyHistory() ([]metainf.DataEnergy, error) {

	q, err := edc.Provider.ReadEnergy()
	if err != nil {
		return nil, metainf.ErrDBRead
	}

	return q, nil
}
