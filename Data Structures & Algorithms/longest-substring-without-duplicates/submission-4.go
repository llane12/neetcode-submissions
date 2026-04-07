func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {return 0}
	if len(s) == 1 {return 1}
	runes := []rune(s)

	maxL := 0
	substr := make(map[rune]int)
	substr[runes[0]] = 0
	for l, r := 0, 1; r < len(runes); r++ {
		char := runes[r]
		if idx, found := substr[char]; found {
			for ; l < idx; l++ {
				delete(substr, runes[l])
			}
			l++
		}
		substr[char] = r
		maxL = max(maxL, r - l + 1)
	}
	return maxL
}
