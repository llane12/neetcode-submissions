func hammingWeight(n int) int {
	mask := 1
	ones := 0
	for range 32 {
		if n & mask == mask {
			ones++
		}
		mask = mask << 1
	}
	return ones
}
