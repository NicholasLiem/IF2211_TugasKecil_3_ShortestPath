package handlers

import (
	"encoding/json"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"io"
	"log"
	"net/http"
	"strings"
)

type resultDTO struct {
	Route []int
	Cost  int64
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var objmap map[string]json.RawMessage
	err = json.Unmarshal(data, &objmap)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var dst int
	err = json.Unmarshal(objmap["dst"], &dst)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var src int
	err = json.Unmarshal(objmap["src"], &src)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var method string
	err = json.Unmarshal(objmap["method"], &method)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var graph models.Graph
	err = json.Unmarshal(objmap["graph"], &graph)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println(src)
	log.Println(dst)
	log.Println(method)
	log.Println(graph)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := resultDTO{}
	if strings.ToLower(method) == "a*" {
		result.Route, result.Cost = models.AStarSearch(graph, src, dst)
	} else if strings.ToLower(method) == "ucs" {
		result.Route, result.Cost = models.UniformCostSearch(graph, src, dst)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
