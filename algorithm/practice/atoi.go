package practice

import (
//	"../../lib"
)

func atoi(str string) int {
	r := []rune(str)
	ten := 1
	for i := 0; i < len(str)-1; i++ {
		ten *= 10
	}
	ans := 0
	for i := 0; i < len(str); i++ {
		tmp := int(r[i] - '0')
		if tmp < 0 || tmp > 9 {
			return -1
		}
		ans += tmp * ten
		ten /= 10
	}
	return ans
}
