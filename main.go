package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
)

type Options struct {
	Width        int
	Height       int
	RaysPerPixel int
	World        int
}

func buildWorldChapter7(width int, height int) (Camera, ObjectList) {
	lookFrom := Vector{0, 0.0, 3.0}
	lookAt := Vector{Z: -1.0}
	aperture := 0.0
	distToFocus := 2.0
	camera := MakeCamera(lookFrom, lookAt, Vector{Y: 1.0}, 20, float64(width)/float64(height), aperture, distToFocus)

	world := ObjectList{
		Sphere{center: Vector{Z: -1.0}, radius: 0.5, material: Lambert{Color{R: 1.0}}},
		Sphere{center: Vector{Y: -100.5, Z: -1.0}, radius: 100, material: Lambert{Color{G: 1.0}}},
	}
	return camera, world
}

func buildWorldMetalSpheres(width, height int) (Camera, ObjectList) {
	lookFrom := Vector{0, 0.0, 3.0}
	lookAt := Vector{Z: -1.0}
	aperture := 0.0
	distToFocus := 1.0
	camera := MakeCamera(lookFrom, lookAt, Vector{Y: 1.0}, 20, float64(width)/float64(height), aperture, distToFocus)

	world := ObjectList{
		Sphere{center: Vector{Z: -1.0}, radius: 0.5, material: Lambert{Color{R: 0.8, G: 0.3, B: 0.3}}},
		Sphere{center: Vector{Y: -100.5, Z: -1.0}, radius: 100, material: Lambert{Color{R: 0.8, G: 0.8}}},
		Sphere{center: Vector{X: 1.0, Y: 0, Z: -1.0}, radius: 0.5, material: Metal{Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}},
		Sphere{center: Vector{X: -1.0, Y: 0, Z: -1.0}, radius: 0.5, material: Metal{Color{R: 0.8, G: 0.8, B: 0.8}, 0.3}},
	}

	return camera, world
}

func buildWorldDielectrics(width, height int) (Camera, ObjectList) {
	lookFrom := Vector{-2.0, 2.0, 1.0}
	lookAt := Vector{Z: -1.0}
	aperture := 0.0
	distToFocus := 1.0
	camera := MakeCamera(lookFrom, lookAt, Vector{Y: 1.0}, 20, float64(width)/float64(height), aperture, distToFocus)

	world := ObjectList{
		Sphere{center: Vector{Z: -1.0}, radius: 0.5, material: Lambert{Color{R: 0.1, G: 0.2, B: 0.5}}},
		Sphere{center: Vector{Y: -100.5, Z: -1.0}, radius: 100, material: Lambert{Color{R: 0.8, G: 0.8}}},
		Sphere{center: Vector{X: 1.0, Y: 0, Z: -1.0}, radius: 0.5, material: Metal{Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}},
		Sphere{center: Vector{X: -1.0, Y: 0, Z: -1.0}, radius: 0.5, material: Dielectric{1.5}},
		Sphere{center: Vector{X: -1.0, Y: 0, Z: -1.0}, radius: -0.45, material: Dielectric{1.5}},
	}

	return camera, world
}

func color(world ObjectList, ray *Ray, depth int) Color {
	if hitAnObject, hit := world.hit(ray, 0.001, math.MaxFloat64); hitAnObject {
		if depth > 50 {
			return Black
		}
		if wasScattered, attenuation, scattered := hit.material.scatter(ray, hit); wasScattered {
			return attenuation.Mult(color(world, scattered, depth+1))
		} else {
			return Black
		}

	}
	unitDirection := ray.Direction.MakeUnitVector()
	t := 0.5 * (unitDirection.Y + 1.0)
	return White.Scale(1.0 - t).Add(SkyBlue.Scale(t))
}

func main() {
	options := Options{}

	worlds := []func(width int, height int) (Camera, ObjectList){buildWorldChapter7, buildWorldMetalSpheres, buildWorldDielectrics}

	flag.IntVar(&options.Width, "w", 800, "width of rendered image")
	flag.IntVar(&options.Height, "h", 400, "height of rendered Image")
	flag.IntVar(&options.RaysPerPixel, "rays", 8, "passes per pixel")
	flag.IntVar(&options.World, "world", 0, fmt.Sprintf("select which world to render (0-%d)", len(worlds)-1))

	flag.Parse()

	rand.Seed(1337)
	rnd := rand.New(rand.NewSource(rand.Int63()))

	options.World = min(len(worlds)-1, options.World)
	camera, world := worlds[options.World](options.Width, options.Height)

	fmt.Println("P3")
	fmt.Println(options.Width, " ", options.Height)
	fmt.Println("255")
	for ny := options.Height; ny > 0; ny-- {
		for nx := 0; nx < options.Width; nx++ {
			c := Color{}
			for ns := 0; ns < options.RaysPerPixel; ns++ {
				u := (float64(nx) + rnd.Float64()) / float64(options.Width)
				v := (float64(ny) + rnd.Float64()) / float64(options.Height)
				ray := camera.ray(rnd, u, v)
				c = c.Add(color(world, ray, 0))
			}
			c = c.Scale(1.0 / float64(options.RaysPerPixel))
			fmt.Println(c)
		}
	}

}
