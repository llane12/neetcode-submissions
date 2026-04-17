public class Solution {
    public int MaxProfit(int[] prices) {
        if (prices.Length == 0) return 0;

        // Initial values for Day 0
        int holding = -prices[0]; // Profit after buying on day 0
        int sold = 0;             // Profit after selling (impossible on day 0)
        int free = 0;             // Profit while being free to buy

        for (int i = 1; i < prices.Length; i++) {
            int prevHolding = holding;
            int prevSold = sold;
            int prevFree = free;

            // 1. To be HOLDING: Either you held from yesterday, 
            //    OR you were FREE and you BUY today.
            holding = Math.Max(prevHolding, prevFree - prices[i]);

            // 2. To be in SOLD: You must have been HOLDING and you SELL today.
            sold = prevHolding + prices[i];

            // 3. To be FREE: Either you were already FREE, 
            //    OR you just finished your COOLDOWN (you were in SOLD yesterday).
            free = Math.Max(prevFree, prevSold);
        }

        // Your max profit will be either you just sold, or you're free.
        return Math.Max(sold, free);
    }
}
