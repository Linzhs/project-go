package leetcode

import (
	"testing"
)

func TestCreateTree(t *testing.T) {
	root := createTreeByRecursive(0, []int{1, 3, 5, 4, 6, 9})
	prevOrder(root)
	inOrder(root)
	postOrder(root)
}
