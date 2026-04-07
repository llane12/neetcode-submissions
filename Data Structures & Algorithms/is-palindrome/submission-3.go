func isPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	runes := []rune(s)
	l, r := 0, len(runes) - 1

	for l < r {
		if !isAlphaNumeric(runes[l]) {
			l++
			continue
		}
		
		if !isAlphaNumeric(runes[r]) {
			r--
			continue
		}

		if unicode.ToLower(runes[l]) != unicode.ToLower(runes[r]) {
			return false
		}

		l++
		r--
	}
	return true
}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
