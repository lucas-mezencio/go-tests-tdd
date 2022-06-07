package slices

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	length := len(numbers)
	sums := make([]int, length)
	for i, num := range numbers {
		sums[i] = Sum(num)
	}
	return sums
}
