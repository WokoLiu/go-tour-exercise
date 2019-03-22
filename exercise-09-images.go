// https://tour.golang.org/methods/25
// https://tour.go-zh.org/methods/25

package main

import (
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct {
	Width, Height int
	PixelData     [][]uint8
}

type MyColorModel struct {}

func (m MyColorModel) Convert(c color.Color) color.Color {
	return c
}

func (img *Image) ColorModel() color.Model {
	return MyColorModel{}
}

func (img *Image) Bounds() image.Rectangle {
	return image.Rectangle{image.Point{0, 0}, image.Point{img.Width, img.Height}}
}

func (img *Image) At(x, y int) color.Color {
	AtColor := img.PixelData[x][y]
	//b := int(AtColor) << 8 / (256 * 256)
	//g := (int(AtColor) << 8 - b*256*256) / 256;
	//r := int(AtColor) << 8 - b*256*256 - g*256;
	// TODO transfer uint8 to beautiful RGB
	r := AtColor
	g := AtColor
	b := AtColor
	// A in RGBA is alpha (透明度)
	return color.Color(color.RGBA{uint8(r), uint8(g), uint8(b), 255})
}

// Pic() from exercise-slices
func SlicesPic(dx, dy int) [][]uint8 {
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
	PicData := SlicesPic(256, 256)
	m := Image{100, 100, PicData}
	pic.ShowImage(&m)
}
