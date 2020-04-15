package main

import (
	"fmt"
	"math/rand"
)

func main() {

	nums := make([]int, 10)
	correct := make([]int, 10)
	error := false


	for times := 0; times < 100; times++ {
		for i,_ := range nums {
			nums[i] = rand.Intn(10)
			correct[i] = nums[i]
		}
		//fmt.Print("src     :")
		//fmt.Println(nums)
		quickSort(nums)
		mergeSort(correct)
		error = false
		for i,v := range nums {
			if v != correct[i] {
				error = true
				break
			}
		}
		if error {
			fmt.Print("error   :")
			fmt.Println(nums)
			fmt.Print("correct :")
			fmt.Println(correct)

			fmt.Println("ERROR")
		}
		//else {
		//	fmt.Println("success")
		//}
	}


	//nums := []int{3,6,2,8,7,4,7,3,6,0}
	//quickSort(nums)
	//fmt.Println(nums)

}
