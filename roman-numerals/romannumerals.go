package romannumerals

import "errors"

// ToRomanNumeral converts number to roman numeral string
func ToRomanNumeral(n int) (string, error) {
	if n > 3000 || n <= 0 {
		return "", errors.New("invalid number")
	}
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	result := ""

	for _, conversion := range conversions {
		for n >= conversion.value {
			result += conversion.digit
			n -= conversion.value
		}
	}

	return result, nil
}
