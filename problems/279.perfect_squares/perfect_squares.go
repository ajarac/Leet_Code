package perfect_squares

import "leetcode/utils"

/*
   const dp = new Array(n + 1).fill(Infinity)
   dp[0] = 0
   for (let i = 0; i <= n; i++) {
       for (let j = 0; j * j < i + 1; j++) {
           const square = j * j
           if (i - square >= 0) {
               dp[i] = Math.min(dp[i], dp[i - square] + 1)
           }
       }
   }
   return dp[n]
*/
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
