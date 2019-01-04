package flood_fill_algorithm

//package main

import (
	"fmt"
)

func solve(place []int, color int, data [][]int) [][]int {
	maxX := len(data[0])
	maxY := len(data)

	visited := make([][]int, len(data))
	for i := 0; i < len(data); i++ {
		visited[i] = make([]int, len(data[0]))
	}
	q := [][]int{place}
	visited[place[0]][place[1]] = 1
	base := data[place[0]][place[1]]

	for len(q) > 0 {
		tmp := q[0]
		q = q[1:]
		x := tmp[1]
		y := tmp[0]
		if x-1 >= 0 && data[y][x-1] == base && visited[y][x-1] == 0 {
			data[y][x-1] = color
			q = append(q, []int{y, x - 1})
		}
		if y-1 >= 0 && data[y-1][x] == base && visited[y-1][x] == 0 {
			data[y-1][x] = color
			q = append(q, []int{y - 1, x})
		}
		if y+1 < maxY && data[y+1][x] == base && visited[y+1][x] == 0 {
			data[y+1][x] = color
			q = append(q, []int{y + 1, x})
		}
		if x+1 < maxX && data[y][x+1] == base && visited[y][x+1] == 0 {
			data[y][x+1] = color
			q = append(q, []int{y, x + 1})
		}
		data[y][x] = color
		visited[y][x] = 1
	}

	return data
}

func main() {
	data := [][]int{
		{0, 1, 1},
		{0, 1, 1},
		{1, 1, 0},
		{1, 2, 3},
	}
	place := []int{0, 1}
	color := 5

	ans := solve(place, color, data)
	fmt.Printf("%+v\n", ans) // output for debug

}

/*
    0 1 2
0 : 0 1 1
1 : 0 1 1
2 : 1 1 0
3 : 1 2 3

1 2 3 4
1 2 3 4
1 2 3 4
1 3 2 4
*/
