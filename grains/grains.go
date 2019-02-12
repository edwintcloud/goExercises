package grains

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Square determines how many grains are on each square
func Square(n int) (uint64, error) {

	// edge case: if n <= 0 or > 64, return an error
	if n <= 0 || n > 64 {
		return 0, errors.New("invalid")
	}

	// edge case: if n is 1 return 1
	if n == 1 {
		return 1, nil
	}

	// return result
	return conv(math.Pow(2, float64(n-1))), nil
}

// Total determines the total amount of grains
func Total() uint64 {
	return conv(math.Pow(2, 64))
}

// conv is a helper function that converts a float64 to uint64
func conv(n float64) uint64 {
	s := fmt.Sprintf("%.0f", n)
	result, _ := strconv.ParseUint(s, 10, 64)
	return result
}
