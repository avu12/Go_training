package main

import (
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten"

	"image/color"
)

const (
	screenWidth, screenHeigth = 640, 360
	boidCount                 = 1000
	viewRadius                = 13
	adjRate                   = 0.015
)

var (
	green   = color.RGBA{10, 255, 50, 255}
	boids   [boidCount]*Boid
	boidMap [screenWidth + 1][screenHeigth + 1]int
	rwlock  = sync.RWMutex{}
)

func update(screen *ebiten.Image) error {
	if !ebiten.IsDrawingSkipped() {
		for i, boid := range boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y), green)
			screen.Set(int(boid.position.x-1), int(boid.position.y), green)
			screen.Set(int(boid.position.x), int(boid.position.y-1), green)
			screen.Set(int(boid.position.x), int(boid.position.y+1), green)
			boid.ttl--
			if boid.ttl == 0 {
				createBoid(i)
			}
		}

	}
	return nil
}

func main() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	err := ebiten.Run(update, screenWidth, screenHeigth, 2, "Boid in box")
	if err != nil {
		log.Fatal(err)
	}

}
