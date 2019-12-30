package main

import "math"

type Sphere struct {
	center   Vector
	radius   float64
	material Material
}

/**
 * Ray intersects with Sphere when:
 * O = Ray.origin
 * D = Ray.direction
 * C = center of sphere
 * r = radius of sphere
 *
 * dot(O + t*D - C, O + t*D - C) = r*r
 * t*t*dot(D, D) + 2*t*dot(D, O-C) + dot(O-D, O-D) - r*r = 0
 *
 * which is a quadratic formula (a*x*x + b*x + c) (with t=x)
 * x = (-b +- sqrt(b*b - 4*a*c))/ 2*a
 *
 * The discriminant is the part of the quadratic formula under the square root.
 *
 * we can leave out some 2's
 */
func (s Sphere) hit(ray *Ray, tMin float64, tMax float64) (bool, *Hit) {
	oc := ray.Origin.Sub(s.center)
	a := Dot(ray.Direction, ray.Direction)
	b := Dot(oc, ray.Direction)
	c := Dot(oc, oc) - s.radius*s.radius
	discriminant := b*b - a*c

	if discriminant > 0 {
		temp := (-b - math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			hitPoint := ray.PointAt(temp)
			hr := Hit{
				t:        temp,
				p:        hitPoint,
				normal:   hitPoint.Sub(s.center).Scale(1.0 / s.radius),
				material: s.material}
			return true, &hr
		}
		temp = (-b + math.Sqrt(discriminant)) / a
		if temp < tMax && temp > tMin {
			hitPoint := ray.PointAt(temp)
			hr := Hit{
				t:        temp,
				p:        hitPoint,
				normal:   hitPoint.Sub(s.center).Scale(1.0 / s.radius),
				material: s.material}
			return true, &hr
		}
	}
	return false, nil
}
