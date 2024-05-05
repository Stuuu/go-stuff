package main

import "fmt"

type Point struct {
	x, y float64
}

func (p Point) Offset(x, y float64) Point {
	return Point{p.x+x, p.y+y}
}

func (p *Point) Move(x, y float64) {
	p.x += x
	p.y += y
}

func main() {

	p1 := Point{1, 2}
	fmt.Println(p1)
	p1 =p1.Offset(1, 1)
	fmt.Println(p1)
	p1.Move(1,1)
	fmt.Println(p1)
}