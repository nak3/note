
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
