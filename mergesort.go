package main

//归并排序O(nlog2n)
func merge(nums []int, mid int) {
	//拷贝右半边
	right := make([]int, len(nums) - mid)
	copy(right, nums[mid:])
	//归并
	i, k := len(nums) - 1, len(right) - 1
	for j := mid - 1; j > -1 && k > -1; i-- {
		if nums[j] > right[k] {
			nums[i] = nums[j]
			j--
		} else {
			nums[i] = right[k]
			k--
		}
	}

	//将right归并
	for k > -1 {
		nums[i] = right[k]
		i, k = i - 1, k - 1
	}
	//left已在原数组中，不需要归并
}

func mergeSort(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	}
	if length == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return
	}
	mid := length/2
	mergeSort(nums[:mid])
	mergeSort(nums[mid:])
	merge(nums, mid)
}


func mergeSortCurCore(nums []int, ch chan int) {
	length := len(nums)
	if length < 3 {
		if length == 2 && nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		for _, c := range nums {
			ch <- c
		}
		close(ch)
		return
	}
	mid := length/2
	leftCh := make(chan int, mid - 0)
	rightCh := make(chan int, len(nums) - mid)
	go mergeSortCurCore(nums[:mid], leftCh)
	go mergeSortCurCore(nums[mid:], rightCh)
	//for i := range leftCh {
	//	ch <- i
	//}
	//for i := range rightCh {
	//	ch <- i
	//}


	i, okI := <- leftCh
	j, okJ := <- rightCh
	for okI && okJ {
		if i > j {
			ch <- j
			j, okJ = <- rightCh
		} else {
			ch <- i
			i, okI = <- leftCh
		}
	}
	for okI {
		ch <- i
		i, okI = <- leftCh
	}
	for okJ {
		ch <- j
		j, okJ = <- rightCh
	}
	close(ch)

}


func mergeSortCur(nums []int) {
	cpy := make([]int, len(nums))
	copy(cpy, nums)
	ch := make(chan int, len(nums))
	go mergeSortCurCore(nums, ch)
	i := 0
	for num := range ch {
		nums[i] = num
		i++
	}
}