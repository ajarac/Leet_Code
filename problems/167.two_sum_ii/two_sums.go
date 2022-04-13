package two_sums_2

func twoSum2(numbers []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range numbers {
		wanted := target - num
		if index, ok := numMap[num]; ok == true {
			return []int{index + 1, i + 1}
		}
		numMap[wanted] = i
	}
	return []int{}
}
