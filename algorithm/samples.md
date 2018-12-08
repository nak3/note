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
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
```

## sort (struct)
## reverse
## sort with codition
