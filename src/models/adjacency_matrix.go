package models

import (
	"errors"
	"fmt"
	"github.com/NicholasLiem/Tucil3_13521083_13521135/utils"
	"os"
	"strconv"
	"strings"
)

type AdjacencyMatrix struct {
	Matrix       [][]float64
	ColumnLabels []string
	Latitudes    []float64
	Longitudes   []float64
	NodesCount   int
}

func (am AdjacencyMatrix) GetNodesCount() int {
	return am.NodesCount
}

func NewAdjacencyMatrix(nodeCount int, columnLabels []string, latitudes, longitudes []float64) AdjacencyMatrix {
	matrix := make([][]float64, nodeCount)
	for i := range matrix {
		matrix[i] = make([]float64, nodeCount)
	}

	return AdjacencyMatrix{
		Matrix:       matrix,
		ColumnLabels: columnLabels,
		NodesCount:   nodeCount,
		Latitudes:    latitudes,
		Longitudes:   longitudes,
	}
}

func ParseToAdjacencyMatrix(buf string) (AdjacencyMatrix, error) {
	var res AdjacencyMatrix
	lines := strings.Split(buf, "\n")
	count, err := strconv.ParseInt(strings.TrimSpace(lines[0]), 10, 32)
	if err != nil {
		return res, err
	}
	columnLabels := make([]string, count)
	latitudes := make([]float64, count)
	longitudes := make([]float64, count)
	for i := 0; i < len(columnLabels); i++ {
		label, latitude, longitude, err := utils.ParseNode(lines[i+1])
		if err != nil {
			return res, err
		}
		columnLabels[i] = label
		latitudes[i] = latitude
		longitudes[i] = longitude
	}
	matrix := make([][]float64, count)
	for i := 0; i < len(matrix); i++ {
		matrix[i], err = utils.ParseRow(lines[i+int(count)+1], int(count))
		if err != nil {
			return res, err
		}
	}
	res = NewAdjacencyMatrix(int(count), columnLabels, latitudes, longitudes)
	res.Matrix = matrix
	return res, nil
}

func AdjacencyMatrixFromFile(filepath string) (AdjacencyMatrix, error) {
	var res AdjacencyMatrix
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
		label, latitude, longitude, ok := utils.ParseNode(lines[i+1])
		if ok != nil {
			msg := fmt.Sprintf("[ERROR] %s at %s:%d", ok.Error(), filepath, i+1)
			return res, errors.New(msg)
		}
		columnLabels[i] = label
		latitudes[i] = latitude
		longitudes[i] = longitude
	}
	matrix := make([][]float64, count)
	for i := 0; i < len(matrix); i++ {
		matrix[i], err = utils.ParseRow(lines[i+int(count)+1], int(count))
		if err != nil {
			msg := fmt.Sprintf("[ERROR] %s at %s:%d", err.Error(), filepath, i+int(count)+1)
			return res, errors.New(msg)
		}
	}
	res = NewAdjacencyMatrix(int(count), columnLabels, latitudes, longitudes)
	res.Matrix = matrix
	return res, nil
}
