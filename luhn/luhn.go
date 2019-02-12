package luhn

import (
	"strconv"
	"strings"
)

// Valid checks the string representation of a number to determine if
// it is a valid luhn number or not.
func Valid(input string) bool {

	// strip spaces before checking
	input = strings.Replace(input, " ", "", -1)

	// strings of length 1 or less are not valid
	if len(input) <= 1 {
		return false
	}

	// convert to integer slice, if there are non-digit characters err will
	// be set and function will return false
	numbers := []int{}
	for _, c := range input {
		parsed, err := strconv.ParseInt(string(c), 10, 64)
		if err != nil {
			return false
		}
		numbers = append(numbers, int(parsed))
	}

	// sum is the sum of all of our numbers, this will be used below
	sum := 0

	// double every second digit, if number is greater than 9, subtract 9
	// from the product; sum all the digits also.
	// note that we must have a i (reverse count) to keep track of the index
	// and a j (counting from 1 foward) to keep track of when we are on a second digit
	for i, j := len(numbers)-1, 1; i >= 0; i, j = i-1, j+1 {

		if j%2 == 0 {
			numbers[i] *= 2
			if numbers[i] > 9 {
				numbers[i] -= 9
			}
		}
		sum += numbers[i]

	}

	// if the sum is evenly divisible by 10, then the input is valid - return true
	if sum%10 == 0 {
		return true
	}

	// otherwise return false
	return false
}
