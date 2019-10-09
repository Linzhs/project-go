package leetcode

func fibV1(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibV1(n-1) + fibV1(n-2)
}

func fibV2(n int, memo []int) int {
	if n == 0 || n == 1 {
		return n
	}
	if memo[n] == 0 {
		memo[n] = fibV2(n-1, memo) + fibV2(n-2, memo)
	}
	return memo[n]
}

func fibV3(n int) int {
	fn := make([]int, n+1)
	fn[0], fn[1] = 0, 1
	for i := 2; i <= n; i++ {
		fn[i] = fn[i-1] + fn[i-2]
	}
	return fn[n]
}
