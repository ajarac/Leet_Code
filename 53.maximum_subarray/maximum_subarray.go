package maximum_subarray

func maxSubArray(nums []int) int {
	currentMax, globalMax := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		currentMax = getMax(nums[i], nums[i]+currentMax)

		globalMax = getMax(currentMax, globalMax)
	}

	return globalMax
}

func getMax(num1 int, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}
