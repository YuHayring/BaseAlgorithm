package sort

//worst:O(n2) avg:O(n2) best:O(n)
func InsertSort(nums []int){
	for i, _ := range nums {
		for i > 0 && nums[i] < nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			i--
		}
	}
}
