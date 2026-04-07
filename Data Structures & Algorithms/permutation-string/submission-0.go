func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) { return false }

	charCounts := make([]int, 26)
	for _, char := range s1 {
		charCounts[char - 'a']++
	}

	l := -len(s1)
	for _, char := range s2 {
		charCounts[char-'a']--
		if l >= 0 {
			charCounts[s2[l]-'a']++
		}
		if perm(charCounts) { return true }
		l++
	}

	return false
}

func perm(cc []int) bool {
	for _, count := range cc {
		if count > 0 { return false }
	}
	return true
}
