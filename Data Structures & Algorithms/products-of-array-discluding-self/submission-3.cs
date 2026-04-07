public class Solution {
    public int[] ProductExceptSelf2(int[] nums) {
        int n = nums.Length;
        int[] prefix = new int[n], suffix = new int[n], output = new int[n];

        prefix[0] = nums[0];
        for (int i = 1; i < n; i++)
        {
            prefix[i] = prefix[i - 1] * nums[i];
        }

        suffix[n - 1] = nums[n - 1];
        for (int i = n - 2; i >= 0; i--)
        {
            suffix[i] = suffix[i + 1] * nums[i];
        }

        for (int i = 0; i < n; i++)
        {
            int left = i > 0 ? prefix[i - 1] : 1;
            int right = i < n - 1 ? suffix[i + 1] : 1;
            output[i] = left * right;
        }

        return output;
    }

    public int[] ProductExceptSelf(int[] nums) {
        int n = nums.Length;
        int[] output = new int[n];

        // Left products
        output[0] = 1;
        for (int i = 1; i < n; i++) {
            output[i] = output[i - 1] * nums[i - 1];
        }

        // Right products
        int right = 1;
        for (int i = n - 1; i >= 0; i--) {
            output[i] *= right;
            right *= nums[i];
        }

        return output;
    }
}
