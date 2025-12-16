// InMemory DataBase. Хранилище данных
package storage

import "github.com/b612lpp/goprj001/domain"

type IMDB struct {
	GasTable    []domain.DataGas    //таблица для хранения данных по газу в формате заданной структуры
	EnergyTable []domain.DataEnergy //таблица для хранения данных по электричеству в формате заданной структуры
	Title       string
}

//создание нового экземпляра. создается на уровне создания сервера
func NewDB() *IMDB {
	return &IMDB{Title: "inmemory DB"}

}

//метод добавления данных соответствующий интерфейсу. принимает на вход структуру заполненную на этапе бизнес логики в пакете application
func (db *IMDB) AddGas(gt domain.DataGas) error {
	db.GasTable = append(db.GasTable, gt)
	return nil
}

//метод добавления данных соответствующий интерфейсу. принимает на вход структуру заполненную на этапе бизнес логики в пакете application
func (db *IMDB) AddEnergy(gt domain.DataEnergy) error {
	db.EnergyTable = append(db.EnergyTable, gt)
	return nil
}

//возвращаем табличку газа
func (db *IMDB) ReadGas() ([]domain.DataGas, error) {
	q := db.GasTable
	if q == nil {
		return q, domain.ErrDBConn
	}
	return q, nil
}

//возвращаем табличку электричества
func (db *IMDB) ReadEnergy() ([]domain.DataEnergy, error) {
	q := db.EnergyTable
	if q == nil {
		return q, domain.ErrDBConn
	}
	return q, nil
}
