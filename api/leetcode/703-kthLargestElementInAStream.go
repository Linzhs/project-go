package leetcode

import "container/heap"

type MinIntHeap []int

type KthLargest struct {
	kth         int
	minIntHeaps MinIntHeap
}

func Constructor(k int, nums []int) KthLargest {
	var p KthLargest
	switch len(nums) <= k {
	case true:
		p = KthLargest{kth: k, minIntHeaps: nums}
		heap.Init(&p.minIntHeaps)
	default:
		p = KthLargest{kth: k, minIntHeaps: nums[:k]}
		heap.Init(&p.minIntHeaps)
		for i := k; i < len(nums); i++ {
			p.Add(nums[i])
		}
	}

	return p
}

func (this *KthLargest) Add(val int) int {
	switch this.minIntHeaps.Len() < this.kth {
	case true:
		heap.Push(&this.minIntHeaps, val)
	default:
		if this.minIntHeaps[0] < val {
			this.minIntHeaps[0] = val
			heap.Fix(&this.minIntHeaps, 0)
		}
	}

	return this.minIntHeaps[0]
}

func (p MinIntHeap) Len() int {
	return len(p)
}

func (p MinIntHeap) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p MinIntHeap) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *MinIntHeap) Push(x interface{}) {
	*p = append(*p, x.(int))
}

func (p *MinIntHeap) Pop() interface{} {
	old := *p
	l := len(old)
	x := old[l-1]
	*p = old[:l-1]
	return x
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
