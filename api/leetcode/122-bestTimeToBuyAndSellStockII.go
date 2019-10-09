package leetcode

// Greedy Algorithm
func maxProfit2(prices []int) int {

	var sum int
	for i, v := range prices {
		if i+1 < len(prices) && v < prices[i+1] {
			sum += prices[i+1] - v
		}
	}

	return sum
}
