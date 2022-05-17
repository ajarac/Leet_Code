package min_cost_climbing_stairs

import "leetcode/utils"

func minCostClimbingStairs(cost []int) int {
	for i := len(cost) - 2; i > 0; i-- {
		cost[i-1] += utils.Min(cost[i], cost[i+1])
	}
	return utils.Min(cost[0], cost[1])
}
