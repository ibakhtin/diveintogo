package main

import (
	"fmt"
	"image/color"
	"github.com/ibakhtin/diveintogo/methods/geom"
)

func main() {
	// Chapter 6.3
	var cp geom.ColoredPoint
	cp.X = 3
	cp.Point.Y = 4
	cp.Color = color.RGBA{A: 255}
	fmt.Println(cp)

	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	var p1 = geom.ColoredPoint{Point: geom.Point{X: 1, Y: 1}, Color: red}
	var p2 = geom.ColoredPoint{Point: geom.Point{X: 5, Y: 4}, Color: blue}
	fmt.Println(p1.Distance(p2.Point))
	p1.ScaleBy(2)
	p2.ScaleBy(2)
	fmt.Println(p1.Distance(p2.Point))

	p3 := geom.ColoredPointPtr{Point: &geom.Point{X: 1, Y: 1}, Color: red}
	p4 := geom.ColoredPointPtr{Point: &geom.Point{X: 5, Y: 4}, Color: blue}
	fmt.Println(p3.Distance(*p4.Point))
	p4.Point = p3.Point
	p3.ScaleBy(2)
	fmt.Println(p3.Point, p4.Point)

	p := geom.Point{X: 1, Y: 2}
	q := geom.Point{X: 4, Y: 6}

	distance := geom.Point.Distance
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)
}