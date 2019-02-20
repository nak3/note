package main

// import (
// 	"fmt"
// )

// is_prime O(sqrt(n))
func is_prime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// divisor O(sqrt(n))
func divisor(n int) []int {
	res := []int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
			if i != n/i {
				res = append(res, n/i)
			}
		}
	}
	return res

}

// prime_factor O(sqrt(n))
func prime_factor(n int) map[int]int {
	mp := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			mp[i]++
			n /= i
		}
	}
	if n != 1 {
		mp[n] = 1
	}
	return mp
}

// sieve ...
func sieve(n int) ([]bool, []int) {
	prime := []int{}
	is_prime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		is_prime[i] = true
	}

	is_prime[0], is_prime[1] = false, false

	for i := 2; i <= n; i++ {
		if is_prime[i] {
			prime = append(prime, i)
			for j := 2 * i; j <= n; j += i {
				is_prime[j] = false
			}

		}
	}
	return is_prime, prime

}

// func main() {
// 	pr := is_prime(10)
// 	fmt.Printf("%+v\n", pr) // output for debug

// 	pf := prime_factor(10)
// 	fmt.Printf("%+v\n", pf) // output for debug

// 	dv := divisor(10)
// 	fmt.Printf("%+v\n", dv) // output for debug

// 	is_prime, prime := sieve(10)
// 	fmt.Printf("%+v\n", is_prime) // output for debug
// 	fmt.Printf("%+v\n", prime)    // output for debug

// }
