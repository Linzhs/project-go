package leetcode

import "sort"

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	slice := intSlice(nums)
	sort.Sort(slice)

	return slice[len(slice)/2]
}
