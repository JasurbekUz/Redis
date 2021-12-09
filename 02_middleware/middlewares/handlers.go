package middlewares

import (
	"net/http"
)

func middleware (fn func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {

	return nil
}