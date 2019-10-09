package leetcode

import (
	"fmt"
	"testing"
)

func TestMaxSidingWindow(t *testing.T) {
	nums := []int{7, 2, 4}
	fmt.Println(maxSlidingWindow(nums, 2))
}
