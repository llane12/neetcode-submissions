public class Solution {
    public int MaxProfit(int[] prices) {
        if (prices.Length < 2) return 0;

        int maxProfit = 0;
        int buyPrice = int.MaxValue;

        for (int i = 0; i < prices.Length; i++)
        {
            if (prices[i] < buyPrice)
            {
                buyPrice = prices[i];
            }
            else
            {
                int profit = prices[i] - buyPrice;
                if (profit > maxProfit)
                {
                    maxProfit = profit;
                }
            }
        }

        return maxProfit;
    }
}
