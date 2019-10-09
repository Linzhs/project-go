package leetcode

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	var result int
	for l, r := 1, x; l <= r; {
		m := (l + r) / 2
		if m < x/m {
			l = m + 1
			result = m
		} else if m > x/m {
			r = m - 1
		} else {
			return m
		}
	}

	return result
}
