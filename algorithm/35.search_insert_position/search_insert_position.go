package search_insert_position

func searchInsert(nums []int, target int) int {

	if target < nums[0] {
		return 0
	}
	size := len(nums)
	if target > nums[size-1] {
		return size
	}

	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}

	return left
}
