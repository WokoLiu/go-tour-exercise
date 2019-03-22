package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.

// start with 0
func fibonacci() func() int {
	first, second := 1, 0
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
