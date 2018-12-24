language processing
---

problems: http://www.cl.ecei.tohoku.ac.jp/nlp100/

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
		s[i] = strings.Trim(s[i], ".")
		s[i] = strings.Trim(s[i], ",")
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

### 07. strings from temprate

```go
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	var x, y, z string
	if len(args) != 3 {
		os.Exit(1)
	}
	x = args[0]
	y = args[1]
	z = args[2]

	fmt.Printf("%s o'clock's %s is %s\n", x, y, z) // output for debug

}
```

### 08. cipher

```go
package main

import (
	"fmt"
)

func main() {
	ss := "FOO2foo!"
	r := []rune(ss)
	for i := 0; i < len(r); i++ {
		if r[i] >= 'a' && r[i] <= 'z' {
			r[i] = rune(219 - int(r[i]))
		}
	}
	fmt.Printf("%s\n", string(r)) // output for debug
}
```

### 09. Typoglycemia

```go
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func shuffle(s string) string {
	r := []rune(s)
	idx := len(r) - 1
	for i := len(r) - 1; i >= 1; i-- {
		tmp := rand.Intn(idx)
		r[i], r[tmp] = r[tmp], r[i]
		idx--
	}
	return string(r)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ss := "I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind."
	sslice := strings.Split(ss, " ")
	for i := 0; i < len(sslice); i++ {
		if len(sslice[i]) > 4 {
			sslice[i] = shuffle(sslice[i])
		}
	}
	ans := sslice[0]
	for i := 1; i < len(sslice); i++ {
		ans += " " + sslice[i]
	}
	fmt.Printf("%s\n", ans) // output for debug
}
```

### 10. Count number of lines

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("hightemp.txt")
	if err != nil {
		os.Exit(2)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	cnt := 0
	for scanner.Scan() {
		cnt += 1
	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", cnt) // output for debug
}
```

### 11. Replace tab with space

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("hightemp.txt")
	if err != nil {
		os.Exit(2)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		str := scanner.Text()
		r := []rune(str)

		for i := 0; i < len(r); i++ {
			if r[i] == rune('\t') {
				r[i] = rune(' ')
			}
		}
		fmt.Printf("%+v\n", string(r)) // output for debug

	}
	if err = scanner.Err(); err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
}
```

### 12. Write 1st and 2nd colomn into each file

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("hightemp.txt")
	if err != nil {
		// TODO
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sl [][]string
	for scanner.Scan() {
		str := scanner.Text()
		sl = append(sl, strings.Split(str, "\t"))
	}
	for i := 0; i < len(sl); i++ {

	}

	fc1, err := os.Create(`/tmp/col1.txt`)
	if err != nil {
		// TODO
	}
	defer fc1.Close()
	fc2, err := os.Create(`/tmp/col2.txt`)
	if err != nil {
		// TODO
	}
	defer fc2.Close()

	for i := 0; i < len(sl); i++ {
		//		fmt.Printf("%+v\n", sl[i][0]) // output for debug
		fc1.Write(([]byte)(sl[i][0]))
		fc1.Write(([]byte)("\n"))
		fc2.Write(([]byte)(sl[i][1]))
		fc2.Write(([]byte)("\n"))
	}
	fmt.Printf("done...\n") // output for debug
}
```

### 13. merge files created in 12

### 14. First N lines

```go
package main

import (
	"fmt"
	"bufio"
	"flag"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		//
	}
	n, _ := strconv.Atoi(args[0])

	f, err := os.Open("hightemp.txt")
	if err != nil {
		//
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	idx := 0
	for scanner.Scan() {
		if idx == n {
			break
		}

		str := scanner.Text()
		fmt.Printf("%+v\n", str) // output for debug
		idx += 1
	}
}
```

### 15. Last N lines

NOTE: This code does not work if big file data due to out of memory.

```go
package main

import (
	"fmt"
	"bufio"
	"flag"
	"os"
	"strconv"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		//
	}
	n, _ := strconv.Atoi(args[0])

	f, err := os.Open("hightemp.txt")
	if err != nil {
		//
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i := len(lines) - n; i < len(lines); i++ {
		fmt.Printf("%+v\n", lines[i]) // output for debug
	}
}
```

### (Optional) 10-19 by Linux commands

```sh
#!/bin/bash

# wget http://www.cl.ecei.tohoku.ac.jp/nlp100/data/hightemp.txt

# 10.
cat hightemp.txt | wc -l

# 11.
cat hightemp.txt | sed -e 's/\t/ /g'

# 12.

cat hightemp.txt | awk '{print $1}' > /tmp/col1.txt
cat hightemp.txt | awk '{print $2}' > /tmp/col2.txt

# 13.
paste /tmp/col1.txt /tmp/col2.txt

# 14.
head -n 3 /tmp/col2.txt

# 15.
tail -n 3 /tmp/col2.txt

# 16.

# 17.
cat hightemp.txt | awk '{print $1}' | sort | uniq | wc -l

# 18.
cat hightemp.txt | sort -r -n -k 3,3

# 19.
cat hightemp.txt | awk '{print $1}' | sort | uniq -c | sort -nr
```
