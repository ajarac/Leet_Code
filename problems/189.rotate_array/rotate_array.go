package rotate_array

func rotate(nums []int, k int) []int {
	if k > len(nums) {
		k = k % len(nums)
	}
	divisionPoint := len(nums) - 1 - k
	return append(nums[divisionPoint+1:], nums[:divisionPoint+1]...)
}
