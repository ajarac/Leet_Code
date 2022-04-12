package two_sums

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		wanted := target - num
		if index, ok := numMap[num]; ok == true {
			return []int{index, i}
		}
		numMap[wanted] = i
	}
	return []int{}
}
