basic multithread related samples
---

Basic multi threading in Go. Keep it simple as much as possible.


## thread unsafe

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	cnt := 0
	for i := 0; i < 10; i++ {
		go func() {
			cnt += 1
			fmt.Printf("%+v ", cnt)
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("%v", cnt)
}
```

## mutex simple example

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cnt := 0
	for i := 0; i < 10; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			cnt += 1
			fmt.Printf("%+v ", cnt)
		}()
	}
	time.Sleep(3 * time.Second)
	fmt.Printf("%v", cnt)
}
```

## WaitGroup example

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	wg := sync.WaitGroup{}
	cnt := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mu.Lock()
			defer mu.Unlock()
			cnt += 1
			fmt.Printf("%+v ", cnt)
			wg.Done()
		}()
	}
	wg.Wait()
}
```
