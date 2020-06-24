package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	heigth float64
	base   float64
}
type shape interface {
	getArea() float64
}

func main() {
	s := square{}
	s.sideLength = 3.4
	t := triangle{}
	t.base = 2
	t.heigth = 6.7
	fmt.Println(s.getArea(), t.getArea())
	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.heigth
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}
