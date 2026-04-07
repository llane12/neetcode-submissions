func topKFrequent(nums []int, k int) []int {
	counts := map[int]int{}
	for _, n := range nums {
		counts[n]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range counts {
		buckets[freq] = append(buckets[freq], num)
	}

	result := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		for _, n := range buckets[i] {
			result = append(result, n)
			if len(result) == k {
				return result
			}
		}
	}
	return result
}
