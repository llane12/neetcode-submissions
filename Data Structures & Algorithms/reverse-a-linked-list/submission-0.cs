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
    public ListNode ReverseList(ListNode head) {
        ListNode prev = null, cur = head != null ? new ListNode(head.val, head.next) : null;
        while (cur != null)
        {
            ListNode tmp = cur.next != null ? new ListNode(cur.next.val, cur.next.next) : null;
            cur.next = prev;
            prev = cur;
            cur = tmp;
        }
        return prev;
    }
}
