package search

//O(log2n)
func BinarySearch(sortedNums []int, target int) int {
	length := len(sortedNums)
	if length == 0 {
		return  -1
	}
	if length == 1 {
		if sortedNums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	mid := length/2
	if sortedNums[mid] == target {
		return mid
	}
	if target < sortedNums[mid] {
		return BinarySearch(sortedNums[:mid], target)
	} else {
		return BinarySearch(sortedNums[mid + 1:], target) + mid + 1
	}

}
