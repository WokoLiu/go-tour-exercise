// https://tour.golang.org/moretypes/23
// https://tour.go-zh.org/moretypes/23

package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	//m := map[string]int{}
	m := make(map[string]int)

	sl := strings.Split(s, " ")

	for i := 0; i < len(sl); i++ {
		m[sl[i]] ++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
