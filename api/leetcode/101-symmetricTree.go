package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	//return isSymmetricUsingRecursive(root, root)
	return isSymmetricUsingIterative(root)
}

func isSymmetricUsingRecursive(node1 *TreeNode, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.Val == node2.Val &&
		isSymmetricUsingRecursive(node1.Left, node2.Right) &&
		isSymmetricUsingRecursive(node1.Right, node2.Left)
}

func isSymmetricUsingIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}

	queue := []*TreeNode{root, root}
	for len(queue) != 0 {
		n1 := queue[0]
		n2 := queue[1]
		queue = queue[2:]
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		queue = append(queue, n1.Left)
		queue = append(queue, n2.Right)
		queue = append(queue, n1.Right)
		queue = append(queue, n2.Left)
	}

	return true
}
