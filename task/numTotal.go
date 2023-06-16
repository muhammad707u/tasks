package task

func NumTotal(a int, b int, c int) int {
	var res1, res2 int
	for i := a; i < b; i++ {
		res1 = res1 + i
	}
	for j := b; j <= c; j++ {
		res2 = res2 + j
	}
	result := res1 + res2
	return result
}
