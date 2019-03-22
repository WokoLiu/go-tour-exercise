// https://tour.golang.org/methods/18
// https://tour.go-zh.org/methods/18

package main

import "fmt"
import "strings"

type IPAddr [4]byte

// Add a "String() string" method to IPAddr.

func (ip IPAddr) _String() string {
	// if IPAddr has varying length
	// this maybe ok
	j := []string{}
	for _, i := range ip {
		j = append(j, fmt.Sprintf("%d", i))
	}
	return strings.Join(j, ".")
}

func (ip IPAddr) String() string {
	// since we know type IPAddr is [4]byte,
	// it's no need to use strings.Join
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
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
