package find_the_closest_prime_number_to_a_given_number

//package main

import "fmt"

func prime() []int {
	prime := make([]bool, 1000001)
	for i := 0; i < 1000001; i++ {
		prime[i] = true
	}
	prime[0] = false
	prime[1] = false

	for i := 2; i < 1000001; i++ {
		for j := i * 2; j < 1000001; j += i {
			prime[j] = false
		}
	}
	primes := []int{}
	for i := 0; i < 1000001; i++ {
		if prime[i] {
			primes = append(primes, i)

		}
	}

	return primes
}

func binarySearch(p []int, target int) int {
	l := 0
	r := len(p) - 1
	for l <= r {
		mid := (l + r) / 2
		if p[mid] == target {
			return target
		}
		if mid == 0 || mid == len(p)-1 {
			return p[mid]
		}
		if p[mid] < target && p[mid+1] > target {
			return p[mid]
		}
		if target < p[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return 0
}

func solve(target int) int {
	p := prime()

	return binarySearch(p, target)
}

func main() {
	fmt.Printf("%+v\n", solve(14)) // output for debug
}
