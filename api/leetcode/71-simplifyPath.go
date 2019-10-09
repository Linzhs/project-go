package leetcode

import (
	"strings"
)

func simplifyPath(path string) string {

	var outputStrings []string

	pathSlice := strings.Split(path, "/")
	for _, v := range pathSlice {
		switch v {
		case "", ".": // current dir
		case "..": // prev dir
			if len(outputStrings) > 0 {
				outputStrings = outputStrings[:len(outputStrings)-1]
			}
		default:
			outputStrings = append(outputStrings, v)
		}
	}

	ret := "/"
	ret += strings.Join(outputStrings, "/")

	return ret
}
