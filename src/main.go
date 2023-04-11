package main

import (
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/handlers"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/parse", handlers.ParseHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
