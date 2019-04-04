package count_the_reversals

// package REPLACE

import (
	"fmt"
	//	"github.com/nak3/note/lib"
)

// solve ...
func solve(s string) int {
	stk := []string{}
	cnt := 0

	for i := 0; i < len(s); i++ {
		if string(s[i]) == "{" {
			stk = append(stk, "{")
		} else {
			if len(stk) == 0 {
				stk = append(stk, "}")
			} else if string(stk[len(stk)-1]) == "{" {
				stk = stk[0 : len(stk)-1]
			} else {
				stk = append(stk, "}")
			}
		}
	}

	tmp := []string{}
	for i := 0; i < len(stk); i++ {
		if string(stk[i]) == "}" {
			if len(tmp) == 0 {
				tmp = append(tmp, "{")
				cnt++
			} else {
				if tmp[0] == "{" {
					tmp = tmp[1:]
				} else {
					tmp = tmp[1:]
					cnt++
				}
			}
		} else {
			if len(tmp) != 0 {
				tmp = tmp[1:]
				cnt++
			} else {
				tmp = append(tmp, "{")
			}
		}
	}
	if len(tmp) != 0 {
		return -1
	}
	return cnt

}

func main() {
	fmt.Printf("%+v\n", solve("}{{}}{{{"))      // output for debug
	fmt.Printf("%+v\n", solve("{{}}}}"))        // output for debug
	fmt.Printf("%+v\n", solve("{{}{{{}{{}}{{")) // output for debug
	fmt.Printf("%+v\n", solve("{{{{}}}}"))      // output for debug

}
