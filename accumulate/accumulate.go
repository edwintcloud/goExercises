package accumulate

// Accumulate takes an input slice of strings and an input function to apply to slice
func Accumulate(s []string, fn func(string) string) []string {
	result := []string{}
	for _, v := range s {
		result = append(result, fn(v))
	}
	return result
}
