package data

import (
	"encoding/json"
	"net/http"
	"time"
)

type DataGas struct {
	Vlue int `json:"day"`
	Data time.Time
}

type DataEnergy struct {
	Day   int `json:"day"`
	Night int `json:"night"`
	Summ  int `json:"summ"`
	Data  time.Time
}

type DataFiller interface {
	FillData()
}

type TypeGetter interface {
	Type() string
}

func (s *DataGas) FillData(r *http.Request) {
	json.NewDecoder(r.Body).Decode(&s)
	s.Data = time.Now()

}

func (s *DataEnergy) FillData(r *http.Request) {
	json.NewDecoder(r.Body).Decode(&s)
	s.Data = time.Now()
}

func (s *DataEnergy) Type() string {
	return "energy"
}

func (s *DataGas) Type() string {
	return "gas"
}
