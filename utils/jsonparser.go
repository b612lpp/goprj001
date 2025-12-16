package utils

import (
	"encoding/json"
	"net/http"

	"github.com/b612lpp/goprj001/metainf"
)

// Всё просто. Парсим боди на json в готовую структуру переданную через дженерик. надо бы типизировать дженерик воизбежании
func ParseUserData[T any](r *http.Request, v *T) error {
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {

		return metainf.ErrWrongData
	}
	return nil

}
