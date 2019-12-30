package main

import (
	"flag"
	"fmt"
	"image"
	"image/png"
	"math"
	"math/rand"
	"os"
	"runtime"
)

type Options struct {
	Width        int
	Height       int
	RaysPerPixel int
	World        int
	Threads      int
	Filename     string
}

type WorldBuildFunction struct {
	Name          string
	BuildFunction func(int, int) (Camera, ObjectList)
}

func buildWorldChapter7(width int, height int) (Camera, ObjectList) {
	lookFrom := Vector{0, 0.0, 3.0}
	lookAt := Vector{Z: -1.0}
	aperture := 0.0
	distToFocus := 2.0
	camera := MakeCamera(lookFrom, lookAt, Up, 20, float64(width)/float64(height), aperture, distToFocus)

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
	camera := MakeCamera(lookFrom, lookAt, Up, 20, float64(width)/float64(height), aperture, distToFocus)

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
	camera := MakeCamera(lookFrom, lookAt, Up, 20, float64(width)/float64(height), aperture, distToFocus)

	world := ObjectList{
		Sphere{center: Vector{Z: -1.0}, radius: 0.5, material: Lambert{Color{R: 0.1, G: 0.2, B: 0.5}}},
		Sphere{center: Vector{Y: -100.5, Z: -1.0}, radius: 100, material: Lambert{Color{R: 0.8, G: 0.8}}},
		Sphere{center: Vector{X: 1.0, Y: 0, Z: -1.0}, radius: 0.5, material: Metal{Color{R: 0.8, G: 0.6, B: 0.2}, 1.0}},
		Sphere{center: Vector{X: -1.0, Y: 0, Z: -1.0}, radius: 0.5, material: Dielectric{1.5}},
		Sphere{center: Vector{X: -1.0, Y: 0, Z: -1.0}, radius: -0.45, material: Dielectric{1.5}},
	}

	return camera, world
}

func buildWorldOneWeekend(width, height int) (Camera, ObjectList) {
	world := ObjectList{}

	maxSpheres := 500
	world = append(world, Sphere{center: Vector{Y: -1000.0}, radius: 1000, material: Lambert{Color{0.5, 0.5, 0.5}}})

	for a := -11; a < 11 && len(world) < maxSpheres; a++ {
		for b := -11; b < 11 && len(world) < maxSpheres; b++ {
			chooseMaterial := rand.Float64()
			center := Vector{float64(a) + 0.9*rand.Float64(), 0.2, float64(b) + 0.9*rand.Float64()}

			if center.Sub(Vector{4.0, 0.2, 0}).Length() > 0.9 {
				switch {
				case chooseMaterial < 0.8: // diffuse
					world = append(world,
						Sphere{
							center:   center,
							radius:   0.2,
							material: Lambert{Color{R: rand.Float64() * rand.Float64(), G: rand.Float64() * rand.Float64(), B: rand.Float64() * rand.Float64()}}})
				case chooseMaterial < 0.95: // metal
					world = append(world,
						Sphere{
							center:   center,
							radius:   0.2,
							material: Metal{Color{R: 0.5 * (1 + rand.Float64()), G: 0.5 * (1 + rand.Float64()), B: 0.5 * (1 + rand.Float64())}, 0.5 * rand.Float64()}})
				default:
					world = append(world,
						Sphere{
							center:   center,
							radius:   0.2,
							material: Dielectric{1.5}})

				}
			}
		}
	}

	world = append(world,
		Sphere{
			center:   Vector{0, 1, 0},
			radius:   1.0,
			material: Dielectric{1.5}},
		Sphere{
			center:   Vector{-4, 1, 0},
			radius:   1.0,
			material: Lambert{Color{0.4, 0.2, 0.1}}},
		Sphere{
			center:   Vector{4, 1, 0},
			radius:   1.0,
			material: Metal{Color{0.7, 0.6, 0.5}, 0}})

	lookFrom := Vector{13, 2, 3}
	lookAt := Vector{}
	aperture := 0.01
	distToFocus := (lookFrom.Sub(lookAt)).Length()
	camera := MakeCamera(lookFrom, lookAt, Up, 20, float64(width)/float64(height), aperture, distToFocus)

	return camera, world
}

func getColorOfRay(world ObjectList, ray *Ray, depth int) Color {
	if hitAnObject, hit := world.hit(ray, 0.001, math.MaxFloat64); hitAnObject {
		if depth > 50 {
			return Black
		}
		if wasScattered, attenuation, scattered := hit.material.scatter(ray, hit); wasScattered {
			return attenuation.Mult(getColorOfRay(world, scattered, depth+1))
		} else {
			return Black
		}

	}
	unitDirection := ray.Direction.MakeUnitVector()
	t := 0.5 * (unitDirection.Y + 1.0)
	return White.Scale(1.0 - t).Add(SkyBlue.Scale(t))
}

type PixelJob struct {
	x, y, passes int
	color        Color
}

func pixelRenderer(in chan PixelJob, out chan PixelJob, quit chan int, options Options, world ObjectList, camera Camera) {
	rnd := rand.New(rand.NewSource(rand.Int63()))
	for {
		select {
		case job := <-in:
			for ns := 0; ns < job.passes; ns++ {
				u := (float64(job.x) + rnd.Float64()) / float64(options.Width)
				v := (float64(job.y) + rnd.Float64()) / float64(options.Height)
				ray := camera.ray(rnd, u, v)
				job.color = job.color.Add(getColorOfRay(world, ray, 0))
			}
			job.color = job.color.Scale(1.0 / float64(job.passes))
			out <- job
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	options := Options{}

	worlds := []WorldBuildFunction{
		WorldBuildFunction{"Chapter 7", buildWorldChapter7},
		WorldBuildFunction{"Metal", buildWorldMetalSpheres},
		WorldBuildFunction{"Dieletrics", buildWorldDielectrics},
		WorldBuildFunction{"OneWeekend", buildWorldOneWeekend}}

	flag.IntVar(&options.Width, "w", 800, "width of rendered image")
	flag.IntVar(&options.Height, "h", 400, "height of rendered Image")
	flag.IntVar(&options.RaysPerPixel, "rays", 8, "passes per pixel")
	flag.IntVar(&options.World, "world", 0, fmt.Sprintf("select which world to render (0-%d)", len(worlds)-1))
	flag.IntVar(&options.Threads, "t", runtime.NumCPU(), "number of parallel threads")
	flag.StringVar(&options.Filename, "filename", "out.png", "the name of the output png")
	flag.Parse()

	rand.Seed(1337)
	//rnd := rand.New(rand.NewSource(rand.Int63()))

	options.World = min(len(worlds)-1, options.World)
	camera, world := worlds[options.World].BuildFunction(options.Width, options.Height)
	fmt.Printf("Rendering '%s' in %dx%d with %d passes on %d CPUs\n", worlds[options.World].Name, options.Width, options.Height, options.RaysPerPixel, options.Threads)
	image := image.NewRGBA(image.Rect(0, 0, options.Width, options.Height))

	jobChannel := make(chan PixelJob)
	resultChannel := make(chan PixelJob)
	quitChannel := make(chan int)

	for i := 0; i < options.Threads; i++ {
		go pixelRenderer(jobChannel, resultChannel, quitChannel, options, world, camera)
	}

	for i := 0; i < options.Threads; i++ {
		go func() {
			for {
				result := <-resultChannel
				image.SetRGBA(result.x, options.Height-1-result.y, result.color.asRGBA())
			}
		}()
	}

	for ny := 0; ny < options.Height; ny++ {
		for nx := 0; nx < options.Width; nx++ {
			jobChannel <- PixelJob{x: nx, y: ny, passes: options.RaysPerPixel, color: Black}
		}
	}

	outputFile, err := os.Create(options.Filename)
	if err != nil {
		//Handle error
	}
	defer outputFile.Close()

	png.Encode(outputFile, image)

}
