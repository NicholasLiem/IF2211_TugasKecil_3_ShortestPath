package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const MaxMultipartMemory = 100 << 20 // 100MB

func UploadHandler(w http.ResponseWriter, r *http.Request) {
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
