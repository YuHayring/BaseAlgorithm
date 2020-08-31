package str

func Horspool(pattern, target string) int {
	if len(pattern) > len(target) {
		return -1
	}
	table := getShiftTable(pattern)
	i := len(pattern) - 1
	str := []byte(target)
	pat := []byte(pattern)
	pLen := len(pattern)
	for i < len(target) {
		k := 0
		for  k < len(pattern) && pat[pLen - 1 - k] == str[i-k] {
			k++
		}
		if k == pLen {
			return i - pLen + 1
		} else {
			i += table[str[i] - 'A']
		}
	}
	return -1
}

func getShiftTable(pattern string) []int {
	result := make([]int,26)
	length := len(pattern)
	str := []byte(pattern)
	for i, _ := range result {
		result[i] = length
	}
	for i := 0; i < len(pattern) - 1; i++ {
		result[str[i]-'A'] =  length - 1 - i
	}
	return result
}
