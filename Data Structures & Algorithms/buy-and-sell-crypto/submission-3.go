func maxProfit(prices []int) int {
	if len(prices) == 1 { return 0 }

	profit := 0
	for l, r := 0, 1; r < len(prices); r++ {
		profit = max(profit, prices[r] - prices[l])
		if prices[r] < prices[l] {
			l = r
		}
	}
	return profit
}
