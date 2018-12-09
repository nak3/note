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


## sort with codition
