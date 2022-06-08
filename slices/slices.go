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

func SumAllTails(numbers ...[]int) []int {
	var sums []int

	for _, nums := range numbers {
		tail := nums[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
