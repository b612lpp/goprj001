package metainf

import "time"

type DataGas struct {
	Value int       `json:"value"`
	Time  time.Time `json:"data"`
}

type DataEnergy struct {
	Day   int       `json:"day"`
	Night int       `json:"night"`
	Summ  int       `json:"summ"`
	Time  time.Time `json:"data"`
}
