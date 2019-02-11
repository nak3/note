package all_pairs_shortest_path

import (
	"errors"
	"fmt"
)

const (
	INFTY = 1 << 32
)

type floyd struct {
	matrix [][]int64
}

// wf ...
func (f *floyd) wf() {
	length := len(f.matrix)
	for k := 0; k < length; k++ {
		for i := 0; i < length; i++ {
			if f.matrix[i][k] == INFTY {
				continue
			}
			for j := 0; j < length; j++ {
				if f.matrix[k][j] == INFTY {
					continue
				}
				tmp := f.matrix[i][k] + f.matrix[k][j]
				if tmp < f.matrix[i][j] {
					f.matrix[i][j] = tmp
				}
			}
		}
	}
}

// solve ...
func solve(v, e int, data [][]int) ([][]int64, error) {
	f := &floyd{make([][]int64, v)}
	for i := 0; i < v; i++ {
		f.matrix[i] = make([]int64, v)
	}

	for i := 0; i < v; i++ {
		for j := 0; j < v; j++ {
			if i == j {
				f.matrix[i][j] = 0
			} else {
				f.matrix[i][j] = INFTY
			}
		}
	}

	// initialize by input data
	for i := 0; i < e; i++ {
		s := uint(data[i][0])
		g := data[i][1]
		d := data[i][2]
		f.matrix[s][g] = int64(d)
	}

	f.wf()
	for i := 0; i < v; i++ {
		if f.matrix[i][i] < 0 {
			return nil, errors.New("NEGATIVE CYCLE\n")
		}
	}

	return f.matrix, nil
}

func main() {
	v := 4
	e := 6
	data := [][]int{
		{0, 1, 1},
		{0, 2, 5},
		{1, 2, 2},
		{1, 3, 4},
		{2, 3, 1},
		{3, 2, 7},
	}
	ans, _ := solve(v, e, data)
	fmt.Printf("%+v\n", ans)
}
