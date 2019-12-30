package main

import (
	"fmt"
	"image/color"
	"math"
)

type Color struct {
	R, G, B float64
}

var (
	White   = Color{1.0, 1.0, 1.0}
	Black   = Color{0.0, 0.0, 0.0}
	SkyBlue = Color{0.5, 0.7, 1.0}
)

func (c Color) Scale(f float64) Color {
	return Color{c.R * f, c.G * f, c.B * f}
}

func (c Color) Mult(c2 Color) Color {
	return Color{c.R * c2.R, c.G * c2.G, c.B * c2.B}
}

func (c Color) Add(c2 Color) Color {
	return Color{c.R + c2.R, c.G + c2.G, c.B + c2.B}
}

func (c Color) asRGBA() color.RGBA {
	r := uint8(math.Min(255.0, c.R*255.99))
	g := uint8(math.Min(255.0, c.G*255.99))
	b := uint8(math.Min(255.0, c.B*255.99))
	return color.RGBA{r, g, b, 255}
}

func (c Color) String() string {
	r := uint32(math.Min(255.0, c.R*255.99))
	g := uint32(math.Min(255.0, c.G*255.99))
	b := uint32(math.Min(255.0, c.B*255.99))

	return fmt.Sprintf("%d %d %d", r, g, b)
}
