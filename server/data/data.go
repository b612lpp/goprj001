package data

type DataGas struct {
	Vlue int `json:"day"`
}

type DataEnergy struct {
	Day   int `json:"day"`
	Night int `json:"night"`
	Summ  int `json:"summ"`
}
