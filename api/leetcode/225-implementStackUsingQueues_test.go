package leetcode

import (
	"fmt"
	"testing"
)

func TestMyStack(t *testing.T) {
	stack := MyStack{}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Pop()
	stack.Pop()
	fmt.Println(stack.Empty())
}
