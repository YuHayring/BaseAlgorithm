package str

func Kmp(src string, target string) int {
	next := NextKmp(target)
	i, j := 0, 0
	for i < len(src) && j < len(target) {
		if src[i] == target[j] {
			j++
		} else {
			if j != 0 {
				j = next[j-1]
				continue
			}
		}
		i++
	}
	if j == len(target) {
		return i - j
	}
	return -1
}

func NextKmp(target string) []int {
	next := make([]int, len(target))
	k := 0
	for i := 1; i < len(target); i++ {
		if target[k] == target[i] {
			k++
			next[i] = k
		} else {
			if k > 0 {

				k = next[k - 1]

				if target[k] == target[i] {
					k++
				}
				next[i] = k
			}
		}
	}
	return next
}