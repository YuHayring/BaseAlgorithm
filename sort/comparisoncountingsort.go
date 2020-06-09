package sort

func ComparisonCountingSort(num []int) []int {
	index := make([]int, len(num))
	for i, val := range num {
		for j := i + 1; j < len(num); j++ {
			if val > num[j] {
				index[i]++
			} else {
				index[j]++
			}
		}
	}
	result := make([]int, len(num))
	for i, val := range index {
		result[val] = num[i]
	}
	return result
}