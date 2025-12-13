package server

import (
	"net/http"

	"github.com/b612lpp/goprj001/server/data"
)

type mySrv struct {
	Port string
	Mux  http.ServeMux
	DB   *data.DataBase
}

func NewMySrv(p string) *mySrv {
	h := mySrv{p, *http.NewServeMux(), data.NewDb()}
	h.Mux.HandleFunc("/", setCorsMW(h.home))
	h.Mux.HandleFunc("/sendgas", setCorsMW(h.sendgas))
	h.Mux.HandleFunc("/sendenergy", setCorsMW(h.sendenergy))
	h.Mux.HandleFunc("/gethistory", setCorsMW(h.gethistory))

	return &h
}

func setCorsMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next(w, r)
	}
}
