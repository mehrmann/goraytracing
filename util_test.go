package main

import "testing"

func TestReflect(t *testing.T) {
	tests := []struct {
		v, n Vector
		out  Vector
	}{
		{v: Vector{-1.0, -1.0, 0.0}, n: Vector{0.0, 1.0, 0.0}, out: Vector{-1.0, 1.0, 0.0}},
		{v: Vector{-2.0, -1.0, 0.0}, n: Vector{0.0, 1.0, 0.0}, out: Vector{-2.0, 1.0, 0.0}},
		{v: Vector{0.0, -1.0, -2.0}, n: Vector{0.0, 1.0, 0.0}, out: Vector{0.0, 1.0, -2.0}},
	}

	for idx, test := range tests {
		r := reflect(test.v, test.n)
		if r != test.out {
			t.Errorf("Test %d: expected %s but got %s", idx, test.out, r)
		}
	}
}

func TestRefract(t *testing.T) {
	tests := []struct {
		op        func() (bool, *Vector)
		refracted bool
		out       *Vector
	}{
		{func() (bool, *Vector) {
			return refract(Vector{0.0, 0.0, 0.0}, Vector{0.0, 1.0, 0.0}, 1.5)
		}, false, nil},
		{func() (bool, *Vector) {
			return refract(Vector{0.1, -0.3, 0.0}, Vector{0.0, 1.0, 0.0}, 1.5)
		}, true, &Vector{0.474342, -0.880341, 0.000000}},
	}

	for idx, test := range tests {
		ref, out := test.op()
		if ref != test.refracted || test.out != nil && *out == *test.out {
			t.Errorf("Test %d: expected %t,%s but got %t,%s", idx, test.refracted, test.out, ref, out)
		}
	}
}
