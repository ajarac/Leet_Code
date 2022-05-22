package subsets

func subsets(nums []int) [][]int {
	var res [][]int
	var dfs func(i int, subset []int)
	dfs = func(i int, subset []int) {
		if i >= len(nums) {
			res = append(res, subset)
			return
		}
		dfs(i+1, append(subset, nums[i]))
		dfs(i+1, subset)
	}
	dfs(0, []int{})
	return res
}
