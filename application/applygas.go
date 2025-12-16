// Блок бизнес логики. Сюда приходит уже готовая структура которую надо проверить на валидность значений
// и можно совершать действия над данными перед записью в хранилище
package application

import (
	"time"

	"github.com/b612lpp/goprj001/metainf"
)

// Создается экземпляр объекта исполняющего бизнеслогику
func NewGasDataCase(dp GasDataProvider) *GasDataCase {

	return &GasDataCase{Gdp: dp}

}

// Прогон данных п о бизнесс сценарию
func (gdc *GasDataCase) GasDataProcessor(dg metainf.DataGas) error {
	if dg.Value < 0 {

		return metainf.ErrWrongData
	}
	dg.Time = time.Now()
	//Запись данных в хранилище через интерфейс
	if err := gdc.Gdp.AddGas(dg); err != nil {
		return metainf.ErrDBConn
	}
	return nil
}

func (gdc *GasDataCase) GasDataGetter() ([]metainf.DataGas, error) {

	q, err := gdc.Gdp.ReadGas()
	if err != nil {
		return nil, metainf.ErrDBRead
	}

	return q, nil
}
