package leetcode

// input: [2,3,-2,4]
// output: 6
func maxProduct(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	curMax, curMin, res := nums[0], nums[0], nums[0]
	for _, v := range nums[1:] {
		curMax, curMin = curMax*v, curMin*v
		curMin, curMax = threeMin(curMin, curMax, v), threeMax(v, curMax, curMin)
		if curMax > res {
			res = curMax
		}
	}

	return res
}

func threeMin(x, y, z int) int {
	s := []int{x, y, z}

	for _, v := range s {
		if x > v {
			x = v
		}
	}
	return x
}

func threeMax(x, y, z int) int {
	s := []int{x, y, z}

	for _, v := range s {
		if x < v {
			x = v
		}
	}
	return x
}
