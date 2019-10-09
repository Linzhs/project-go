package leetcode

// 二叉搜索树中序遍历有序
func isValidBST(root *TreeNode) bool {

	cur := root
	var stack []*TreeNode
	var sortSlice []int
	for cur != nil || len(stack) != 0 {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = top.Right
			if len(sortSlice) != 0 && top.Val < sortSlice[len(sortSlice)-1] {
				return false
			}
			sortSlice = append(sortSlice, top.Val)
		}
	}

	return true
}
