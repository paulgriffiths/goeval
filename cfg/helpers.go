package cfg

func intRange(n int) []int {
	r := []int{}
	for i := 0; i < n; i++ {
		r = append(r, i)
	}
	return r
}
