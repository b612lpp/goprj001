package utils

import "encoding/json"

func JsonFormer[T any](dot T) ([]byte, error) {

	q, err := json.Marshal(dot)
	if err != nil {
		return nil, err
	}
	return q, nil
}
