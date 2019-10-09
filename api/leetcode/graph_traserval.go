package leetcode

import "fmt"

func bfs(root *TreeNode) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) != 0 {
		top := queue[0]
		queue = queue[1:]
		fmt.Print(top.Val, " ")
		if top.Left != nil {
			queue = append(queue, top.Left)
		}
		if top.Right != nil {
			queue = append(queue, top.Right)
		}
	}
}

func dfsUsingRec(root *TreeNode) {

	if root == nil {
		return
	}

	fmt.Print(root.Val, " ")
	if root.Left != nil {
		dfsUsingRec(root.Left)
	}
	if root.Right != nil {
		dfsUsingRec(root.Right)
	}
}

func dfsUsingIter(root *TreeNode) {

	cur := root
	var stack []*TreeNode
	for len(stack) != 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
		} else {
			top := stack[len(stack)-1]
			fmt.Println(top.Val)
			stack = stack[:len(stack)-1]
			cur = top.Right
		}
	}
}
