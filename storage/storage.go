package storage

type DB struct {
	GasTable    []DataGas
	EnergyTable []DataEnergy
}

func (db *DB) AddGas(gt DataGas) error {
	db.GasTable = append(db.GasTable, gt)
	return nil
}

func (db *DB) AddEnergy(gt DataEnergy) error {
	db.EnergyTable = append(db.EnergyTable, gt)
	return nil
}

type DataWriter interface {
	AddGas(DataGas) error
	AddEnergy(DataEnergy) error
}

type DataGas struct {
	Gv int
}

type DataEnergy struct {
	Day, Night, Summ int
}
