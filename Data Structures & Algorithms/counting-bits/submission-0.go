func countBits(n int) []int {
	res := make([]int, n+1)
	for i := 0; i <= n; i++ {
		res[i] = hammingWeight(i)
	}
	return res
}

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

