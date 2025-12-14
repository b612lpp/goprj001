package storage

import "github.com/b612lpp/goprj001/metainf"

type IMDB struct {
	GasTable    []metainf.DataGas
	EnergyTable []metainf.DataEnergy
}

func NewDB() *IMDB {
	return &IMDB{}

}

func (db *IMDB) AddGas(gt metainf.DataGas) error {
	db.GasTable = append(db.GasTable, gt)
	return nil
}

func (db *IMDB) AddEnergy(gt metainf.DataEnergy) error {
	db.EnergyTable = append(db.EnergyTable, gt)
	return nil
}

type DataWriter interface {
	AddGas(metainf.DataGas) error
	AddEnergy(metainf.DataEnergy) error
}
