// https://tour.golang.org/methods/18
// https://tour.go-zh.org/methods/18

package main

import "fmt"
import "strings"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (a IPAddr) String() string {
	j := []string{}
	for _, i := range a {
		j = append(j, fmt.Sprintf("%d", i))
	}
	return strings.Join(j, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
