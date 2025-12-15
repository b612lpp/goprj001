package utils

import (
	"encoding/json"
	"net/http"
)

func ParseUsrData(r *http.Request) err {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}
