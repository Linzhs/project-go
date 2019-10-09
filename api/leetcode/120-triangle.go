package leetcode

// example:
// 		2
// 	   3 3
//	  6 5 7
// 	 4 1 8 3

// dp[i] below->up方向 当前节点最优路径值
func minimumTotalV2(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}

	dp := make([]int, len(triangle))
	for i := 0; i < len(triangle[len(triangle)-1]); i++ {
		dp[i] = triangle[len(triangle)-1][i]
	}
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = triangle[i][j] + min(dp[j], dp[j+1])
		}
	}

	return dp[0]
}

func minimumTotalV1(triangle [][]int) int {

	if len(triangle) == 0 {
		return 0
	}

	memo := make([][]int, len(triangle))
	for i := 0; i < len(triangle); i++ {
		memo[i] = make([]int, len(triangle[i]))
	}
	memo[0][0] = triangle[0][0]

	return minimumTotalRecursion(triangle, &memo, 0, 0)
}

func minimumTotalRecursion(triangle [][]int, memo *[][]int, row, col int) int {
	if row+1 >= len(triangle) || col+1 >= len(triangle[row+1]) {
		return triangle[row][col]
	}

	var left, right int

	leftValue := (*memo)[row+1][col]
	if leftValue != 0 {
		left = triangle[row][col] + leftValue
	} else {
		left = triangle[row][col] + minimumTotalRecursion(triangle, memo, row+1, col)
	}

	rightValue := (*memo)[row+1][col+1]
	if rightValue != 0 {
		right = triangle[row][col] + rightValue
	} else {
		right = triangle[row][col] + minimumTotalRecursion(triangle, memo, row+1, col+1)
	}

	if left > right {
		(*memo)[row][col] = right
		return right
	} else {
		(*memo)[row][col] = left
	}

	return left
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
