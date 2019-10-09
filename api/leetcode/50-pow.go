package leetcode

func myPowRecursive(x float64, n int) float64 {
	if n == 0 {
		return 1
	} else if n < 0 { // 翻转
		n = -n
		x = 1 / x
	}

	switch n%2 == 0 {
	case true:
		return myPowRecursive(x*x, n/2)
	default:
		return myPowRecursive(x*x, n/2) * x
	}
}

func myPow(x float64, n int) float64 {

	if n == 0 {
		return 1
	} else if n < 0 {
		n = -n
		x = 1 / x
	}

	var result float64 = 1
	for n != 0 {
		if (n & 1) == 1 {
			result *= x
		}
		x *= x
		n >>= 1
	}
	return result
}
