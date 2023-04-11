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

func ParseNode(line string) (string, float64, float64, error) {
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

func ParseRow(line string, columns int) ([]int64, error) {
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

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	p := 0.017453292519943295 // Math.PI / 180
	c := math.Cos
	a := 0.5 - c((lat2-lat1)*p)/2 +
		c(lat1*p)*c(lat2*p)*
			(1-c((lon2-lon1)*p))/2

	return 12742 * math.Asin(math.Sqrt(a)) // 2 * R; R = 6371 km
}
