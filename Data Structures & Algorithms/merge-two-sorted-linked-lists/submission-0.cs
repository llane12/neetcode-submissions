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
    public ListNode MergeTwoLists(ListNode list1, ListNode list2) {
        ListNode dummy = new ListNode(0);
        ListNode node = dummy;

        while (list1 != null && list2 != null)
        {
            if (list1.val < list2.val)
            {
                node.next = list1;
                list1 = list1.next;
            }
            else
            {
                node.next = list2;
                list2 = list2.next;
            }
            node = node.next;
        }

        if (list1 != null)
        {
            node.next = list1;
        }
        else
        {
            node.next = list2;
        }

        return dummy.next;
    }

    private ListNode FindNextNode(ListNode list1, ListNode list2)
    {
        ListNode current;
        if (list1 == null)
            current = list2;
        else if (list2 == null)
            current = list1;
        else if (list2.val > list1.val)
            current = list1;
        else
            current = list2;
        return current;
    }
}