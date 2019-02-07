package hamming

import "errors"

// Distance is a function that returns the Hamming distance between two DNA strand sequences
func Distance(a, b string) (int, error) {
	count := 0

	// if a and b are not equal length, return error
	if len(a) != len(b) {
		return 0, errors.New("a and b must be of equal length")
	}

	// calculate Hamming distance
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	// return result
	return count, nil
}
