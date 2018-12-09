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
