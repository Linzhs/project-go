package leetcode

import "testing"

func TestSimplifyPath(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"/home/", "/home"},
		{"home/", "/home"},
		{"/../", "/"},
		{"/home//foo/", "/home/foo"},
		{"/a/./b/../../c/", "/c"},
		{"/a/../../b/../c//.//", "/c"},
		{"/a//b////c/d//././/..", "/a/b/c"},
	}

	for _, test := range tests {
		got := simplifyPath(test.path)
		if got != test.want {
			t.Fatalf("test input:%s, want %s but got %s", test.path, test.want, got)
		}
	}
}
