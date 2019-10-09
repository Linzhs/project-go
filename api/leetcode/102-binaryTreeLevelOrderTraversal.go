package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var ret [][]int
	queue := []*TreeNode{root}
	for len(queue) != 0 { // 针对每一层做处理
		var t []int
		for i, l := 0, len(queue); i < l; i++ {
			top := queue[0]
			t = append(t, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			switch len(queue) {
			case 0:
				queue = []*TreeNode{}
			default:
				queue = queue[1:]
			}
		}
		ret = append(ret, t)
	}

	return ret
}
