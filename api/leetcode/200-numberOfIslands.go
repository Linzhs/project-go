package leetcode

type unionSet struct {
	roots []int
}

func newUnionSet(grid [][]byte) *unionSet {
	u := &unionSet{}
	for i, row := range grid {
		l := len(grid[i])
		for j, v := range row {
			u.roots = append(u.roots, -1)
			if v == 1 {
				u.roots[i*l+j] = i*l + j
			}
		}
	}
	return u
}

func (u *unionSet) findRoot(i int) int {
	root := i
	for root != u.roots[root] {
		root = u.roots[root]
	}
	return root
}

func (u *unionSet) connected(p, q int) bool {
	return u.findRoot(p) == u.findRoot(q)
}

func (u *unionSet) union(p, q int) {
	xRoot := u.findRoot(p)
	yRoot := u.findRoot(q)
	u.roots[xRoot] = yRoot
}

func numIslands(grid [][]byte) int {

	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	//u := newUnionSet(grid)
	//for i, row := range grid {
	//	l := len(row)
	//	for j, v := range row {
	//		if v == 0 {
	//			continue
	//		}
	//		nl, nu := i * l + j -1, (i - 1) * l + j
	//		if nl <= 0 && nl <
	//	}
	//}

	return 0
}
