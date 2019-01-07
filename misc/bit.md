
## List factorial of 2 by N

```go
package main

import "fmt"

func main() {
	v := 100000
	var i, data, a uint64
	a = 1
	for i = 0; data <= uint64(v); i++ {
		data = a << i
		fmt.Printf("%+v\n", data) // output for debug
	}
}
```

## List all subsets

```go
package main

import (
	"fmt"
)

func main() {
	subsets([]int{1, 2, 3})
}

func subsets(nums []int) [][]int {
	for i := 0; i < (1 << uint(len(nums))); i++ {
		fmt.Printf("[") // output for debug

		for j := uint(0); j < uint(len(nums)); j++ {
			if (i & (1 << j)) > 0 {
				fmt.Printf("%v", nums[j])
			}
		}
		fmt.Printf("]\n") // output for debug
	}
	return [][]int{}
}
```

## Hex from/to String

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("%02x\n", 10)

	x, _ := strconv.ParseInt("a", 16, 64)
	fmt.Printf("%+v\n", x)
}
```
