package geom

import (
	"image/color"
	"math"
)

// Point type
type Point struct {
	X, Y float64
}

// Distance method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X- q.X, p.Y- q.Y)
}

// Add method
func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// Sub method
func (p Point) Sub(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

// ScaleBy method
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// ColoredPoint type
type ColoredPoint struct {
	Point
	Color color.RGBA
}

// ColoredPointPtr type
type ColoredPointPtr struct {
	*Point
	Color color.RGBA
}

// Path type
type Path []Point