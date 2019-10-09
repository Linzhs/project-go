package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		depth++
		elemNum := len(queue)
		for i := 0; i < elemNum; i++ {
			elem := queue[i]
			if elem.Left == nil && elem.Right == nil {
				return depth
			}
			if elem.Left != nil {
				queue = append(queue, elem.Left)
			}
			if elem.Right != nil {
				queue = append(queue, elem.Right)
			}
		}
		queue = queue[elemNum:]
	}

	return depth
}
