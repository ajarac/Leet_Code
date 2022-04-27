package container_with_most_water

import "leetcode/utils"

func maxArea(height []int) int {

	left, right := 0, len(height)-1
	max := 0

	for left < right {
		area := utils.Min(height[left], height[right]) * (right - left)
		max = utils.Max(max, area)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}
