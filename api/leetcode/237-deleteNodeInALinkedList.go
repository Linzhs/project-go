package leetcode

//Note:
//
//The linked list will have at least two elements.
//All of the nodes' values will be unique.
//The given node will not be the tail and it will always be a valid node of the linked list.
//Do not return anything from your function.
func deleteNode(node *ListNode) {
	// 将下一个节点的信息拷贝到当前节点
	// 此时当前节点和其下一个节点相同
	// 删除下一个节点
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
