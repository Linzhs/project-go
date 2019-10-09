package leetcode

func lengthOfLIS(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	var res int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}

	return res
}

func lengthOfLISV2(nums []int) int {

	//s := make([]int, 0)
	//for i, v := range nums {
	//
	//}

	return 0
}
