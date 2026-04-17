public class Solution {
    public int MaxProfit(int[] prices)
    {
        Dictionary<(int, bool), int> dp = new();
        return Trade(prices, 0, true, dp);
    }

    private int Trade(int[] prices, int i, bool canBuy, Dictionary<(int, bool), int> dp)
    {
        if (i >= prices.Length)
        {
            return 0;
        }

        if (dp.TryGetValue((i, canBuy), out int p))
        {
            return p;
        }

        int hold = Trade(prices, i + 1, canBuy, dp);

        int step = 2;
        int price = prices[i];
        if (canBuy)
        {
            price = -price;
            step = 1;
        }

        int transact = price + Trade(prices, i + step, !canBuy, dp);

        int profit = Math.Max(transact, hold);
        dp[(i, canBuy)] = profit;
        return profit;
    }
}
