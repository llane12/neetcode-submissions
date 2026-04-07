func groupAnagrams(strs []string) [][]string {
	groups := make(map[int][]string)

	for _, str := range strs {
		hash := hashCode(str)
		group, found := groups[hash]
		if found {
			groups[hash] = append(group, str)
		} else {
			groups[hash] = []string{str}
		}
	}

	result := make([][]string, 0, len(groups))
    for _, group := range groups {
        result = append(result, group)
    }
	return result
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"
var primes = []int{3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

func hashCode(s string) int {
    sum := 0
    for _, c := range s {
      sum += primes[c - 'a']
    }
    return sum
}
