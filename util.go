package main

import (
	"math"
)

type Rnd interface {
	Float64() float64
}

func randomInUnitSphere(rnd Rnd) Vector {
	x := rnd.Float64()
	y := rnd.Float64()
	z := rnd.Float64()
	s := math.Sqrt(1.0 / (x*x + y*y + z*z))
	return Vector{s * x, s * y, s * z}
}

func randomInUnitDisk(rnd Rnd) Vector {
	phi := rnd.Float64() * math.Pi * 2.0
	r := math.Sqrt(rnd.Float64())
	return Vector{r * math.Cos(phi), r * math.Sin(phi), 0.0}
}

/**
 * \     /|
 *  \v r/ | B
 *   \ /  |
 *  --*-----
 *     \  |
 *      \v| B
 *       \|
 *
 * length of B = dot(v,n)
 * direction of B is n
 */
func reflect(v, n Vector) Vector {
	return v.Sub(n.Scale(2.0 * Dot(v, n)))
}

func refract(v, n Vector, niOverNt float64) (bool, *Vector) {
	uv := v.MakeUnitVector()
	dt := Dot(uv, n)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0 {
		refracted := uv.Sub(n.Scale(dt)).Scale(niOverNt).Sub(n.Scale(math.Sqrt(discriminant)))
		return true, &refracted
	}
	return false, nil
}

func schlick(cosine, ref_idx float64) float64 {
	r0 := (1.0 - ref_idx) / (1.0 + ref_idx)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow((1.0-cosine), 5)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
