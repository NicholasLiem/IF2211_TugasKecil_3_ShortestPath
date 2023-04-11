package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/utils"
	"io"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/parse", parse)
	http.HandleFunc("/upload", handlers.UploadHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}

func parse(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	start := bytes.IndexByte(data, byte(','))
	decodedData := make([]byte, base64.StdEncoding.DecodedLen(len(data[start+1:])))
	_, err = base64.StdEncoding.Decode(decodedData, data[start+1:])
	if err != nil {
		_ = fmt.Errorf("error: %s", err.Error())
		return
	}
	adjMat, err := utils.ParseToAdjacencyMatrix(string(decodedData))
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	graph := models.NewGraphFromAdjacencyMatrix(adjMat)
	err = json.NewEncoder(w).Encode(graph)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
