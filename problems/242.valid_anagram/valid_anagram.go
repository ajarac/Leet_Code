package valid_anagram

func isAnagram(s string, t string) bool {
	tmpRansom := make([]int, 26)
	tmpMagazine := make([]int, 26)
	for _, v := range s {
		tmpRansom[v-'a']++
	}
	for _, v := range t {
		tmpMagazine[v-'a']++
	}
	for i := range tmpRansom {
		if tmpRansom[i] != tmpMagazine[i] {
			return false
		}
	}
	return true
}
