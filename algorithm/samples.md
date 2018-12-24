basic algorithm
---

Basic algorithm in Go. Keep it simple as much as possible.

## sort

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 1, 2, 4}
	sort.Ints(arr)
	fmt.Println(arr)
}
```

## sort (struct)

```go
package main

import (
	"fmt"
	"sort"
)

type Node struct {
	val  int
	name string
}

type NodeList []Node

func (n NodeList) Len() int {
	return len(n)
}

func (n NodeList) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n NodeList) Less(i, j int) bool {
	return n[i].val < n[j].val
}

func main() {
	var arr NodeList = []Node{
		Node{1, "a"},
		Node{3, "d"},
		Node{2, "c"},
	}
	sort.Sort(arr)
	fmt.Println(arr)
}
```

## reverse

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 1, 2, 4}
	sort.Ints(arr)
	fmt.Println(arr)

	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}
```

## sort map

This is different from C++. [blog.golang.org](https://blog.golang.org/go-maps-in-action#TOC_7.) is the reference.

```
package main

import (
	"fmt"
	"sort"
)

func main() {
	mp := map[int]string{1: "a", 3: "c", 2: "b"}
	var keys []int
	for k, _ := range mp {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%+v ", mp[k]) // output for debug
	}
}
```

## sort with codition

## dynamic programming

```go
package main

import "fmt"

func main() {
	n := 9
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	fmt.Printf("%+v\n", dp[n])
}
```

## DFS, BFS, level order traveral (Tree)

```go
package main

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func New(n, k int) *Tree {
	var t *Tree
	for _, v := range rand.Perm(n) {
		t = insert(t, (1+v)*k)
	}
	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{v, nil, nil}
	}
	if v < t.Val {
		t.Left = insert(t.Left, v)
		return t
	}
	t.Right = insert(t.Right, v)
	return t
}

func dfs(root *Tree) {
	if root == nil {
		return
	}
	fmt.Printf("%+v ", root.Val) // output for debug
	dfs(root.Left)
	dfs(root.Right)
}

func bfs(root *Tree) {
	if root == nil {
		return
	}
	q := []*Tree{root}
	for {
		if len(q) == 0 {
			break
		}
		tmp := q[0]
		fmt.Printf("%v ", tmp.Val)
		if tmp.Left != nil {
			q = append(q, tmp.Left)
		}
		if tmp.Right != nil {
			q = append(q, tmp.Right)
		}
		q = q[1:]
	}
}

func levelOrder(root *Tree) {
	if root == nil {
		return
	}
	q := []*Tree{root}
	for {
		if len(q) == 0 {
			break
		}
		size := len(q)
		for i := 0; i < size; i++ {
			tmp := q[0]
			fmt.Printf("%v ", tmp.Val)
			if tmp.Left != nil {
				q = append(q, tmp.Left)
			}
			if tmp.Right != nil {
				q = append(q, tmp.Right)
			}
			q = q[1:]
		}
		fmt.Printf("\n")
	}
}

func main() {
	fmt.Printf("hello, world\n")
	t1 := New(10, 1)
	fmt.Printf("\nDFS:")
	dfs(t1)
	fmt.Printf("\nBFS:")
	bfs(t1)
	fmt.Printf("\nLevel:\n")
	levelOrder(t1)
}
```
