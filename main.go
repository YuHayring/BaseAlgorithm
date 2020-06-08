package main

import (
	"./basestruct"
	"./str"
	"fmt"
)
func main() {
	fmt.Println("start")

	fmt.Println(str.Kmp("aabaaabaaac","aabaaac"))

	//fmt.Println(next("aabaaac"))

	fmt.Println("end")
}


func RBTreeTest() {
	a := []basestruct.Integer{10, 40, 30, 60, 90, 70, 20, 50, 80}
	ilen := len(a)

	rootPointer := new(basestruct.RedBlackTree)
	fmt.Printf("== 原始数据: ")

	for i := 0; i < ilen; i++ {
		fmt.Printf("%d ", a[i])
	}

	fmt.Printf("\n")

	for i := 0; i < ilen; i++ {
		var num basestruct.Comparable = a[i]
		rootPointer.Insert(&num)
	}
	root,_ := rootPointer.GetRoot()
	basestruct.RBTreePrint(root,-1)
	fmt.Print("\n")
	basestruct.RBTreePrint(root,0)
	fmt.Print("\n")
	basestruct.RBTreePrint(root,1)
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("\n")
	for i := 0; i < ilen; i++ {
		var num basestruct.Comparable = a[i]
		rootPointer.DeleteVal(&num)
		basestruct.RBTreePrint(root,0)
		fmt.Print("\n")

	}

}

func next(findStr string) (next []int) {
	k := 0
	next = make([]int, len(findStr))
	next[0] = k
	for i := 1; i < len(findStr); i++ {
		for k > 0 && findStr[k] != findStr[i] {
			k = next[k-1]
		}
		if findStr[k] == findStr[i] {
			k++
		}
		next[i] = k
	}
	return next
}