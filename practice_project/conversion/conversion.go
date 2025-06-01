package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloats(strings []string) ([]float64, error) {
	var floats []float64 = make([]float64, len(strings))
	for strIndex, str := range strings {
		value, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("failed to string to float")
		}
		floats[strIndex] = value
	}
	return floats, nil
}
