package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	var result []int
	var stack []*TreeNode

	cur := root
	for cur != nil || len(stack) > 0 {
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
	reverseIntSlice(result)
	return result
}

func reverseIntSlice(s []int) {
	first, last := 0, len(s)-1
	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}
