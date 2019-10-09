package leetcode

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	set := make(map[int]int)
	for i, v := range nums {
		set[v] = i
	}
	for i, v := range nums {
		x := target - v
		val, ok := set[x]
		if ok && val != i {
			return []int{i, val}
		}
	}
	return []int{}
}
