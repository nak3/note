package weighted_union_find_trees

import (
	"fmt"
	// "strings"
	"strconv"
)

type UnionFind struct {
	p      []int
	weight []int
	rank   []int
}

func (u *UnionFind) init(size int) {
	u.p = make([]int, size)
	for i := 0; i < size; i++ {
		u.p[i] = i
	}
	u.weight = make([]int, size)
	u.rank = make([]int, size)
}

func (u *UnionFind) root(x int) int {
	if u.p[x] == x {
		return x
	} else {
		r := u.root(u.p[x])
		u.weight[x] += u.weight[u.p[x]]
		u.p[x] = r
		return u.p[x]
	}
}

func (u *UnionFind) unite(x, y, w int) {
	rx := u.root(x)
	ry := u.root(y)

	if u.rank[rx] < u.rank[ry] {
		u.p[rx] = ry
		u.weight[rx] = w - u.weight[x] + u.weight[y]
	} else {
		u.p[ry] = rx
		u.weight[ry] = -w - u.weight[y] + u.weight[x]
		if u.rank[rx] == u.rank[ry] {
			u.rank[rx] += 1
		}
	}
}

func (u *UnionFind) diff(x, y int) int {
	return u.weight[x] - u.weight[y]
}

func (u *UnionFind) same(x, y int) bool {
	rx := u.root(x)
	ry := u.root(y)
	return rx == ry
}

// solve ...
func solve(n, q int, data [][]int) []string {
	ans := []string{}
	u := UnionFind{}
	u.init(n)
	for i := 0; i < q; i++ {
		switch data[i][0] {
		case 0:
			u.unite(data[i][1], data[i][2], data[i][3])
		case 1:
			if !u.same(data[i][1], data[i][2]) {
				ans = append(ans, "?")
			} else {
				diff := u.diff(data[i][1], data[i][2])
				ans = append(ans, strconv.Itoa(diff))
			}
		}
	}
	return ans
}

func main() {
	n := 5
	q := 6
	data := [][]int{
		{0, 0, 2, 5},
		{0, 1, 2, 3},
		{1, 0, 1, 0},
		{1, 1, 3, 0},
		{0, 1, 4, 8},
		{1, 0, 4, 0},
	}
	ans := solve(n, q, data)
	fmt.Printf("%+v\n", ans) // output for debug

}
