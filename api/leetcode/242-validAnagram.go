package leetcode

import (
	"reflect"
)

// 也可以使用排序的方法
func isAnagram(s string, t string) bool {

	set1 := make(map[rune]int)
	for _, v := range s {
		set1[v]++
	}

	set2 := make(map[rune]int)
	for _, v := range t {
		set2[v]++
	}

	return reflect.DeepEqual(set1, set2)
}
