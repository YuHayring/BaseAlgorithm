package permutationgenerate

import obj "../basestruct"

func BackTracking(src []int) [][]int {
	length := len(src)
	if length == 0 {
		return [][]int{}
	}
	for i := length - 1; i > 1; i-- {
		length *= i
	}
	result := make([][]int, 0, length)
	queue := obj.BuildQueueByIntArray(src)
	cur := make([]int, 0, len(src))
	for i := 0; i < len(src); i++ {
		e := queue.Pull().(int)
		cur = append(cur, e)
		backTrackingCore(queue, &result, cur)
		cur = cur[:len(cur)-1]
		queue.Push(e)
	}
	return result
}

func backTrackingCore(queue *obj.Queue, result *[][]int, cur []int) {
	if queue.Len() == 0 {
		res := make([]int,len(cur))
		copy(res, cur)
		*result = append(*result, res)
	}
	for i := 0; i < queue.Len(); i++ {
		e := queue.Pull().(int)
		cur = append(cur, e)
		backTrackingCore(queue, result, cur)
		cur = cur[:len(cur)-1]
		queue.Push(e)
	}
}