package main

type DB struct {
	GasTable    []DataGas
	EnergyTable []DataEnergy
}

func (db *DB) AddGas(gt GasTable) error {
	db = append(db.GasTable, gt)
}

func (db *DB) AddEnergy(gt EnergyTable) error {
	db = append(db.EnergyTable, gt)
}

type DataWriter interface {
	AddGas(DataGas) error
	AddEnergy(DataEnergy) error
}
