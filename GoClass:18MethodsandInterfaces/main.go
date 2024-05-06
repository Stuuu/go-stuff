package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

type Line struct {
	Begin, End Point
}

type Path []Point

func (l Line) Distance() float64 {
	return math.Hypot(l.End.X - l.Begin.X, l.End.Y - l.Begin.Y)
}


func (l Line) ScaleBy(f float64) Line {
	l.End.X += (f-1)*(l.End.X - l.Begin.X)
	l.End.Y += (f-1)*(l.End.Y - l.Begin.Y)
	return Line{l.Begin, Point{l.End.X, l.End.Y}}
}

func main() {
	side := Line{Point{1,2}, Point{4,6}}

	s2 := side.ScaleBy(2.5)
	fmt.Println(s2.Distance())
	fmt.Println(Line{Point{1,2}, Point{4,6}}.ScaleBy(2).Distance())	

}