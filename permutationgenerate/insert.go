package permutationgenerate

func InsertGenerate(src []int) [][]int {
	length := len(src)
	if length == 0 {
		return [][]int{}
	}
	if length == 1 {
		return [][]int{src}
	}
	for i := length - 1; i > 1; i-- {
		length *= i
	}
	result := make([][]int, length)
	for i, _ := range  result {
		result[i] = make([]int, 0, len(src))
	}
	insertGenratorCore(result, src, 0)
	return result
}

func insertGenratorCore(result [][]int, src []int, start int) {
	if start == len(src) - 1 {
		result[0] = append(result[0], src[start])
		return
	}
	insertGenratorCore(result, src, start + 1)
	i := 0
	for  len(result[i]) != 0 {
		result[i] = append(result[i] , src[start])
		i++
	}
	index := i
	for j :=0; j < i; j++ {
		cur := make([]int, len(result[j]))
		copy(cur, result[j])
		for k := len(cur) - 1; k > 0; k, index = k - 1, index + 1 {
			cur[k], cur[k-1] = cur[k-1], cur[k]
			result[index] = append(result[index], cur...)
		}
	}


}
