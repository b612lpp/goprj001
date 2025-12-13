package server

import "net/http"

type mySrv struct {
	Port string
	Mux  http.ServeMux
}

func NewMySrv(p string) *mySrv {
	h := mySrv{p, *http.NewServeMux()}
	h.Mux.HandleFunc("/", setCorsMW(h.home))
	h.Mux.HandleFunc("/sendgas", setCorsMW(h.sendgas))
	h.Mux.HandleFunc("/sendenergy", setCorsMW(h.sendenergy))

	return &h
}

func setCorsMW(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next(w, r)
	}
}
