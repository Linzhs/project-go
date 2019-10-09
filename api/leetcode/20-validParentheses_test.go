package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		str  string
		want bool
	}{
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
	}

	for _, test := range tests {
		got := isValid(test.str)
		if got != test.want {
			t.Fatalf("got %v but want %v", got, test.want)
		}
	}
}
