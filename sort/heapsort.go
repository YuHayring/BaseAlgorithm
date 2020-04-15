package sort

//堆排序
func HeapSort(nums []int, comparator func(a int, b int) bool ) {
	end := len(nums)
	for index := (end/2) - 1; index > -1; index-- {
		adjust(nums, index, comparator, end)
	}

	end--
	for end > -1 {
		nums[0], nums[end] = nums[end], nums[0]
		adjust(nums, 0, comparator, end)
		end--
	}

}

func adjust(nums []int, index int, comparator func(a int, b int) bool, end int) {
	var left, right, target int
	left = 2 * index + 1
	for left < end {
		right = left + 1
		target = left
		if right < end && comparator(nums[right], nums[left]) {
			target = right
		}
		if comparator(nums[target], nums[index]) {
			nums[index], nums[target] = nums[target], nums[index]
		} else {
			break
		}
		index = target
		left = 2 * index + 1
	}
}

func Bigger(a int, b int) bool {
	return a > b
}

func Smaller(a int, b int) bool {
	return a < b
}
