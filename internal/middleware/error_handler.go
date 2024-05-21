package middleware

import (
	"crud-app/internal/errors"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				var appErr *errors.AppError
				switch e := err.(type) {
				case *errors.AppError:
					appErr = e
				case error:
					appErr = errors.InternalServerError(e)
				default:
					appErr = errors.InternalServerError(fmt.Errorf("%v", e))
				}

				log.Printf("HTTP %d - %s", appErr.Code, appErr.Message)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(appErr.Code)
				json.NewEncoder(w).Encode(appErr)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
