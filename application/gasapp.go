// Блок бизнес логики. Сюда приходит уже готовая структура которую надо проверить на валидность значений
// и можно совершать действия над данными перед записью в хранилище
package application

import "github.com/b612lpp/goprj001/metainf"

//Описываем интерфейс записи данных в хранилище. За интерефейсом может быть что угодно созданное на уровне создания сервера
type DataWriter interface {
	AddGas(metainf.DataGas) error
}
type GasDataCase struct {
	Dw DataWriter
}

//Создается экземпляр объекта исполняющего бизнеслогику
func NewGasDataCase(dw DataWriter) *GasDataCase {

	q := GasDataCase{Dw: dw}
	return &q
}

//Вызов интерфейса записи данных в хранилище
func (gdc *GasDataCase) AddGasRow(dg metainf.DataGas) error {
	err := gdc.Dw.AddGas(dg)
	return err
}
