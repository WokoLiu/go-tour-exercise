// https://tour.golang.org/concurrency/7
// https://tour.go-zh.org/concurrency/7

package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

func TestWalk(f func(t *tree.Tree, ch chan int)) bool {
	ok := true
	ch := make(chan int)
	for k := 1; k < 10; k++ {
		t := tree.New(k)
		go f(t, ch)
		for i := 1; i <= 10; i++ {
			if i*k != <-ch {
				ok = false
				break
			}
		}
	}
	return ok
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

func TestSame(f func(t1, t2 *tree.Tree) bool) bool {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			same := f(tree.New(i), tree.New(j))
			if i == j && same != true {
				return false
			}
			if i != j && same == true {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(TestWalk(Walk))
	fmt.Println(TestSame(Same))
}
