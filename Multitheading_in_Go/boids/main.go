package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"image/color"
)

const (
	screenWidth, screenHeigth = 640, 360
	boidCount                 = 5
)

var (
	green = color.RGBA{10, 255, 50, 255}
	boids [boidCount]*Boid
)

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for _, boid := range boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y), green)
			screen.Set(int(boid.position.x-1), int(boid.position.y), green)
			screen.Set(int(boid.position.x), int(boid.position.y-1), green)
			screen.Set(int(boid.position.x), int(boid.position.y+1), green)
		}

	}
	return nil
}

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	err := ebiten.Run(update, screenWidth, screenHeigth, 2, "Boid in box")
	if err != nil {
		log.Fatal(err)
	}

}
