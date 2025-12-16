package application

import (
	"github.com/b612lpp/goprj001/metainf"
)

// Создаёт экземпляр объекта в котором данные по электричеству обрабатываются по бизнесс сценарию и сохраняются в БД
func NewEnergyDataCase(dp EnergyDataProvider) *EnergyDataCase {
	return &EnergyDataCase{Edp: dp}
}

// Бизнес сценарий. В данном случае проверка данных по шаблону бизнеса, а не типов.
func (edc *EnergyDataCase) EnergyDataProcessor(ed metainf.DataEnergy) error {
	if ed.Day < 0 || ed.Night < 0 || ed.Day+ed.Night != ed.Summ {
		return metainf.ErrWrongData
	}
	//запуск метода записи в БД через интерфейс
	if err := edc.Edp.AddEnergy(ed); err != nil {
		return metainf.ErrDBConn
	}
	return nil
}

func (edc *EnergyDataCase) EnergyDataGetter() ([]metainf.DataEnergy, error) {

	q, err := edc.Edp.ReadEnergy()
	if err != nil {
		return nil, metainf.ErrDBRead
	}

	return q, nil
}
