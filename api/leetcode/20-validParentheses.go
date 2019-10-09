package leetcode

import "strings"

func isValid(s string) bool {

	for length := 0; len(s) != length; {
		length = len(s)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "()", "", -1)
		s = strings.Replace(s, "[]", "", -1)
	}
	return len(s) == 0
}
