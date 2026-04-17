public class Solution {
    public int SplitArray(int[] nums, int k) {
        int l = 0, r = 0;
        foreach (int num in nums) {
            l = Math.Max(l, num);
            r += num;
        }

        int ans = r;
        while (l <= r)
        {
            int mid = l + ((r - l) / 2);
            if (CanSplit(nums, k, mid))
            {
                ans = mid;
                r = mid - 1;
            }
            else
            {
                l = mid + 1;
            }
        }
        return ans;
    }

    private bool CanSplit(int[] nums, int k, int targetSum)
    {
        int count = 1;
        int curSum = 0;
        for (int i = 0; i < nums.Length; i++)
        {
            if (curSum + nums[i] > targetSum)
            {
                count++;
                curSum = nums[i];
            }
            else
            {
                curSum += nums[i];
            }
        }
        return count <= k;
    }
}