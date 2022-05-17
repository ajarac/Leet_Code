package climbing_stairs

func climbStairs(n int) int {
	if n < 2 {
		return n
	}
	var a uint32 = 1
	var b uint32 = 1
	var ret uint32 = 0

	for i := 2; i <= n; i++ {
		ret = a + b
		b = a
		a = ret
	}
	return int(ret)
}
