package sort

//worst:O(n2) avg:O(n2) best:O(n)
func InsertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		j := i - 1
		for j > -1 && tmp < nums[j] {
			nums[j+1] = nums[j]
			j--
		}
		nums[j] = tmp
	}
}
