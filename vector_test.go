package main

import (
	"testing"
)

func TestVectorOps(t *testing.T) {
	tests := []struct {
		in, out Vector
		op      func(Vector) Vector
	}{
		{in: Vector{1.0, 1.0, 1.0}, out: Vector{1.0, 1.0, 1.0}, op: func(in Vector) Vector { return in.Scale(1.0) }},
		{in: Vector{1.0, 1.0, 1.0}, out: Vector{2.0, 2.0, 2.0}, op: func(in Vector) Vector { return in.Scale(2.0) }},
		{in: Vector{1.0, 1.0, 1.0}, out: Vector{0.5, 0.5, 0.5}, op: func(in Vector) Vector { return in.Scale(0.5) }},

		{in: Vector{1.0, 1.0, 1.0}, out: Vector{1.0, 1.0, 1.0}, op: func(in Vector) Vector { return in.Multiply(Vector{1.0, 1.0, 1.0}) }},
		{in: Vector{1.0, 1.0, 1.0}, out: Vector{1.1, 1.2, 1.3}, op: func(in Vector) Vector { return in.Add(Vector{0.1, 0.2, 0.3}) }},
		{in: Vector{1.0, 1.0, 1.0}, out: Vector{0.9, 0.8, 0.7}, op: func(in Vector) Vector { return in.Sub(Vector{0.1, 0.2, 0.3}) }},

		{in: Vector{5.0, 0.0, 0.0}, out: Vector{1.0, 0.0, 0.0}, op: func(in Vector) Vector { return in.MakeUnitVector() }},

		{in: Vector{5.0, -1.0, 0.0}, out: Vector{-5.0, 1.0, 0.0}, op: func(in Vector) Vector { return in.Negate() }},

		{in: Vector{1.0, 1.0, 1.0}, out: Vector{0.0, 0.0, 0.0}, op: func(in Vector) Vector { return Cross(in, Vector{1.0, 1.0, 1.0}) }},
		{in: Vector{1.0, 0.0, 0.0}, out: Vector{0.0, 0.0, 1.0}, op: func(in Vector) Vector { return Cross(in, Vector{0.0, 1.0, 0.0}) }},
		{in: Vector{0.0, 1.0, 1.0}, out: Vector{-1.0, 1.0, -1.0}, op: func(in Vector) Vector { return Cross(in, Vector{1.0, 1.0, 0.0}) }},
	}

	for idx, test := range tests {
		r := test.op(test.in)
		if r != test.out {
			t.Errorf("Test %d: expected %s but got %s", idx, test.out, r)
		}
	}
}

func TestVectorOps2(t *testing.T) {
	tests := []struct {
		in  Vector
		out float64
		op  func(Vector) float64
	}{
		{in: Vector{42.0, 0.0, 0.0}, out: 42.0, op: func(in Vector) float64 { return in.Length() }},
		{in: Vector{2.0, 3.0, 7.0}, out: 62.0, op: func(in Vector) float64 { return in.SquaredLength() }},
		{in: Vector{0.0, 1.0, 0.0}, out: 0.0, op: func(in Vector) float64 { return Dot(in, Vector{1.0, 0.0, 0.0}) }},
		{in: Vector{1.0, 0.0, 0.0}, out: 1.0, op: func(in Vector) float64 { return Dot(in, Vector{1.0, 0.0, 0.0}) }},
		{in: Vector{0.0, 0.0, 1.0}, out: 0.0, op: func(in Vector) float64 { return Dot(in, Vector{1.0, 0.0, 0.0}) }},
	}

	for idx, test := range tests {
		r := test.op(test.in)
		if r != test.out {
			t.Errorf("Test %d: expected %f but got %f", idx, test.out, r)
		}
	}
}
