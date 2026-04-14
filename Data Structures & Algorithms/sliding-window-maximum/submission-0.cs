public class Solution {
    public int[] MaxSlidingWindow(int[] nums, int k) {
       if (nums.Length == 0) return new int[0];
        
        int[] res = new int[nums.Length - k + 1];

        // Max heap
        var comparer = Comparer<int>.Create((x, y) => y.CompareTo(x));
        PriorityQueue<int, int> window = new(comparer);

        int r = 0;
        for (r = 0; r < k; r++)
        {
            window.Enqueue(r, nums[r]);
        }

        for (int l = 0; l <= nums.Length - k; l++)
        {
            int i = window.Peek();
            res[l] = nums[i];

            if (i == l)
            {
                window.Dequeue();
            }
            else
            {
                window.Remove(l, out _, out _);
            }

            if (r < nums.Length)
            {
                window.Enqueue(r, nums[r]);
            }
            r++;
        }

        return res;
    }
}
