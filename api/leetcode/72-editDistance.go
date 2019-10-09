package leetcode

func minDistance(word1 string, word2 string) int {

	n, m := len(word1), len(word2)

	dp := make([][]int, n+1, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1, m+1)
	}

	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			// eq
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}

			// not eq
			// insert/delete/replace
			dp[i][j] = minN(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
		}
	}

	return dp[n][m]
}

func minN(values ...int) int {

	res := values[0]

	for _, v := range values {
		if v < res {
			res = v
		}
	}

	return res
}
