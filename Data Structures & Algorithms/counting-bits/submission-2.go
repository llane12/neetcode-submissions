// Using the pattern of the binary representation of the numbers:
// 0 - 0000
// 1 - 0001 *
// 2 - 0010 *
// 3 - 0011
// 4 - 0100 *
// 5 - 0101
// 6 - 0110
// 7 - 0111
// 8 - 1000 *
// At each new significant bit, the previous bits are repeated
// New significant bits occur at powers of 2
func countBits(n int) []int {
	dp := make([]int, n+1)
	dp[0] = 0

	offset := 1
	for i := 1; i <= n; i++ {
		if offset * 2 == i {
			offset = i
		}
		dp[i] = 1 + dp[i - offset]
	}
	return dp
}
