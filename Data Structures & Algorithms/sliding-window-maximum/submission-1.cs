public class Solution {
    public int[] MaxSlidingWindow(int[] nums, int k) {
       if (nums.Length == 0) return [];

        int[] res = new int[nums.Length - k + 1];
        LinkedList<int> deque = new();

        // Pre-fill the first window
        for (int i = 0; i < k; i++)
        {
            while (deque.Count > 0 && nums[deque.Last!.Value] < nums[i])
                deque.RemoveLast();
            deque.AddLast(i);
        }

        for (int l = 0; l <= nums.Length - k; l++)
        {
            int r = l + k - 1;

            // Record max for this window
            res[l] = nums[deque.First!.Value];

            // Remove l if it's the front (it's leaving the window)
            if (deque.First!.Value == l)
                deque.RemoveFirst();

            // Add the next element if there is one
            if (r + 1 < nums.Length)
            {
                while (deque.Count > 0 && nums[deque.Last!.Value] < nums[r + 1])
                    deque.RemoveLast();
                deque.AddLast(r + 1);
            }
        }

        return res;
    }

    public int[] MaxSlidingWindow_MaxHeap(int[] nums, int k) {
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
