package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 1->2->3->4->5->nil
// 5->4->3->2->1->nil
func ReverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	var prev, curr, next *ListNode = nil, head, head.Next
	for next != nil {
		// 反转当前节点指针方向
		curr.Next = prev

		// 移动3指针
		prev = curr
		curr = next
		next = next.Next
	}

	curr.Next = prev

	return curr
}
