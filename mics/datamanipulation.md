---


## Read file line by line

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		username := strings.Split(str, ":")[0]
		fmt.Printf("%v\n", username)
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
```

## Check string if alphabet or number only or not

```go
package main

import (
	"fmt"
	"unicode"
)

func main() {
	str := "abcdef123"
	j := true
	for _, r := range str {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			j = false
		}
	}
	fmt.Printf("%+v\n", j)

	str2 := "abcdef123>1"
	for _, r := range str2 {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			j = false
		}
	}
	fmt.Printf("%+v\n", j)
}
```
