package leetcode

import "strings"

func convert(s string, numRows int) string {

	if len(s) <= 1 || numRows > len(s) {
		return s
	}

	//interval := numRows - 2
	build := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if i == 0 {
			build.WriteRune(rune(s[i]))
		}
	}

	return build.String()
}
