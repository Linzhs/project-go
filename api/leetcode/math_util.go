package leetcode

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func absInt64(x int64) int64 {
	y := x >> 63
	return (x ^ y) - y
}
