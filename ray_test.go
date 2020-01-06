package main

import "testing"

func TestRayPointAt(t *testing.T) {
	ray := Ray{Origin: Vector{0.0, 0.0, 0.0}, Direction: Vector{1.0, 0.0, 0.0}.MakeUnitVector()}
	v := ray.PointAt(1.0)
	e := Vector{1.0, 0.0, 0.0}

	if v != e {
		t.Errorf("expected %s but got %s", e, v)
	}
}
