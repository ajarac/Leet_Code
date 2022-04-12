package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	i1 := m - 1
	i2 := n - 1
	for i := len(nums1) - 1; i >= 0; i-- {
		if (i1 == -1 && i2 >= 0) || (i1 >= 0 && i2 >= 0 && nums1[i1] <= nums2[i2]) {
			nums1[i] = nums2[i2]
			i2--
		} else if (i2 == -1 && i1 >= 0) || (i2 >= 0 && i1 >= 0 && nums2[i2] <= nums1[i1]) {
			nums1[i] = nums1[i1]
			i1--
		}
	}
	if m == 0 {
		nums1 = nums2
	}
	return nums1
}
