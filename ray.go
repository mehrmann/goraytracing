package main

type Ray struct {
	Origin    Vector
	Direction Vector
	rnd       Rnd
}

func (r *Ray) PointAt(t float64) Vector {
	return r.Origin.Add(r.Direction.Scale(t))
}
