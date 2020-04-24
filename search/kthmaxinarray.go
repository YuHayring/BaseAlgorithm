package search

// 1 <= k <= len(nums)
func KthMaxInArray(nums []int, k int) int {

	left, right := 0, len(nums) - 1
	targetNum := nums[left]
	for left < right{
		for nums[right] <= targetNum && right > left {
			right--
		}
		nums[left] = nums[right]
		for nums[left] >= targetNum && right > left {
			left++
		}
		nums[right] = nums[left]
	}
	nums[right] = targetNum
	if right + 1 == k {
		return targetNum
	} else {
		if right + 1 > k {
			return KthMaxInArray(nums[:right], k)
		}
		return KthMaxInArray(nums[right+1:], k - 5) + 5
	}


}