package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.x - q.x, p.y - q.y)
}

func (p *Point) Scale(n float64) {
	p.x *= n
	p.y *= n
}

func main() {
	pv := Point{1, 1}
	qv := Point{3, 4}
	pp := &pv

	// 1.
	// parameter is T and argument is T
	fmt.Println(pv.Distance(qv))
	// parameter is *T and argument is *T
	pp.Scale(10)
	fmt.Println(*pp)

	// 2.
	// parameter is *T and argument is T
	pv.Scale(0.5)
	fmt.Println(pv)

	// 3.
	// parameter is T and argument is *T
	fmt.Println(pp.Distance(qv))
}