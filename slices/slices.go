package slices

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	var sums []int
	for _, num := range numbers {
		sums = append(sums, Sum(num))
	}
	return sums
}
