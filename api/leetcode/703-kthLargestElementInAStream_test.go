package leetcode

import "testing"

func TestKthLargest_Add(t *testing.T) {

	tests := []struct {
		element int
		want    int
	}{
		{3, 4},
		{5, 5},
		{10, 5},
		{9, 8},
		{4, 8},
	}

	p := Constructor(3, []int{4, 5, 8, 2})
	for i, v := range tests {
		got := p.Add(v.element)
		if got != v.want {
			t.Fatalf("test case %d:%+v want %d but got %d", i, v, v.want, got)
		}
	}
}
