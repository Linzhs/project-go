package leetcode

import "testing"

func TestJudgeSquareSum(t *testing.T) {
	testTable := []struct {
		c    int
		want bool
	}{
		{0, true},
		{3, false},
		{5, true},
	}

	for _, test := range testTable {
		got := JudgeSquareSum(test.c)
		if got != test.want {
			t.Fatalf("want %v but got %v", test.want, got)
		}
	}
}

func TestJudgeSquareSumV2(t *testing.T) {
	testTable := []struct {
		c    int
		want bool
	}{
		{3, false},
		{5, true},
	}

	for _, test := range testTable {
		got := JudgeSquareSumV2(test.c)
		if got != test.want {
			t.Fatalf("want %v but got %v", test.want, got)
		}
	}
}
