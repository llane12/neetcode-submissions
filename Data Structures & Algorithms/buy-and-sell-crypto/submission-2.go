func maxProfit(prices []int) int {
	if len(prices) == 1 { return 0 }

	profit := 0
	l, r := 0, 1
	for l < len(prices) && r < len(prices) {
		profit = max(profit, prices[r] - prices[l])
		if prices[r] > prices[l] {
			r++
		} else {
			l++
			r = max(r, l + 1)
		}
	}
	return profit
}
