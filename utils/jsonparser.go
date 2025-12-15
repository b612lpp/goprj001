package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

/*func ParseUsrData_(w http.ResponseWriter, r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("некорректные данные")
	}
	return nil
}*/

func ParseUserData[T any](w http.ResponseWriter, r *http.Request, v T) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return errors.New("некорректные данные")
	}
	return nil

}
