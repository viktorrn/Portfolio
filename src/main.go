package main

import (
	"fmt"
	"log"
	"net/http"
	"portfolio/handler"
)

func main() {
	http.HandleFunc("/", handler.HandleRouter(handler.IndexHandler))
	http.HandleFunc("/contact", handler.HandleRouter(handler.ContactHandler))
	http.Handle("POST /$", handler.HandleRouter(handler.PostHandler))
	http.Handle("/static/", http.StripPrefix("/static/", handler.HandleRouter(handler.StaticHandler)))

	fmt.Println("Server started at http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
