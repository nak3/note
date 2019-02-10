package range_minimum_quer

import (
	"fmt"
	// "strings"
	"math"
)

type SegmentTree struct {
	size int
	data []int
	tree []int
}

func (s *SegmentTree) init(original []int) {
	s.data = original
	n := len(s.data)
	N := 1
	for N < n {
		N *= 2
	}
	s.tree = make([]int, N*2)
	for i := 0; i < N*2; i++ {
		s.tree[i] = math.MaxInt32
	}

	for i := 0; i < len(s.data); i++ {
		s.update(i, s.data[i])
	}
}

func (s *SegmentTree) update(i, x int) {
	N := len(s.tree) / 2
	i = N + i - 1

	s.tree[i] = x
	fmt.Printf("%+v\n", s.tree) // output for debug

	for i > 0 {
		i = (i - 1) / 2
		if s.tree[i*2+1] < s.tree[i*2+2] {
			s.tree[i] = s.tree[i*2+1]
		} else {
			s.tree[i] = s.tree[i*2+2]
		}
	}
}

func (s *SegmentTree) query(start, end int) int {
	return s.mqr(start, end, 0, 0, len(s.tree)/2)
}

func (s *SegmentTree) mqr(a, b, k, l, r int) int {
	if r <= a || b <= l {
		return math.MaxInt32
	}
	if a <= l && r <= b {
		return s.tree[k]
	}
	lv := s.mqr(a, b, 2*k+1, l, (l+r)/2)
	rb := s.mqr(a, b, 2*k+2, (l+r)/2, r)
	if lv < rb {
		return lv
	}
	return rb
}

// solve ...
func solve(n, q int, data [][]int) []int {
	ans := []int{}
	s := &SegmentTree{}
	dummy := make([]int, n)
	for i := 0; i < n; i++ {
		dummy[i] = math.MaxInt32
	}
	s.init(dummy)

	for i := 0; i < q; i++ {
		c := data[i][0]
		x := data[i][1]
		y := data[i][2]
		if c == 0 {
			s.update(x, y)
		} else {
			ans = append(ans, s.query(x, y+1))
			fmt.Printf("%+v\n", s.query(x, y+1))
		}
	}
	return ans
}

func main() {
	n := 3
	q := 5
	data := [][]int{
		{0, 0, 1},
		{0, 1, 2},
		{0, 2, 3},
		{1, 0, 2},
		{1, 1, 2},
	}
	solve(n, q, data)
}
