// example 6
package main

import "fmt"

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Scale(s float64) {
	r.Width *= s
	r.Height *= s
}

func main() {
	rect := Rectangle{Width: 3, Height: 4}
	fmt.Println("Luas sebelum scaling:", rect.Area()) // 12

	rect.Scale(2)
	fmt.Println("Luas setelah scaling:", rect.Area()) // 48
}
