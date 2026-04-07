public class Solution {
    public int MinEatingSpeed(int[] piles, int h) {
        int l = 1, r = piles.Max();
        int k = 0;

        while (l <= r)
        {
            int m = l + (r - l) / 2;

            long hours = 0;
            foreach (int pile in piles)
            {
                hours += (long)Math.Ceiling((double)pile / m);                
            }

            if (hours <= h)
            {
                k = m;
                r = m - 1;
            }
            else
            {
                l = m + 1;
            }
        }

        return k;
    }
}
