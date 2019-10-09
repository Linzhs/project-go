package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
// Output: 7 -> 0 -> 8
// Explanation: 342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var (
		head          = &ListNode{}
		curNode, q, p = head, l1, l2
		carry         int
	)

	for p != nil || q != nil {
		var val1, val2 int
		if p != nil {
			val1 = p.Val
			p = p.Next
		}
		if q != nil {
			val2 = q.Val
			q = q.Next
		}

		sum := val1 + val2 + carry
		curNode.Next = &ListNode{
			Val: sum % 10,
		}
		curNode = curNode.Next
		carry = sum % 10
	}

	if carry != 0 {
		curNode.Next = &ListNode{
			Val: carry,
		}
	}

	return head.Next
}
