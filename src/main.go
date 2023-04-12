package main

import (
	"github.com/NicholasLiem/Tucil3_13521083_13521135/handlers"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/parse", handlers.ParseHandler)
	http.HandleFunc("/search", handlers.SearchHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
