package main

import (
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/handlers"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", handlers.IndexHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
