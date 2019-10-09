package leetcode

func isPowerOfTwo(n int) bool {

	var cnt int
	for n != 0 {
		cnt++
		n = n & (n - 1)
		if cnt > 1 {
			return false
		}
	}
	return cnt == 1
}
