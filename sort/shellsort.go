package sort

func Shellsort(src []int) {
	//外层步长控制
	for step := len(src) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(src); i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && src[j+step] < src[j]; j -= step {
				src[j], src[j+step] = src[j+step], src[j]
			}
		}
	}
}
