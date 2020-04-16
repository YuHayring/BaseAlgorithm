package main

import (
	pg "./permutationgenerate"
	"fmt"
)
func main() {

	//nums := make([]int, 10)
	//correct := make([]int, 10)
	//error := false
	//
	//
	//for times := 0; times < 1000 ; times++ {
	//	for i,_ := range nums {
	//		nums[i] = rand.Intn(10)
	//		correct[i] = nums[i]
	//	}
	//	//fmt.Print("src     :")
	//	//fmt.Println(nums)
	//	sort.InsertSort(nums)
	//	sort.MergeSort(correct)
	//	error = false
	//	for i,v := range nums {
	//		if v != correct[i] {
	//			error = true
	//			break
	//		}
	//	}
	//	if error {
	//		fmt.Print("error   :")
	//		fmt.Println(nums)
	//		fmt.Print("correct :")
	//		fmt.Println(correct)
	//
	//		fmt.Println("ERROR")
	//	}
	//	//else {
	//	//	fmt.Println("success")
	//	//}
	//}


	//nums := []int{5,4,3,2,1}
	//sort.MergeSortCur(nums)
	//fmt.Println(nums)

	//fmt.Println(BinarySearch([]int{0,1,2,3,4,5,6,7,8,9},6))



	src := []int{1,2,3,4,5}
	result := pg.BackTracking(src)
	for _,arr := range result {
		fmt.Println(arr)
	}

}
