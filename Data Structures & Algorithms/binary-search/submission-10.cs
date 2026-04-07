public class Solution {
    public int Search(int[] nums, int target) {
        if (nums.Length == 0)
        {
            return nums[0] == target ? 0 : -1;
        }

        int l = 0, r = nums.Length - 1;
        int m;

        while (l <= r)
        {
            m = l + (r - l) / 2;
            if (nums[m] == target)
                return m;
            else if (nums[m] < target)
                l = m + 1;
            else
                r = m - 1;
        }
        return -1;
    }
}
