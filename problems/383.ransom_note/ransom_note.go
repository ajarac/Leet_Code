package ransom_note

func canConstruct(ransomNote string, magazine string) bool {
	tmpRansom := make([]int, 26)
	tmpMagazine := make([]int, 26)
	for _, v := range ransomNote {
		tmpRansom[v-'a']++
	}
	for _, v := range magazine {
		tmpMagazine[v-'a']++
	}
	for i := range tmpRansom {
		if tmpRansom[i] > tmpMagazine[i] {
			return false
		}
	}
	return true
}
