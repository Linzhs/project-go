package leetcode

func largestRectangleArea(heights []int) int {

	if len(heights) == 0 {
		return 0
	}
	if len(heights) == 1 {
		return heights[0]
	}

	stack := make([]int, 0, len(heights))
	stack = append(stack, -1)
	heights = append(heights, -1)

	var res int
	for i, v := range heights {
		top := stack[len(stack)-1]
		for top != -1 && heights[top] >= v {
			stack = stack[:len(stack)-1]
			top2 := stack[len(stack)-1]
			res = max(res, (i-top2-1)*heights[top])
			top = top2
		}

		stack = append(stack, i)
	}

	return res
}
