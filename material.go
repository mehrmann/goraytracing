package main

type Material interface {
	scatter(ray *Ray, hit *Hit) (wasScattered bool, attenuation *Color, scattered *Ray)
}

type Lambert struct {
	albedo Color
}

type Metal struct {
	albedo Color
	fuzz   float64
}

type Dielectric struct {
	ri float64
}

func (mat Lambert) scatter(ray *Ray, hit *Hit) (bool, *Color, *Ray) {
	target := hit.p.Add(hit.normal).Add(randomInUnitSphere(ray.rnd))
	scattered := &Ray{hit.p, target.Sub(hit.p), ray.rnd}
	attenuation := &mat.albedo
	return true, attenuation, scattered
}

func (mat Metal) scatter(ray *Ray, hit *Hit) (bool, *Color, *Ray) {
	reflected := reflect(ray.Direction.MakeUnitVector(), hit.normal)
	scattered := &Ray{hit.p, reflected, ray.rnd}
	attenuation := &mat.albedo
	if Dot(scattered.Direction, hit.normal) > 0 {
		return true, attenuation, scattered
	}
	return false, nil, nil
}

func (mat Dielectric) scatter(ray *Ray, hit *Hit) (bool, *Color, *Ray) {

	outward_normal := Vector{}
	reflected := reflect(ray.Direction.MakeUnitVector(), hit.normal)
	ni_over_nt := 0.0
	reflect_prob := 0.0
	cosine := 0.0
	refracted := Vector{}
	scattered := Ray{}
	if Dot(ray.Direction, hit.normal) > 0 {
		outward_normal = hit.normal.Scale(-1)
		ni_over_nt = mat.ri
		cosine = mat.ri * Dot(ray.Direction, hit.normal) / ray.Direction.Length()
	} else {
		outward_normal = hit.normal
		ni_over_nt = 1.0 / mat.ri
		cosine = -Dot(ray.Direction, hit.normal) / ray.Direction.Length()
	}
	if wasRefracted, ref := refract(ray.Direction, outward_normal, ni_over_nt); wasRefracted {
		reflect_prob = schlick(cosine, mat.ri)
		refracted = ref.Scale(1.0)
	} else {
		scattered = Ray{hit.p, reflected, ray.rnd}
		reflect_prob = 1.0
	}
	if ray.rnd.Float64() < reflect_prob {
		scattered = Ray{hit.p, reflected, ray.rnd}
	} else {
		scattered = Ray{hit.p, refracted, ray.rnd}
	}
	return true, &White, &scattered
}
