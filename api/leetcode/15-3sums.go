package leetcode

import (
	"sort"
)

type intSlice []int

func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// target = a + b + c
func threeSum(nums []int) [][]int {
	if len(nums) < 2 {
		return [][]int{}
	}

	var ret [][]int
	slice := intSlice(nums)
	sort.Sort(slice)
	for i, v := range slice {
		if i > 0 && slice[i] == slice[i-1] {
			continue
		}
		for j, k := i+1, len(slice)-1; j < k; {
			sum := v + slice[j] + slice[k]
			if sum < 0 {
				j++
			} else if sum > 0 {
				k--
			} else {
				ret = append(ret, []int{slice[i], slice[j], slice[k]})
				for j < k && slice[j] == slice[j+1] {
					j++
				}
				for j < k && slice[k] == slice[k-1] {
					k--
				}
				j++
				k--
			}
		}
	}

	return ret
}
