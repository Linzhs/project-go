package leetcode

func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	curr, next := head, head.Next
	for ; next.Next != nil && next.Next.Next != nil; curr, next = next.Next, next.Next.Next {
		swapVal(curr, next)
	}
	swapVal(curr, next)

	return head
}

func swapVal(cur *ListNode, other *ListNode) {
	intVal := cur.Val
	cur.Val = other.Val
	other.Val = intVal
}

func createLinkedList(intSlice []int) *ListNode {

	head := &ListNode{}
	p := head
	for _, item := range intSlice {
		p.Val = item
		p.Next = &ListNode{}
		p = p.Next
	}

	return head
}
