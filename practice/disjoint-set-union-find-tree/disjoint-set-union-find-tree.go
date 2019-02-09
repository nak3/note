package disjoint_set_union_find_tree

import (
//	"fmt"
// "strings"
)

type UnionFind struct {
	p []int
}

func (u *UnionFind) init(size int) {
	u.p = make([]int, size)
	for i := 0; i < size; i++ {
		u.p[i] = i
	}
}

func (u *UnionFind) root(x int) int {
	if u.p[x] == x {
		return x
	}
	u.p[x] = u.root(u.p[x])
	return u.p[x]
}

func (u *UnionFind) unite(x, y int) {
	rx := u.root(x)
	ry := u.root(y)
	if rx != ry {
		u.p[rx] = ry
	}
}

func (u *UnionFind) same(x, y int) bool {
	rx := u.root(x)
	ry := u.root(y)
	return rx == ry
}

func solve(n, q int, data [][]int) []int {
	ans := []int{}
	u := UnionFind{}
	u.init(n)
	for i := 0; i < q; i++ {
		com := data[i][0]
		x := data[i][1]
		y := data[i][2]
		if com == 0 {
			u.unite(x, y)
		} else {
			b := u.same(x, y)
			if b == false {
				ans = append(ans, 0)
			} else {
				ans = append(ans, 1)
			}
		}
	}
	return ans
}
