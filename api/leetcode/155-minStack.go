package leetcode

//["MinStack","push","push","push","getMin","pop","top","getMin"]
//[[],[-2],[0],[-3],[],[],[],[]]

type MinStack struct {
	Val []int
	Min []int
}

/** initialize your data structure here. */
func MinStackConstructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {

	if len(this.Val) <= 0 {
		this.Min = append(this.Min, x)
		this.Val = append(this.Val, x)
		return
	}

	this.Val = append(this.Val, x)
	top := this.Min[len(this.Min)-1]
	switch top > x {
	case true:
		this.Min = append(this.Min, x)
	default:
		this.Min = append(this.Min, top)
	}
}

func (this *MinStack) Pop() {
	if len(this.Val) <= 0 {
		return
	}
	this.Val = this.Val[:len(this.Val)-1]
	this.Min = this.Min[:len(this.Min)-1]
}

func (this *MinStack) Top() int {
	if len(this.Val) <= 0 {
		return 0
	}
	return this.Val[len(this.Val)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.Min) <= 0 {
		return 0
	}
	return this.Min[len(this.Min)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
