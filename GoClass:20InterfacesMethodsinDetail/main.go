package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.x -p.x, q.y - p.y)
}
func main() {
	p := Point{1, 1}
	q := Point{5, 4}
	

	fmt.Println(p.Distance(q))
	distanceFromP := p.Distance
	
	
	p = Point{2, 2}
	fmt.Println(distanceFromP(q))
}