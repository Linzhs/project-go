package leetcode

// back-track 左右括号使用配比
func generateParenthesis(n int) []string {

	var result []string
	generator(&result, "", 0, 0, n)

	return result
}

func generator(ans *[]string, result string, left, right, n int) {
	// 左右括号都已使用完
	if left == n && right == n {
		*ans = append(*ans, result)
		return
	}

	if left < n {
		generator(ans, result+"(", left+1, right, n)
	}
	if left > right && right != n {
		generator(ans, result+")", left, right+1, n)
	}
}
