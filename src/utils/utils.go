package utils

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ParseNode(line string) (string, float64, float64, error) {
	words := strings.Split(line, ",")
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

func ParseRow(line string, columns int) ([]float64, error) {
	words := strings.Fields(line)
	if len(words) != columns {
		return nil, fmt.Errorf("invalid number of columns (%d), expected %d", len(words), columns)
	}

	row := make([]float64, columns)
	for i, word := range words {
		val, err := strconv.ParseFloat(strings.Trim(strings.TrimSpace(word), "\x00"), 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse column %d (%w)", i+1, err)
		}
		row[i] = val
	}

	return row, nil
}

func Distance(lat1, lon1, lat2, lon2, weightRange float64) float64 {
	deltaX := lon2 - lon1
	deltaY := lat2 - lat1
	return math.Sqrt(deltaX*deltaX+deltaY*deltaY) * weightRange
}
