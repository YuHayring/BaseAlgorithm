package main


func quickSort(nums []int) {
	if len(nums) < 3 {
		if len(nums) == 2 && nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return
	}

	left, right := 1, len(nums) - 1
	var target int
	for left <= right {
		for left < right && nums[left] < nums[0] {
			left++
		}
		if left == right {
			if nums[right] < nums[0] {
				nums[0], nums[left] = nums[left], nums[0]
				target = left
			} else {
				nums[0], nums[left-1] = nums[left-1], nums[0]
				target = left - 1
			}
			break
		}
		for left < right && nums[right] >= nums[0] {
			right--
		}
		if left == right {
			nums[0], nums[left-1] = nums[left-1], nums[0]
			target = left - 1
			break
		}
		if left + 1 == right {
			nums[0], nums[left], nums[right] = nums[right], nums[0], nums[left]
			target = left
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		left, right = left + 1, right -1
	}
	if target - 1 > - 1 {
		quickSort(nums[:target])
	}
	if target + 1 < len(nums) {
		quickSort(nums[target+1:])
	}


}