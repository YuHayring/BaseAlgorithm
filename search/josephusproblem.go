package search

//O(n) space:O(1)
func lastRemaining(n int, m int) int {
	f := 0;
	for i := 2; i != n + 1; i++ {
		f = (m + f) % i
	}
	return f
}