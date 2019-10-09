package leetcode

import (
	"fmt"
	"testing"
)

var EPSILON float64 = 0.00000001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}

func TestMyPow(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{2.0, 10, 1024.00000},
		{2.1, 3, 9.26100},
		{2.0, -21, 0},
	}

	for _, v := range tests {
		got := myPow(v.x, v.n)
		if !floatEquals(got, v.want) {
			t.Fatalf("want %+v but got %+v", v.want, got)
		}
	}
}

func TestSqrtx(t *testing.T) {
	fmt.Println(mySqrt(8))
}
