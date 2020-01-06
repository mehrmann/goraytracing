package main

import (
	"image"
	"math"
)

type Texture interface {
	texture(u, v float64, p Vector) Color
}

type ConstantTexture struct {
	albedo Color
}

type ImageTexture struct {
	image image.Image
}

type CheckerTexture struct {
	odd  Texture
	even Texture
	size float64
}

func (c ConstantTexture) texture(u, v float64, p Vector) Color {
	return c.albedo
}

func (i ImageTexture) texture(u, v float64, p Vector) Color {
	bounds := i.image.Bounds()
	x := int(u * float64(bounds.Max.X))
	y := int((1.0 - v) * float64(bounds.Max.Y))

	r, g, b, _ := i.image.At(x, y).RGBA()
	off := float64(1 << 16)
	return Color{R: float64(r) / off, G: float64(g) / off, B: float64(b) / off}
}

func (c CheckerTexture) texture(u, v float64, p Vector) Color {
	sines := math.Sin(1.0/c.size*p.X) * math.Sin(1.0/c.size*p.Y) * math.Sin(1.0/c.size*p.Z)
	if sines < 0 {
		return c.odd.texture(u, v, p)
	} else {
		return c.even.texture(u, v, p)
	}
}
