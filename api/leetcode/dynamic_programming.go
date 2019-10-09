package leetcode

func countPathsV1(grid [][]bool, row, col int) int {
	if validSquare(grid, row, col) {
		return 0
	}
	if isAtEnd(grid, row, col) {
		return 1
	}
	return countPathsV1(grid, row+1, col) + countPathsV1(grid, row, col+1)
}

func validSquare(grid [][]bool, row, col int) bool {
	return grid[row][col] == true
}

func isAtEnd(grid [][]bool, row, col int) bool {
	r := len(grid)
	c := len(grid[0])
	return r == row && col == c
}

// dp: opt[i, j] = opt[i, j - 1] + opt[i - 1, j]
// 	if 空地 ==>  opt[i, j] = opt[i, j - 1] + opt[i - 1, j]
//  else 障碍 ==> opt[i, j] = 0
func countPathsV2(grid [][]bool) int {
	r := len(grid)
	c := len(grid[0])

	if r == 0 || c == 0 {
		return 0
	}

	opt := make([][]int, r)
	for i := 0; i < r; i++ {
		opt[i] = make([]int, c)
	}
	opt[r-1][c-1] = 1

	for i := r - 1; i >= 0; i-- {
		for j := c - 1; j >= 0; j-- {
			if validSquare(grid, i, j) {
				opt[i][j] = 0
			} else {
				if i+1 < r {
					opt[i][j] += opt[i+1][j]
				}
				if j+1 < c {
					opt[i][j] += opt[i][j+1]
				}
			}
		}
	}
	return opt[0][0]
}
