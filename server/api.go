package server

import (
	"fmt"
	"net/http"

	"github.com/b612lpp/goprj001/server/data"
)

func (p *mySrv) home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello home")
}

func (p *mySrv) sendgas(w http.ResponseWriter, r *http.Request) {

	qq := data.DataGas{}
	qq.FillData(r)
	p.DB.AddInfo(qq)

}

func (p *mySrv) sendenergy(w http.ResponseWriter, r *http.Request) {
	qq := data.DataEnergy{}
	qq.FillData(r)
	p.DB.AddInfo(qq)
}
