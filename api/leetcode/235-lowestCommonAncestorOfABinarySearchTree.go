package leetcode

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	leftNode := lowestCommonAncestor(root.Left, p, q)
	rightNode := lowestCommonAncestor(root.Right, p, q)

	if leftNode == nil {
		return rightNode
	} else if rightNode == nil {
		return leftNode
	}

	return root
}
