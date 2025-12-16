package midlware

import (
	"fmt"
	"net/http"
)

func MWCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("первый хэндлер вызван")
		next(w, r)
	}

}
