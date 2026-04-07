func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for _, str := range strs {
		counts := countCharacters(str)
		group, found := groups[counts]
		if found {
			groups[counts] = append(group, str)
		} else {
			groups[counts] = []string{str}
		}
	}

	result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
	return result
}

func countCharacters(s string) [26]int {
    counts := [26]int{}
    for _, c := range s {
		counts[c - 'a']++
    }
    return counts
}
