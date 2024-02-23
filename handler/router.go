package handler

import (
	"net/http"
)

func HandleRouter(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Mange allowed URL:s

		m := true
		if m == false {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}
