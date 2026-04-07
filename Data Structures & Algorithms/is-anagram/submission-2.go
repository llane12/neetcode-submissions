func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make([]int, 26)
	for i := 0; i < len(s); i++ {
        counts[s[i] -'a']++
        counts[t[i] -'a']--
    }
    for _, v := range counts {
        if v != 0 {
            return false
        }
    }
	return true
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := map[rune]int{}
	for _, c := range s {
		counts[c]++
	}
	for _, c := range t {
		counts[c]--
		if counts[c] < 0 {
			return false
		}
	}
	return true
}
