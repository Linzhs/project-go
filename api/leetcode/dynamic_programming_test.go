package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountPaths(t *testing.T) {
	grid := [][]bool{
		{false, false, false, false, false, false, false, false},
		{false, false, true, false, false, false, true, false},
		{false, false, false, false, true, false, false, false},
		{true, false, true, false, false, true, false, false},
		{false, false, true, false, false, false, false, false},
		{false, false, false, true, true, false, true, false},
		{false, true, false, false, false, true, false, false},
		{false, false, false, false, false, false, false, false},
	}

	fmt.Println(countPathsV2(grid))
}

func TestMinimumTotal(t *testing.T) {
	triangle := [][]int{
		{2},
		{3, 3},
		{6, 5, 7},
		{4, 1, 8, 3},
	}
	fmt.Println(minimumTotalV1(triangle))
	fmt.Println(minimumTotalV2(triangle))
}

func TestMaxProduct(t *testing.T) {
	nums := []int{2, 3, -2, 4}
	got := maxProduct(nums)
	assert.Equal(t, 6, got)
}

func TestCoinsChange(t *testing.T) {
	tests := []struct {
		arr    []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, coinChange(test.arr, test.amount))
	}
}

func TestLargestRectangleArea(t *testing.T) {
	s := []int{2, 1, 5, 6, 2, 3}
	fmt.Println(largestRectangleArea(s))
}

func TestLIS(t *testing.T) {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}
