package utils

import (
	"encoding/json"
	"net/http"

	"github.com/b612lpp/goprj001/metainf"
)

func ParseUserData[T any](r *http.Request, v *T) error {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {

		return metainf.ErrWrongData
	}
	return nil

}
