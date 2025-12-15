package metainf

type DataGas struct {
	Value int `json:"value"`
}

type DataEnergy struct {
	Day   int `json:"day"`
	Night int `json:"night"`
	Summ  int `json:"summ"`
}
