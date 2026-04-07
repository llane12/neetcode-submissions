func minWindow(s string, t string) string {
    if len(s) < len(t) { return "" }

	cc := make(map[rune]int)
	for _, char := range t {
		cc[char]++
	}

	res := ""
	l := 0
	for r, char := range s {
		countChar(cc, char, -1)
		for substring(cc) {
			if res == "" || (r - l + 1) < len(res) {
				res = s[l:r+1]
			}
			countChar(cc, rune(s[l]), 1)
			l++
		}
	}

	return res
}

func countChar(cc map[rune]int, char rune, val int) {
	if _, ok := cc[char]; ok {
		cc[char] += val
	}
}

func substring(cc map[rune]int) bool {
	for _, v := range cc {
		if v > 0 { return false }
	}
	return true
}
