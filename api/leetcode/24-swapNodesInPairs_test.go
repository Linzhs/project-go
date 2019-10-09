package leetcode

import "testing"

func TestSwapNodesInPairs(t *testing.T) {
	tests := []struct {
		slice []int
	}{
		{[]int{1, 2, 3, 4}},
	}

	for _, test := range tests {
		head := createLinkedList(test.slice)
		swapPairs(head)
	}
}
