package leetcode

func plusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		result := digits[i]
		if result%10 != 0 {
			return digits
		} else {
			digits[i] = result % 10
		}
	}

	// 999...情况
	digits = make([]int, len(digits)+1)
	digits[0] = 1
	return digits
}
