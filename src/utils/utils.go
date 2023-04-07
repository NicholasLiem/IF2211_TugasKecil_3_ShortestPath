package utils

import (
	"errors"
	"fmt"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"os"
	"strconv"
	"strings"
)

func Hello() {
	fmt.Println("Hello")
}

// TODO: proper parser

func parseNode(line string) (string, float64, float64, error) {
	words := strings.Split(line, " ")
	name := strings.TrimSpace(words[0])
	latitude, ok := strconv.ParseFloat(strings.TrimSpace(words[1]), 64)
	if ok != nil {
		return "", 0, 0, errors.New("cannot parse latitude (" + ok.Error() + ")")
	}
	longitude, ok := strconv.ParseFloat(strings.TrimSpace(words[2]), 64)
	if ok != nil {
		return "", 0, 0, errors.New("cannot parse longitude (" + ok.Error() + ")")
	}
	return name, latitude, longitude, nil
}

func parseRow(line string, columns int) ([]int64, error) {
	words := strings.Split(line, " ")
	row := make([]int64, columns)
	var ok error
	for i := range row {
		row[i], ok = strconv.ParseInt(strings.TrimSpace(words[i]), 10, 32)
		if ok != nil {
			return []int64{}, errors.New("cannot parse column (" + ok.Error() + ")")
		}
	}
	return row, nil
}

func AdjacencyMatrixFromFile(filepath string) (*models.AdjacencyMatrix, error) {
	buf, ok := os.ReadFile(filepath)
	if ok != nil {
		msg := fmt.Sprintf("[ERROR] cannot read file %s (%s)", filepath, ok.Error())
		return nil, errors.New(msg)
	}
	lines := strings.Split(string(buf), "\n")
	count, ok := strconv.ParseInt(strings.TrimSpace(lines[0]), 10, 32)
	if ok != nil {
		msg := fmt.Sprintf("[ERROR] cannot parse node count (%s) at %s:1", ok.Error(), filepath)
		return nil, errors.New(msg)
	}
	columnLabels := make([]string, count)
	latitudes := make([]float64, count)
	longitudes := make([]float64, count)
	for i := 0; i < len(columnLabels); i++ {
		label, latitude, longitude, ok := parseNode(lines[i+1])
		if ok != nil {
			msg := fmt.Sprintf("[ERROR] %s at %s:%d", ok.Error(), filepath, i+1)
			return nil, errors.New(msg)
		}
		columnLabels[i] = label
		latitudes[i] = latitude
		longitudes[i] = longitude
	}
	matrix := make([][]int64, count)
	for i := 0; i < len(matrix); i++ {
		matrix[i], ok = parseRow(lines[i+int(count)+1], int(count))
		if ok != nil {
			msg := fmt.Sprintf("[ERROR] %s at %s:%d", ok.Error(), filepath, i+int(count)+1)
			return nil, errors.New(msg)
		}
	}
	res := models.NewAdjacencyMatrix(int(count), columnLabels, latitudes, longitudes)
	res.Matrix = matrix
	return &res, nil
}
