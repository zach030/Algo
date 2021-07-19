package bit

func countBits(n int) []int {
	ret := make([]int, n)
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			ret[i] = ret[i/2]
		} else {
			ret[i] = ret[i/2] + 1
		}
	}
	return ret
}
