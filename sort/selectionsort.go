package sort

//选择排序：O(n^2)
func selectionSort(nums []int) {
	for i := len(nums) - 1; i > 0; i-- {
		max := nums[i]
		index := i
		for j := 0; j < i; j++ {
			if nums[j] > max {
				index, max = j, nums[j]
			}
		}
		nums[i], nums[index] = nums[index], nums[i]
	}
}
