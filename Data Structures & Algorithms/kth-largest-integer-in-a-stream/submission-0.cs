public class KthLargest {

    private readonly int _k;
    private readonly PriorityQueue<int, int> _minHeap;

    public KthLargest(int k, int[] nums) {
        _k = k;

        _minHeap = new PriorityQueue<int, int>();
        for (int i = 0; i < nums.Length; i++)
        {
            _minHeap.Enqueue(nums[i], nums[i]);
        }
        while (_minHeap.Count > k)
        {
            _minHeap.Dequeue();
        }
    }
    
    public int Add(int val) {
        _minHeap.Enqueue(val, val);
        if (_minHeap.Count > _k)
        {
            _minHeap.Dequeue();
        }
        return _minHeap.Peek();
    }
}
