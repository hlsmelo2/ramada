package middlewares

import (
	"net/http"
	"ramada/api/src/auth"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if _error := auth.ValidateToken(request); _error != nil {
			http.Error(writer, "Not authorized", http.StatusOK)

			return
		}

		next(writer, request)
	}
}
