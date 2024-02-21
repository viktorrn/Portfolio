package handler

import (
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/$")

func HandleRouter(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//validate the path
		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}
