import "slices"

func lengthOfLIS(nums []int) int {
	n := len(nums)
    dp := make([]int, n)
	parent := make([]int, n)

	lis := 0
	lisIdx := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		parent[i] = -1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j] + 1 > dp[i] {
					dp[i] = dp[j] + 1
					parent[i] = j
				}
			}
		}
		if dp[i] > lis {
			lis = dp[i]
			lisIdx = i
		}
	}

	seq := make([]int, lis)
	p := lisIdx
	for p > -1 {
		seq = append(seq, nums[p])
		p = parent[p]
	}
	slices.Reverse(seq) 

	return lis
}
