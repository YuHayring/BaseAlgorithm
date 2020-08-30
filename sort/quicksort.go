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
	//
	left := HoarePartition(nums)
	if left > 1 {
		QuickSort(nums[:left])
	}
	if left < len(nums) - 2 {
		QuickSort(nums[left+1:])
	}

}


func LomutoPartition(src []int) int {
	p := src[0]
	s := 0
	for i := 1; i < len(src); i++ {
		if src[i] < p {
			s++
			src[s], src[i] = src[i], src[s]
		}
	}
	src[0], src[s] = src[s], src[0]
	return s
}

func HoarePartition(src []int) int {
	p := src[0]
	left := 1
	right := len(src) - 1
	for left < right {
		for left < right && src[left] < p {
			left++
		}
		for left < right && src[right] >=p {
			right--
		}
		src[left], src[right] = src[right], src[left]
	}
	src[left], src[right] = src[right], src[left]
	var center int
	if src[right] < p {
		center = right
	} else {
		center = right-1
	}
	src[0], src[center] = src[center], src[0]

	return center

}