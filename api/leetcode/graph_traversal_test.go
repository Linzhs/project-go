package leetcode

import "testing"

func TestBfs(t *testing.T) {
	tree := newTree([]int{10, 5, 11, 2, 8, 19, 21, 1, 3, 7, 9})
	bfs(tree)
}
