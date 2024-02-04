package main

import (
	"net/http"

	"github.com/danecwalker/cheez/pkg/cheez"
)

func main() {
	server := http.NewServeMux()

	server.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./examples/static"))))
	server.HandleFunc("/", cheez.NewCheez(App))

	http.ListenAndServe(":3000", server)
}
