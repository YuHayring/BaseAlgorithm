package sort
func InsertSort(nums []int){
	for i, _ := range nums {
		for i > 0 && nums[i] < nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			i--
		}
	}
}
