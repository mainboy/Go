package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s *stockPosition) getvalue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c *car) getvalue() float32 {
	return c.price
}

type valuable interface {
	getvalue() float32
}

func showvalue(asset valuable) {
	fmt.Println("Value :", asset.getvalue())
}

func main() {
	s := stockPosition{
		"OK",
		3.2,
		5,
	}

	c := car{
		"BWM",
		"BWM",
		32,
	}

	showvalue(&s)
	showvalue(&c)
}
