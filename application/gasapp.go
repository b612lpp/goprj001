package application

import "github.com/b612lpp/goprj001/metainf"

type DataWriter interface {
	AddGas(metainf.DataGas) error
}
type GasDataCase struct {
	Dw DataWriter
}

func NewGasDataCase(dw DataWriter) *GasDataCase {

	q := GasDataCase{Dw: dw}
	return &q
}

func (gdc *GasDataCase) AddGasRow(dg metainf.DataGas) {
	gdc.Dw.AddGas(dg)
}
