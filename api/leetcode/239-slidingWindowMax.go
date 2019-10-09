package leetcode

type deque []int

func (d *deque) push(i int) {
	*d = append(*d, i)
}

func (d *deque) lPop() {
	if len(*d) > 0 {
		*d = (*d)[1:]
	}
}

func (d *deque) rPop() {
	if len(*d) > 0 {
		*d = (*d)[:len(*d)-1]
	}
}

// Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
// Output: [3,3,5,5,6,7]
func maxSlidingWindow(nums []int, k int) []int {

	if len(nums) == 0 {
		return []int{}
	}

	window := make(deque, 0, k)         // index
	maxVal := make([]int, 0, len(nums)) // result
	for i, v := range nums {
		if i >= k && window[0] <= i-k { // 超过窗口大小的元素移除
			window.lPop()
		}
		// 将当前元素大过滑动窗口左边的元素全部移除 使得最左边元素处于最大
		for j := len(window) - 1; j >= 0; j-- {
			if nums[window[j]] <= v {
				window.rPop()
			}
		}

		window.push(i)
		if i >= k-1 {
			maxVal = append(maxVal, nums[window[0]])
		}
	}

	return maxVal
}
