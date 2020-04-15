package main

//冒泡排序：O(n^2)
func bubbleSort(nums []int) {
	for i := len(nums); i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums[j - 1] > nums[j] {
				nums[j - 1], nums[j] = nums[j], nums[j-1]
			}
		}
	}
}