package shtrutura

import "fmt"

type Rectangle struct {
	Width  float64
	Height float64
}

var (
	r = Rectangle{
		Width:  0,
		Height: 0,
	}
)

func init() {
	fmt.Println(r.Width * r.Height)
}
