func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) { return false }

	var s1cc, s2cc [26]int
	for _, char := range s1 {
		s1cc[char - 'a']++
	}

	l := -len(s1)
	for _, char := range s2 {
		s2cc[char-'a']++
		if l >= 0 {
			s2cc[s2[l]-'a']--
		}
		if s1cc == s2cc { return true }
		l++
	}

	return false
}


func checkInclusion1(s1 string, s2 string) bool {
	if len(s2) < len(s1) { return false }

	var charCounts [26]int
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

func perm(cc [26]int) bool {
	for _, count := range cc {
		if count > 0 { return false }
	}
	return true
}
