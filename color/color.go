package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	c  color.Color
	b  Rectangle
	cm color.RGBAModel
}

func (i *Image) ColorModel() color.RGBAModel {
	return cm
}

func (i *Image) Bounds() image.Rectangle {
	return b
}

func (i *Image) At(x, y int) color.Color {
	return c
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}
