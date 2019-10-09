package leetcode

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	nums := []int{6, 5, 3, 1, 8, 7, 2, 4}
	fmt.Printf("%+v", bubbleSort(nums))
}
