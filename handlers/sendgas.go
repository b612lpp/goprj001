package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/b612lpp/goprj001/application"
	"github.com/b612lpp/goprj001/metainf"
)

type GasHandler struct {
	Gdc *application.GasDataCase
}

func NewGasHandler(gdc *application.GasDataCase) *GasHandler {
	return &GasHandler{Gdc: gdc}
}

func (gh *GasHandler) SendGas(w http.ResponseWriter, r *http.Request) {
	var v metainf.DataGas
	err := json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	gh.Gdc.AddGasRow(v)
	fmt.Println(v)

}
