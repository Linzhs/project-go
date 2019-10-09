package leetcode

func hammingWeight(num uint32) int {

	var cnt int

	for num != 0 {
		cnt++
		num = num & (num - 1)
	}

	return cnt
}
