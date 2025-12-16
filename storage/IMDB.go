// InMemory DataBase. Хранилище данных
package storage

import "github.com/b612lpp/goprj001/metainf"

type IMDB struct {
	GasTable    []metainf.DataGas    //таблица для хранения данных по газу в формате заданной структуры
	EnergyTable []metainf.DataEnergy //таблица для хранения данных по электричеству в формате заданной структуры
	Title       string
}

//создание нового экземпляра. создается на уровне создания сервера
func NewDB() *IMDB {
	return &IMDB{Title: "inmemory DB"}

}

//метод добавления данных соответствующий интерфейсу. принимает на вход структуру заполненную на этапе бизнес логики в пакете application
func (db *IMDB) AddGas(gt metainf.DataGas) error {
	db.GasTable = append(db.GasTable, gt)
	return nil
}

//метод добавления данных соответствующий интерфейсу. принимает на вход структуру заполненную на этапе бизнес логики в пакете application
func (db *IMDB) AddEnergy(gt metainf.DataEnergy) error {
	db.EnergyTable = append(db.EnergyTable, gt)
	return nil
}

//возвращаем табличку газа
func (db *IMDB) ReadGas() ([]metainf.DataGas, error) {
	q := db.GasTable
	if q == nil {
		return q, metainf.ErrDBConn
	}
	return q, nil
}

//возвращаем табличку электричества
func (db *IMDB) ReadEnergy() ([]metainf.DataEnergy, error) {
	q := db.EnergyTable
	if q == nil {
		return q, metainf.ErrDBConn
	}
	return q, nil
}
