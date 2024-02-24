package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func StaticHandler(w http.ResponseWriter, r *http.Request) {

	if strings.HasSuffix(r.URL.Path, ".js") {
		fmt.Println("JS file")
		w.Header().Set("Content-Type", "application/javascript")
		w.Header().Set("Cache-Control", "max-age=2592000")

	}
	if strings.HasSuffix(r.URL.Path, ".css") {
		w.Header().Set("Content-Type", "text/css")

	}

	if strings.HasPrefix(r.URL.Path, "svg/") {
		w.Header().Set("Content-Type", "image/svg+xml")

	}

	http.ServeFile(w, r, "static/"+r.URL.Path)
}
