package leetcode

import "math"

// a^2 + b^2 = c
func JudgeSquareSum(c int) bool {

	for a := 0; a*a <= c; a++ {
		b := c - a*a
		if binarySearch(0, b, b) {
			return true
		}
	}

	return false
}

func JudgeSquareSumV2(c int) bool {

	l, h := 0, int(math.Sqrt(float64(c)))
	for s := l*l + h*h; l < h && s != c; s = l*l + h*h {
		if s > c {
			h--
		}
		l++
	}
	return c == l*l+h*h
}

// binarySearch 二分查找
// s 最小
// l 最大
// n 求值
func binarySearch(s, l, n int) bool {
	if s > l {
		return false
	}
	mid := s + (l-s)/2
	midSquare := mid * mid
	if midSquare > n {
		return binarySearch(s, mid-1, n)
	} else if midSquare < n {
		return binarySearch(mid+1, l, n)
	}

	return true
}
