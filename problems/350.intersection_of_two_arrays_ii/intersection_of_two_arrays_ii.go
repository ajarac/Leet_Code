package intersection_of_two_arrays_ii

func intersect(nums1 []int, nums2 []int) []int {
	map1 := buildMap(nums1)
	map2 := buildMap(nums2)

	var result []int
	for num, count1 := range map1 {
		if count2, ok := map2[num]; ok == true {
			min := Min(count1, count2)
			for i := 0; i < min; i++ {
				result = append(result, num)
			}
		}
	}

	return result
}

func buildMap(nums1 []int) map[int]int {
	m := make(map[int]int)
	for _, num := range nums1 {
		i, ok := m[num]
		if ok == true {
			m[num] = i + 1
		} else {
			m[num] = 1
		}
	}
	return m
}

func Min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
