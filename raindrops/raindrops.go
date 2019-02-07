package raindrops

import (
	"strconv"
)

// Convert is a function that converts a number into a text string based on factors
func Convert(n int) string {
	var result string

	// get number factors
	if n%3 == 0 {
		result += "Pling"
	}
	if n%5 == 0 {
		result += "Plang"
	}
	if n%7 == 0 {
		result += "Plong"
	}

	// if result doesn't have any text, set result to string version of number
	if len(result) == 0 {
		result += strconv.Itoa(n)
	}

	// return result
	return result
}
