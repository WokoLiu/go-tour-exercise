// https://tour.golang.org/moretypes/18
// https://tour.go-zh.org/moretypes/18

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := [][]uint8{}
	for x := 0; x < dy; x++ {
		temp := []uint8{}
		for y := 0; y < dx; y++ {
			temp = append(temp, uint8(x^y))
			//temp = append(temp, uint8((x+y)/2))
			//temp = append(temp, uint8(x*y))
		}
		res = append(res, temp)
	}
	return res
}

func main() {
	pic.Show(Pic)
}
