func minWindow(s string, t string) string {
    if len(s) < len(t) { return "" }

	cc := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		cc[t[i]]++
	}

	required, satisfied := len(cc), 0
	res, resLen := "", math.MaxInt32

	l := 0
	for r := 0; r < len(s); r++ {
		char := s[r]

		if _, ok := cc[char]; ok {
			cc[char]--
			if cc[char] == 0 {
				satisfied++
			}
		}

		for satisfied == required {
			if (r - l + 1) < resLen {
				res = s[l:r+1]
				resLen = r - l + 1
			}
			
			leftChar := s[l]
			if _, ok := cc[leftChar]; ok {
				if cc[leftChar] == 0 {
					satisfied--
				}
				cc[leftChar]++
			}
			l++
		}
	}

	return res
}

// ===================

func minWindow1(s string, t string) string {
    if len(s) < len(t) { return "" }

	cc := make(map[rune]int)
	for _, char := range t {
		cc[char]++
	}

	res, resLen := "", math.MaxInt32
	l := 0
	for r, char := range s {
		countChar(cc, char, -1)
		for substring(cc) {
			if (r - l + 1) < resLen {
				res = s[l:r+1]
				resLen = r - l + 1
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
