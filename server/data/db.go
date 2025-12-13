package data

import "fmt"

type DataBase struct {
	G []DataGas    `json:"datagas"`
	E []DataEnergy `json:"dataenergy"`
}

func NewDb() *DataBase {
	d := DataBase{}
	return &d

}

func (db *DataBase) AddInfo(a any) {
	switch v := a.(type) {
	case DataGas:
		db.G = append(db.G, v)
		fmt.Println(db.G)
	case DataEnergy:
		db.E = append(db.E, v)
		fmt.Println(db.E)

	}

}
