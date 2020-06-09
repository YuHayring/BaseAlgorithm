package str

func BoyerMoore(src string, target string) int {
	end := len(target) - 1
	right := end
	for right < len(src) {
		left := right
		e := end
		for e > -1 && src[left] == target[e] {
			left, e = left - 1, e - 1
		}
		if e == -1 {
			return left + 1
		}
		for e > -1 && src[left] != target[e] {
			e--
		}
		right = left + end - e

	}
	return -1
}
