package leetcode

type MyStack struct {
	FirstQueue  TheQueue
	SecondQueue TheQueue
}

/** Initialize your data structure here. */
func ConstructorMyStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.SecondQueue.enqueue(x)
	for !this.FirstQueue.isEmpty() {
		this.SecondQueue.enqueue(this.FirstQueue.dequeue())
	}
	this.FirstQueue = this.SecondQueue
	this.SecondQueue = TheQueue{}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.FirstQueue.isEmpty() {
		return 0
	}

	x := this.FirstQueue.top()
	this.FirstQueue.dequeue()
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.FirstQueue.top()
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.FirstQueue.isEmpty()
}

type TheQueue struct {
	Queue []int
}

func newQueue() *TheQueue {
	return &TheQueue{Queue: []int{}}
}

func (p *TheQueue) enqueue(x int) {
	p.Queue = append(p.Queue, x)
}

func (p *TheQueue) dequeue() int {
	if p.isEmpty() {
		return 0
	}
	item := p.Queue[0]
	p.Queue = p.Queue[1:len(p.Queue)]
	return item
}

func (p *TheQueue) top() int {
	if len(p.Queue) == 0 {
		return 0
	}
	return p.Queue[0]
}

func (p *TheQueue) isEmpty() bool {
	return len(p.Queue) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
