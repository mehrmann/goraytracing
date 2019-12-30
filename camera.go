package main

import "math"

type Camera interface {
	ray(rnd Rnd, u float64, v float64) *Ray
}

type FullyFletchedCamera struct {
	origin          Vector
	lowerLeftCorner Vector
	horizontal      Vector
	vertical        Vector
	u, v, w         Vector
	lensRadius      float64
}

func MakeCamera(lookFrom, lookAt, up Vector, vfov, aspect, aperture, focusDistance float64) Camera {
	lensRadius := aperture / 2.0
	theta := vfov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight

	origin := lookFrom
	w := lookFrom.Sub(lookAt).MakeUnitVector()
	u := Cross(up, w).MakeUnitVector()
	v := Cross(w, u)
	lowerLeftCorner := origin.Sub(u.Scale(halfWidth * focusDistance)).Sub(v.Scale(halfHeight * focusDistance)).Sub(w.Scale(focusDistance))
	//lowerLeftCorner := origin.Translate(u.Scale(-(halfWidth * focusDist))).Translate(v.Scale(-(halfHeight * focusDist))).Translate(w.Scale(-focusDist))

	horizontal := u.Scale(2.0 * halfWidth * focusDistance)
	vertical := v.Scale(2.0 * halfHeight * focusDistance)

	return FullyFletchedCamera{origin: origin, lowerLeftCorner: lowerLeftCorner, horizontal: horizontal, vertical: vertical, u: u, v: v, w: w, lensRadius: lensRadius}
}

func (c FullyFletchedCamera) ray(rnd Rnd, u float64, v float64) *Ray {
	direction := c.lowerLeftCorner.Add(c.horizontal.Scale(u)).Add(c.vertical.Scale(v)).Sub(c.origin)
	origin := c.origin
	if c.lensRadius > 0 {
		rd := randomInUnitDisk(rnd).Scale(c.lensRadius)
		offset := c.u.Scale(rd.X).Add(c.v.Scale(rd.Y))
		origin.Add(offset)
		direction = direction.Sub(offset)
	}
	return &Ray{origin, direction, rnd}
}
