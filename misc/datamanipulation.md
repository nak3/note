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

## Read stdin line by line

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if scanner.Err() != nil {
		// TODO
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

## Check string if valid IP format or not

```go
package main

import (
	"fmt"
	"regexp"

	"net"
)

func main() {
	ip := "10.1.3.2"

	validIP := regexp.MustCompile(`^[1-9][0-9]{1,3}\.[0-9]{1,4}\.[0-9]{1,4}\.[0-9]{1,4}`) // \b and \d are also available.
	fmt.Printf("%+v\n", validIP.MatchString(ip))

	// This is also possible, but To4() does not care some IP like 0.1.2.3
	ipaddr := net.ParseIP(ip)
	if ipaddr.To4() == nil {
		fmt.Printf("false") // output for debug
	} else {
		fmt.Printf("true") // output for debuga
	}

}
```
