package sort


func QuickSort(nums []int) {

	//1.0
	//left, right := 0, len(nums) - 1
	//var target int
	//for left <= right {
	//	for left < right && nums[left] < nums[0] {
	//		left++
	//	}
	//	if left == right {
	//		if nums[right] < nums[0] {
	//			nums[0], nums[left] = nums[left], nums[0]
	//			target = left
	//		} else {
	//			nums[0], nums[left-1] = nums[left-1], nums[0]
	//			target = left - 1
	//		}
	//		break
	//	}
	//	for left < right && nums[right] >= nums[0] {
	//		right--
	//	}
	//	if left == right {
	//		nums[0], nums[left-1] = nums[left-1], nums[0]
	//		target = left - 1
	//		break
	//	}
	//	if left + 1 == right {
	//		nums[0], nums[left], nums[right] = nums[right], nums[0], nums[left]
	//		target = left
	//		break
	//	}
	//	nums[left], nums[right] = nums[right], nums[left]
	//	left, right = left + 1, right -1
	//}




	//2.0
	//targetNum := nums[left]
	//for left < right{
	//	for nums[right] >= targetNum && right > left {
	//		right--
	//	}
	//	nums[left] = nums[right]
	//	for nums[left] <= targetNum && right > left {
	//		left++
	//	}
	//	nums[right] = nums[left]
	//}
	//nums[right] = targetNum
	//if right - 1 > - 1 {
	//	QuickSort(nums[:right])
	//}
	//if right + 1 < len(nums) {
	//	QuickSort(nums[right+1:])
	//}

	//3.0
	if len(nums) < 2 {
		return
	}
	left, right := 0, len(nums) - 1
	value := nums[left] // 基准值
	for left < right {
		for nums[right] >= value && left < right { // 依次查找大于等于基准值的位置
			right--
		}
		nums[left] = nums[right]
		for nums[left] < value && left < right { // 依次查找小于基准值的位置
			left++
		}
		nums[right] = nums[left]
	}
	nums[left] = value
	if left > 1 {
		QuickSort(nums[:left])
	}
	if left < len(nums) - 2 {
		QuickSort(nums[left+1:])
	}

}