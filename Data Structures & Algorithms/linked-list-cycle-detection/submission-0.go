/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

    fast, slow := head, head
	for {		
		slow = slow.Next
		fast = fast.Next
		if fast == nil {
			return false
		}

		fast = fast.Next
		if fast == nil {
			return false
		} else if fast == slow {
			return true
		}
	}
}
