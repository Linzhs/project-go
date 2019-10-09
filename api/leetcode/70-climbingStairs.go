package leetcode

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	result := 0
	oneStepBefore, twoStepBefore := 1, 1
	for i := 2; i <= n; i++ {
		result = oneStepBefore + twoStepBefore
		oneStepBefore, twoStepBefore = result, oneStepBefore
	}

	return result
}
