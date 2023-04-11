package utils

import (
	"fmt"
	"github.com/NicholasLiem/IF2211_TugasKecil_3_RoutePlanning/models"
	"os"
	"strconv"
	"strings"
)

func Hello() {
	fmt.Println("Hello")
}
func parseNode(line string) (string, float64, float64, error) {
	words := strings.Fields(line)
	if len(words) != 3 {
		return "", 0, 0, fmt.Errorf("invalid number of fields (%d)", len(words))
	}

	name := strings.TrimSpace(words[0])

	latitude, err := strconv.ParseFloat(strings.TrimSpace(words[1]), 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("cannot parse latitude (%w)", err)
	}

	longitude, err := strconv.ParseFloat(strings.TrimSpace(words[2]), 64)
	if err != nil {
		return "", 0, 0, fmt.Errorf("cannot parse longitude (%w)", err)
	}

	return name, latitude, longitude, nil
}

func parseRow(line string, columns int) ([]float64, error) {
	words := strings.Fields(line)
	if len(words) != columns {
		return nil, fmt.Errorf("invalid number of columns (%d), expected %d", len(words), columns)
	}

	row := make([]float64, columns)
	for i, word := range words {
		val, err := strconv.ParseFloat(word, 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse column %d (%w)", i+1, err)
		}
		row[i] = val
	}

	return row, nil
}

func AdjacencyMatrixFromFile(filepath string) (*models.AdjacencyMatrix, error) {
	buf, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s (%w)", filepath, err)
	}
	lines := strings.Split(string(buf), "\n")
	if len(lines) < 3 {
		return nil, fmt.Errorf("invalid file format: expected at least 3 lines, got %d", len(lines))
	}

	count, err := strconv.ParseInt(strings.TrimSpace(lines[0]), 10, 32)
	if err != nil {
		return nil, fmt.Errorf("cannot parse node count (%w) at %s:1", err, filepath)
	}

	columnLabels := make([]string, count)
	latitudes := make([]float64, count)
	longitudes := make([]float64, count)

	for i := 0; i < int(count); i++ {
		name, lat, lon, err := parseNode(lines[i+1])
		if err != nil {
			return nil, fmt.Errorf("%w at %s:%d", err, filepath, i+2)
		}

		columnLabels[i] = name
		latitudes[i] = lat
		longitudes[i] = lon
	}

	matrix := make([][]float64, count)

	for i := 0; i < int(count); i++ {
		row, err := parseRow(lines[i+int(count)+1], int(count))
		if err != nil {
			return nil, fmt.Errorf("%w at %s:%d", err, filepath, i+int(count)+2)
		}

		matrix[i] = make([]float64, count)
		for j, val := range row {
			matrix[i][j] = float64(val)
		}
	}

	res := models.NewAdjacencyMatrix(int(count), columnLabels, latitudes, longitudes)
	res.Matrix = matrix

	return &res, nil
}
