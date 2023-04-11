package main

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type PageData struct {
	FileName string
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/upload", uploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	const MaxMultipartMemory = 100 << 20 // 100MB
	r.ParseMultipartForm(MaxMultipartMemory)

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Redirect(w, r, "/?uploaded=false&message=File upload failed", http.StatusSeeOther)
		return
	}
	defer file.Close()

	ext := filepath.Ext(header.Filename)

	switch ext {
	case ".txt":
	default:
		http.Redirect(w, r, "/?uploaded=false&message=File upload failed", http.StatusSeeOther)
		return
	}

	f, err := os.OpenFile("./uploads/"+header.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Redirect(w, r, "/?uploaded=false&message=File upload failed", http.StatusSeeOther)
		return
	}
	defer f.Close()

	io.Copy(f, file)
	http.Redirect(w, r, "/?uploaded=true&message=File uploaded successfully", http.StatusSeeOther)
}
