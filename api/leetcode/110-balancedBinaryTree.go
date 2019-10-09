package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {

	return searchTreeLength(root) != -1
}

func searchTreeLength(root *TreeNode) int {

	if root == nil {
		return 0
	}

	left := searchTreeLength(root.Left)
	if left == -1 {
		return -1
	}
	right := searchTreeLength(root.Right)
	if right == -1 {
		return -1
	}

	if math.Abs(float64(left-right)) < 2 {
		if left >= right {
			return left + 1
		}
		return right + 1
	}

	return -1
}
