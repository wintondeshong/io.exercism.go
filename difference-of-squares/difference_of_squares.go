package diffsquares

const testVersion = 1

func Difference(i int) int {
	return SquareOfSums(i) - SumOfSquares(i)
}

func SquareOfSums(i int) int {
	sum := 0

	for x := 1; x <= i; x++ {
		sum += x
	}

	return sum * sum
}

func SumOfSquares(i int) int {
	sum := 0

	for x := 1; x <= i; x++ {
		sum += x * x
	}

	return sum
}
