public class Solution {
    public int MaxArea(int[] heights) {
        int l = 0, r = heights.Length - 1;
        int area = 0;

        while (l < r)
        {
            int curArea = Math.Min(heights[r], heights[l]) * (r - l);
            if (curArea > area)
            {
                area = curArea;
            }

            if (heights[l] < heights[r])
            {
                l++;
            }
            else
            {
                r--;
            }
        }

        return area;
    }
}
