package data

import "fmt"

type DataBase struct {
	g []DataGas
	e []DataEnergy
}

func NewDb() *DataBase {
	d := DataBase{}
	return &d

}

func (db *DataBase) AddInfo(a any) {
	switch v := a.(type) {
	case DataGas:
		db.g = append(db.g, v)
		fmt.Println(db.g)
	case DataEnergy:
		db.e = append(db.e, v)
		fmt.Println(db.e)

	}

}
