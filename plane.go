package main

type Plane struct {
	center   Vector
	normal   Vector
	material Material
}

func (p Plane) hit(ray *Ray, tMin float64, tMax float64) (bool, *Hit) {
	denom := Dot(p.normal, ray.Direction)
	if denom > 1e-6 {
		centerToRayOrigin := p.center.Sub(ray.Origin)
		temp := Dot(centerToRayOrigin, p.normal) / denom
		if temp < tMax && temp > tMin {
			hitPoint := ray.PointAt(temp)

			hr := Hit{
				t:        temp,
				p:        hitPoint,
				normal:   p.normal,
				material: p.material}
			return true, &hr
		}
	}
	return false, nil
}
