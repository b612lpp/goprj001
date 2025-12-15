package application

import (
	"errors"

	"github.com/b612lpp/goprj001/metainf"
)

func NewEnergyDataCase(dw EnergyDataWriter) *EnergyDataCase {
	return &EnergyDataCase{Edw: dw}
}

func (edc *EnergyDataCase) EnergyDataProcessor(ed metainf.DataEnergy) error {
	if ed.Day < 0 || ed.Night < 0 || ed.Day+ed.Night != ed.Summ {
		return errors.New("некорректные данные")
	}

	if err := edc.Edw.AddEnergy(ed); err != nil {
		return err
	}
	return nil
}
