// Блок бизнес логики. Сюда приходит уже готовая структура которую надо проверить на валидность значений
// и можно совершать действия над данными перед записью в хранилище
package application

import (
	"time"

	"github.com/b612lpp/goprj001/domain"
)

// Создается экземпляр объекта исполняющего бизнеслогику
func NewGasDataCase(dp GasDataProvider) *GasDataCase {

	return &GasDataCase{Provider: dp}

}

// Прогон данных п о бизнесс сценарию
func (gdc *GasDataCase) GasDataCase(dg domain.DataGas) error {
	if dg.Value < 0 {

		return domain.ErrWrongData
	}
	dg.Time = time.Now()
	//Запись данных в хранилище через интерфейс
	if err := gdc.Provider.AddGas(dg); err != nil {
		return domain.ErrDBConn
	}
	return nil
}

func (gdc *GasDataCase) GasHistory() ([]domain.DataGas, error) {

	q, err := gdc.Provider.ReadGas()
	if err != nil {
		return nil, domain.ErrDBRead
	}

	return q, nil
}
