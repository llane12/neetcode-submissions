func climbStairs(n int) int {
	cache := make([]int, n+1)
	for i := 0; i <= n; i++ {
        cache[i] = -1
    }

    return climbStairsR(n, 0, cache)
}

func climbStairsR(n, i int, cache []int) int {
	if i == n {
		return 1
	} else if i > n {
		return 0
	} else if cache[i] != -1 {
		return cache[i]
	}

	cache[i] = climbStairsR(n, i+1, cache) + climbStairsR(n, i+2, cache)
	return cache[i]	
}
