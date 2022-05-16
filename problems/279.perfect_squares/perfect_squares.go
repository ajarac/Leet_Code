package perfect_squares

import "leetcode/utils"

func numSquares(n int) int {
	dp := make([]int, n+1)

	for i := range dp {
		dp[i] = n + n
	}
	dp[0] = 0

	for i := 0; i <= n; i++ {
		for j := 0; j*j < i+1; j++ {
			square := j * j
			if i-square >= 0 {
				num1 := dp[i]
				num2 := dp[i-square] + 1
				dp[i] = utils.Min(num1, num2)
			}
		}
	}

	return dp[n]
}
