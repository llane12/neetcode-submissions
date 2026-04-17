/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */

public class Solution {    
    public ListNode MergeKLists(ListNode[] lists) {
        ListNode head = new ListNode();
        ListNode cur = head;

        PriorityQueue<ListNode, int> q = new();
        foreach (ListNode n in lists)
        {
            q.Enqueue(n, n.val);
        }

        while (q.Count > 0)
        {
            ListNode n = q.Dequeue();
            cur.next = n;
            cur = n;
            if (cur.next != null)
            {
                q.Enqueue(cur.next, cur.next.val);
            }
        }

        return head.next;
    }
}
