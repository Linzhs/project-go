package leetcode

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	cur := root
	for cur != nil {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
			continue
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
			continue
		}
		return cur
	}
	return nil
}
