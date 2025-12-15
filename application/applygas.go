// Блок бизнес логики. Сюда приходит уже готовая структура которую надо проверить на валидность значений
// и можно совершать действия над данными перед записью в хранилище
package application

import (
	"github.com/b612lpp/goprj001/metainf"
)

// Создается экземпляр объекта исполняющего бизнеслогику
func NewGasDataCase(dw GasDataWriter) *GasDataCase {

	return &GasDataCase{Gdw: dw}

}

// Вызов интерфейса записи данных в хранилище
func (gdc *GasDataCase) GasDataProcessor(dg metainf.DataGas) error {
	if dg.Value < 0 {

		return metainf.ErrWrongData
	}
	if err := gdc.Gdw.AddGas(dg); err != nil {
		return metainf.ErrDBConn
	}
	return nil
}
