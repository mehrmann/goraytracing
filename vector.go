package main

import "math"

type Vector struct {
	X, Y, Z float64
}

func (v Vector) Scale(f float64) Vector {
	return Vector{v.X * f, v.Y * f, v.Z * f}
}

func (v Vector) Multiply(w Vector) Vector {
	return Vector{v.X * w.X, v.Y * w.Y, v.Z * w.Z}
}

func (v Vector) Sub(w Vector) Vector {
	return Vector{v.X - w.X, v.Y - w.Y, v.Z - w.Z}
}

func (v Vector) Add(w Vector) Vector {
	return Vector{v.X + w.X, v.Y + w.Y, v.Z + w.Z}
}

func (v Vector) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y + v.Z*v.Z)
}

func (v Vector) SquaredLength() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

func (v Vector) MakeUnitVector() Vector {
	return v.Scale(1.0 / v.Length())
}

func (v Vector) Negate() Vector {
	return Vector{-v.X, -v.Y, -v.Z}
}

func Dot(a Vector, b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Cross(a Vector, b Vector) Vector {
	return Vector{
		(a.Y*b.Z - a.Z*b.Y),
		-(a.X*b.Z - a.Z*b.X),
		(a.X*b.Y - a.Y*b.X)}
}
