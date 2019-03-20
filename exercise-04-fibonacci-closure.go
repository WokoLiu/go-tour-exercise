package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

// start with 0
func fibonacci() func() int {
	first, second := -1, 1
	fib := func() int {
		first, second = second, first+second
		return second
	}
	return fib
}

// start with 1
func fibonacci2() func() int {
	first, second := 0, 1
	fib := func() int {
		first, second = second, first+second
		return first
	}
	return fib
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
