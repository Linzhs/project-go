package leetcode

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, test := range tests {
		got := isAnagram(test.s, test.t)
		if got != test.want {
			t.Fatalf("got %+v but want %+v", got, test.want)
		}
	}
}
