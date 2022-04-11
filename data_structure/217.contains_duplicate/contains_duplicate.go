package contains_duplicate

func containsDuplicate(nums []int) bool {
	s := map[int]bool{}
	for _, num := range nums {
		_, ok := s[num]
		if ok {
			return true
		} else {
			s[num] = true
		}
	}
	return false
}
