package move_zeroes

func moveZeroes(nums []int) {
	zero := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && zero == -1 {
			zero = i
		}

		if nums[i] != 0 && zero != -1 {
			nums[zero] = nums[i]
			nums[i] = 0
			zero++
		}
	}
}
