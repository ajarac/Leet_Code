package squares_of_a_sorted_array

func sortedSquares(nums []int) []int {
	total := len(nums)
	if total == 0 {
		return []int{}
	}

	var indexNegative int = -1
	var indexPositive int = total

	for i, num := range nums {
		if num >= 0 {
			indexNegative = i - 1
			indexPositive = i
			break
		}
	}
	if indexPositive == total {
		indexNegative = total - 1
	}
	result := make([]int, total)
	index := 0

	for indexNegative > -1 || indexPositive < total {

		if indexPositive == total && (indexNegative > -1) || (indexPositive < total && indexNegative > -1 && (nums[indexNegative]*-1) < nums[indexPositive]) {
			result[index] = nums[indexNegative] * nums[indexNegative]
			indexNegative--
		} else if indexNegative == -1 && (indexPositive < total) || (nums[indexNegative]*-1) >= nums[indexPositive] {
			result[index] = nums[indexPositive] * nums[indexPositive]
			indexPositive++
		}
		index++
	}

	return result
}
