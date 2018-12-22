language processing
---

### 01. reverse string

```go
reverse string

package main

import (
	"fmt"
)

func main() {
	ss := "stressed"
	s := []rune(ss)
	l := 0
	r := len(s) - 1
	for {
		if l >= r {
			break
		}
		s[l], s[r] = s[r], s[l]
		l += 1
		r -= 1
	}
	fmt.Printf("%+v\n", string(s)) // output for debug
}
```

### 03. length of each strings

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	s := strings.Split(ss, " ")
	num := []int{}
	for i := 0; i < len(s); i++ {
		num = append(num, len(s[i]))
	}
	fmt.Printf("%+v\n", num) // output for debug
}
```


### 04. hashmap

```
package main

import (
	"fmt"
	"strings"
)

func main() {
	ss := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	s := strings.Split(ss, " ")
	mp := map[string]string{}
	for i := 0; i < len(s); i++ {
		_, ok := mp[string(s[i][:1])]
		if ok {
			mp[string(s[i])[:2]] = s[i]
		} else {
			mp[string(s[i][:1])] = s[i]
		}
	}
	fmt.Printf("%+v\n", mp) // output for debug
}
```

### 05. n-gram

```go
package main

import (
	"fmt"
	"strings"
)

func bi_gram_char(arg string) []string {
	s := []rune(arg)
	ans := []string{}
	i := 1
	for ; i < len(s); i += 2 {
		tmp := string(s[i-1]) + string(s[i])
		ans = append(ans, tmp)
	}

	if i <= len(s) {
		ans = append(ans, string(s[i-1])+string(""))
	}
	return ans
}

func bi_gram_word(arg string) [][]string {
	s := strings.Split(arg, " ")
	ans := [][]string{}
	i := 1
	for ; i < len(s); i += 2 {
		tmp := []string{s[i-1], s[i]}
		ans = append(ans, tmp)
	}
	if i <= len(s) {
		ans = append(ans, []string{string(s[i-1]), string("")})
	}
	return ans
}

func main() {
	ss := "I am an NLPer"
	fmt.Printf("%+v\n", bi_gram_word(ss)) // output for debug
	fmt.Printf("%+v\n", bi_gram_char(ss)) // output for debug}
}
```
