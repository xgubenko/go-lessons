package arraysandslices

func Sum(numbers []int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func SumAll(arrays ...[]int) []int {
	var sums []int

	for _, arr := range arrays {
		sums = append(sums, Sum(arr))
	}

	return sums
}

func SumAllTails(arrays ...[]int) []int {
	var sums []int

	for _, arr := range arrays {

		if len(arr) == 0 {
			sums = append(sums, 0)
		} else {
			tail := arr[1:]
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
