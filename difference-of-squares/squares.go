package diffsquares

// SquareOfSum does stuff
func SquareOfSum(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i
	}

	return result * result
}

// SumOfSquares does stuff
func SumOfSquares(n int) int {
	result := 0
	for i := 1; i <= n; i++ {
		result += i * i
	}
	return result
}

// Difference also does stuff
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
