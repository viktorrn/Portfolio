package main

import (
	"log"
	"net/http"
	"portfolio/handler"
)

func main() {
	http.HandleFunc("/", handler.HandleRouter(handler.UserHandler))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
