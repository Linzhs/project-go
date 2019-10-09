package leetcode

// time: 16ms
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	set := map[*ListNode]bool{}
	next := head
	for next != nil {
		_, ok := set[next]
		if ok {
			return true
		}
		set[next] = true
		next = next.Next
	}

	return false
}

// 龟兔赛跑法
// time:14ms
func hasCycleV2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	for slow, fast := head, head.Next; slow != fast; {
		if fast == nil || fast.Next == nil {
			return false
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return true
}
