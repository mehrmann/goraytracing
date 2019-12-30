package main

type Hit struct {
	t        float64
	p        Vector
	normal   Vector
	material Material
}

type Object interface {
	hit(ray *Ray, tMin float64, tMax float64) (bool, *Hit)
}

type ObjectList []Object

func (list ObjectList) hit(r *Ray, tMin float64, tMax float64) (bool, *Hit) {
	var resultHit *Hit
	hitAnything := false

	closestSoFar := tMax

	for _, h := range list {
		if wasHit, hit := h.hit(r, tMin, closestSoFar); wasHit {
			hitAnything = true
			resultHit = hit
			closestSoFar = hit.t
		}
	}
	return hitAnything, resultHit
}
