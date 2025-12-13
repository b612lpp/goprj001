package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/b612lpp/goprj001/server/data"
)

func (p *mySrv) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello home")
}

func (p *mySrv) sendgas(w http.ResponseWriter, r *http.Request) {
	qq := data.DataGas{}
	json.NewDecoder(r.Body).Decode(&qq)
	fmt.Printf("hello gas %d", qq.Vlue)
}

func (p *mySrv) sendenergy(w http.ResponseWriter, r *http.Request) {
	qq := data.DataEnergy{}
	json.NewDecoder(r.Body).Decode(&qq)
	fmt.Printf("hello energy %d", qq.Summ)
}
