package utils

import (
	"errors"
	"fmt"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"math"
	"os"
	"strconv"
	"strings"
)

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

func ParseToAdjacencyMatrix(buf string) (models.AdjacencyMatrix, error) {
	var res models.AdjacencyMatrix
	lines := strings.Split(buf, "\n")
	count, err := strconv.ParseInt(strings.TrimSpace(lines[0]), 10, 32)
	if err != nil {
		return res, err
	}
	columnLabels := make([]string, count)
	latitudes := make([]float64, count)
	longitudes := make([]float64, count)
	for i := 0; i < len(columnLabels); i++ {
		label, latitude, longitude, err := parseNode(lines[i+1])
		if err != nil {
			return res, err
		}
		columnLabels[i] = label
		latitudes[i] = latitude
		longitudes[i] = longitude
	}
	matrix := make([][]int64, count)
	for i := 0; i < len(matrix); i++ {
		matrix[i], err = parseRow(lines[i+int(count)+1], int(count))
		if err != nil {
			return res, err
		}
	}
	res = models.NewAdjacencyMatrix(int(count), columnLabels, latitudes, longitudes)
	res.Matrix = matrix
	return res, nil
}

func AdjacencyMatrixFromFile(filepath string) (models.AdjacencyMatrix, error) {
	var res models.AdjacencyMatrix
	buf, err := os.ReadFile(filepath)
	if err != nil {
		msg := fmt.Sprintf("[ERROR] cannot read file %s (%s)", filepath, err.Error())
		return res, errors.New(msg)
	}
	lines := strings.Split(string(buf), "\n")
	count, err := strconv.ParseInt(strings.TrimSpace(lines[0]), 10, 32)
	if err != nil {
		msg := fmt.Sprintf("[ERROR] cannot parse node count (%s) at %s:1", err.Error(), filepath)
		return res, errors.New(msg)
	}
	columnLabels := make([]string, count)
	latitudes := make([]float64, count)
	longitudes := make([]float64, count)
	for i := 0; i < len(columnLabels); i++ {
		label, latitude, longitude, ok := parseNode(lines[i+1])
		if ok != nil {
			msg := fmt.Sprintf("[ERROR] %s at %s:%d", ok.Error(), filepath, i+1)
			return res, errors.New(msg)
		}
		columnLabels[i] = label
		latitudes[i] = latitude
		longitudes[i] = longitude
	}
	matrix := make([][]int64, count)
	for i := 0; i < len(matrix); i++ {
		matrix[i], err = parseRow(lines[i+int(count)+1], int(count))
		if err != nil {
			msg := fmt.Sprintf("[ERROR] %s at %s:%d", err.Error(), filepath, i+int(count)+1)
			return res, errors.New(msg)
		}
	}
	res = models.NewAdjacencyMatrix(int(count), columnLabels, latitudes, longitudes)
	res.Matrix = matrix
	return res, nil
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	p := 0.017453292519943295 // Math.PI / 180
	c := math.Cos
	a := 0.5 - c((lat2-lat1)*p)/2 +
		c(lat1*p)*c(lat2*p)*
			(1-c((lon2-lon1)*p))/2

	return 12742 * math.Asin(math.Sqrt(a)) // 2 * R; R = 6371 km
}
