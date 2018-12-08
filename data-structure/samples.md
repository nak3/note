# basic data structure

Basic data strctures in Go. Keep it simple as much as possible.

## stack

~~~go
package main

import "fmt"

func main() {
	var stack []int

	stack = append(stack, 1)
	stack = append(stack, 2)

	n := len(stack) - 1
	fmt.Println(stack[n]) // top()

	stack = stack[:n] // pop()

	n = len(stack) - 1
	fmt.Println(stack[n]) // top()
}
~~~

## queue

~~~go
package main

import "fmt"

func main() {
	queue := make([]int, 0)

	queue = append(queue, 1) // push()
	queue = append(queue, 2) // push()

	n := queue[0] // front()

	fmt.Println(n)

	queue = queue[1:] // pop()

	n = queue[0] // front()
	fmt.Println(n)
}
~~~

## priority queue
## singly linkedlist 
## doubly linkedlist 
## tree
## graph (djacency matrix)
## graph (incidence matrix)
