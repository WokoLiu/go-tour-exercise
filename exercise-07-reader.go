// https://tour.golang.org/methods/22
// https://tour.go-zh.org/methods/22

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// A Read([]byte) (int, error) method to MyReader.
func (r MyReader) _Read(b []byte) (n int, err error) {
	// One byte at a time is usable but too inefficient
	b[0] = 'A'
	return 1, nil
}

func (r MyReader) Read(b []byte) (n int, err error) {
	// use len(b), not cap(b)
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}
