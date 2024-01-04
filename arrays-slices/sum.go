package arraysslices

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}

	return sum
}

// [[1 2] [0 9]] => slice of slices of integers: [][]int

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
