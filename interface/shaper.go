package main

import "fmt"

type shaper interface {
	Area() float64
}

type Suqare struct {
	Side float64
}

func (s *Suqare) Area() float64 {
	return s.Side * s.Side
}

type Rectangle struct {
	Shape
	length, width float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

type Shape struct {
}

func (s *Shape) Area() float64 {
	return 1.0
}

func main() {
	sh := &Suqare{5}
	rh := Rectangle{Shape: Shape{}, 5, 3}

	shapes := []shaper{sh, rh}
	for n, _ := range shapes {
		fmt.Println("shape details: ", shapes[n])
		fmt.Println("Area: ", shapes[n].Area())
	}

}
