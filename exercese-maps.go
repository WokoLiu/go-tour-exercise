package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := map[string]int{}
	sl := strings.Split(s, " ")

	for i := 0; i < len(sl); i++ {
		m[sl[i]] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
