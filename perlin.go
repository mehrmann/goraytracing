package main

import (
	"math"
	"math/rand"
	"sync"
)

type Noise interface {
	noise(v Vector) float64
	turb(p Vector, depth int) float64
}

type perlin struct {
	ranFloat []Vector
	permX    []int
	permY    []int
	permZ    []int
}

var (
	PerlinNoise *perlin
	once        sync.Once
)

func Perlin() *perlin {
	once.Do(func() {
		rnd := rand.New(rand.NewSource(rand.Int63()))
		PerlinNoise = &perlin{
			ranFloat: perlinGenerate(rnd),
			permX:    perlinGeneratePerm(rnd),
			permY:    perlinGeneratePerm(rnd),
			permZ:    perlinGeneratePerm(rnd)}
	})
	return PerlinNoise
}

func perlinGenerate(rnd Rnd) []Vector {
	p := make([]Vector, 256)
	for i := 0; i < 256; i++ {
		p[i] = Vector{
			2.0*rnd.Float64() - 1.0,
			2.0*rnd.Float64() - 1.0,
			2.0*rnd.Float64() - 1.0}.MakeUnitVector()
	}
	return p
}

func perlinGeneratePerm(rnd Rnd) []int {
	p := make([]int, 256)
	for i := 0; i < 256; i++ {
		p[i] = i
	}
	permute(rnd, p, 256)
	return p
}

func permute(rnd Rnd, p []int, n int) {
	for i := n - 1; i > 0; i-- {
		target := int(rnd.Float64() * float64(i+1))
		tmp := p[i]
		p[i] = p[target]
		p[target] = tmp
	}
}

func perlinInterp(c [2][2][2]Vector, u, v, w float64) float64 {
	uu := u * u * (3 - 2*u)
	vv := v * v * (3 - 2*v)
	ww := w * w * (3 - 2*w)
	accum := 0.0
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				weight_v := Vector{u - float64(i), v - float64(j), w - float64(k)}
				accum += (float64(i)*uu + (1.0-float64(i))*(1.0-uu)) *
					(float64(j)*vv + (1.0-float64(j))*(1.0-vv)) *
					(float64(k)*ww + (1.0-float64(k))*(1.0-ww)) * Dot(c[i][j][k], weight_v)
			}
		}
	}
	return accum
}

func (noise perlin) noise(p Vector) float64 {
	u := p.X - math.Floor(p.X)
	v := p.Y - math.Floor(p.Y)
	w := p.Z - math.Floor(p.Z)

	i := int(math.Floor(p.X))
	j := int(math.Floor(p.Y))
	k := int(math.Floor(p.Z))

	var c [2][2][2]Vector
	for di := 0; di < 2; di++ {
		for dj := 0; dj < 2; dj++ {
			for dk := 0; dk < 2; dk++ {
				c[di][dj][dk] = noise.ranFloat[noise.permX[(i+di)&255]^noise.permY[(j+dj)&255]^noise.permZ[(k+dk)&255]]
			}
		}
	}
	return perlinInterp(c, u, v, w)
}

func (noise perlin) turb(p Vector, depth int) float64 {
	accum := 0.0
	temp_p := p
	weight := 1.0
	for i := 0; i < depth; i++ {
		accum += weight * noise.noise(temp_p)
		weight *= 0.5
		temp_p = temp_p.Scale(2.0)
	}
	return math.Abs(accum)
}
