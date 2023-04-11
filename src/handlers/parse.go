package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"io"
	"log"
	"net/http"
)

func ParseHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	start := bytes.IndexByte(data, byte(','))
	decodedData := make([]byte, base64.StdEncoding.DecodedLen(len(data[start+1:])))
	_, err = base64.StdEncoding.Decode(decodedData, data[start+1:])
	if err != nil {
		_ = fmt.Errorf("error: %s", err.Error())
		return
	}
	mat, err := models.ParseToAdjacencyMatrix(string(decodedData))
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	graph := models.NewGraphFromAdjacencyMatrix(mat)
	err = json.NewEncoder(w).Encode(graph)
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err.Error())
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
