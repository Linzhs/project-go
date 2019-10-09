package leetcode

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func newTree(treeList []int) *TreeNode {

	root := &TreeNode{}
	for _, v := range treeList {
		root = treeInsert(root, v)
	}
	return root
}

func treeInsert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{Val: val}
	}

	if val < node.Val {
		node.Left = treeInsert(node.Left, val)
	} else {
		node.Right = treeInsert(node.Right, val)
	}
	return node
}

func prevOrder(root *TreeNode) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		fmt.Print(cur.Val, " ")
		stack = stack[:len(stack)-1]

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	fmt.Println()
}

// 使用栈存储左儿子做为历史记录， 遍历到最左边
func inOrder(root *TreeNode) {
	cur := root
	var stack []*TreeNode
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		if len(stack) != 0 {
			top := stack[len(stack)-1]
			fmt.Print(top.Val, " ")
			stack = stack[:len(stack)-1]
			cur = top.Right
		}
	}
	fmt.Println()
}

func postOrder(root *TreeNode) {
	cur := root
	var stack []*TreeNode
	var result []int
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			result = append(result, cur.Val)
			cur = cur.Right
		} else {
			top := stack[len(stack)-1]
			cur = top.Left
			stack = stack[:len(stack)-1]
		}
	}

	for i := len(result) - 1; i >= 0; i-- {
		fmt.Print(result[i], " ")
	}
	fmt.Println()
}

func createTreeByRecursive(i int, nums []int) *TreeNode {
	t := &TreeNode{Val: nums[i]}
	if i < len(nums) && i*2+1 < len(nums) {
		t.Left = createTreeByRecursive(2*i+1, nums)
	}
	if i < len(nums) && 2*i+2 < len(nums) {
		t.Right = createTreeByRecursive(2*i+2, nums)
	}
	return t
}

func createTreeByIteration() *TreeNode {
	return nil
}
