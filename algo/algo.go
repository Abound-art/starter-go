// Package algo implements a Lorenz attractor, see
// https://en.wikipedia.org/wiki/Lorenz_system
package algo

import (
	"image"
	"image/color"
	"math"
)

type Config struct {
	Beta       float64 `json:"beta"`
	Rho        float64 `json:"rho"`
	Sigma      float64 `json:"sigma"`
	DT         float64 `json:"dt"`
	Iterations int     `json:"iterations"`
	ResultSize int     `json:"result_size"`
}

func Run(c *Config) image.Image {
	points := make([]*Point, c.Iterations)
	points[0] = &Point{1, 1, 1}
	bounds := &Bounds{&Point{1, 1, 1}, &Point{1, 1, 1}}
	for iteration := 1; iteration < c.Iterations; iteration++ {
		point := c.NextStep(points[iteration-1])
		points[iteration] = point
		bounds.Expand(point)
	}
	for i, point := range points {
		points[i] = bounds.Translate(point, c.ResultSize)
	}
	counts := make([][]int, c.ResultSize)
	for i := 0; i < c.ResultSize; i++ {
		counts[i] = make([]int, c.ResultSize)
	}
	maxCount := 0
	for _, point := range points {
		x := int(math.Floor(point.X))
		y := int(math.Floor(point.Y))
		counts[x][y]++
		if counts[x][y] > maxCount {
			maxCount = counts[x][y]
		}
	}
	img := image.NewRGBA(image.Rect(0, 0, c.ResultSize, c.ResultSize))
	for i := range counts {
		for j := range counts[i] {
			count := counts[i][j]
			if count == 0 {
				img.SetRGBA(i, j, color.RGBA{A: 255, R: 0, G: 0, B: 0})
			} else {
				pos := math.Sqrt(math.Sqrt(float64(count) / float64(maxCount)))
				b := uint8(pos*200) + 55
				g := uint8((1-pos)*200) + 55
				img.SetRGBA(i, j, color.RGBA{A: 255, G: g, B: b, R: 0})
			}
		}
	}
	return img
}

type Point struct {
	X, Y, Z float64
}

func (c *Config) NextStep(p *Point) *Point {
	dxdt := c.Sigma * (p.Y - p.X)
	dydt := p.X*(c.Rho-p.Z) - p.Y
	dzdt := p.X*p.Y - c.Beta*p.Z
	return &Point{
		X: p.X + dxdt*c.DT,
		Y: p.Y + dydt*c.DT,
		Z: p.Z + dzdt*c.DT,
	}
}

type Bounds struct {
	Min *Point
	Max *Point
}

func (b *Bounds) Expand(p *Point) {
	if p.X < b.Min.X {
		b.Min.X = p.X
	}
	if p.Y < b.Min.Y {
		b.Min.Y = p.Y
	}
	if p.Z < b.Min.Z {
		b.Min.Z = p.Z
	}
	if p.X > b.Max.X {
		b.Max.X = p.X
	}
	if p.Y > b.Max.Y {
		b.Max.Y = p.Y
	}
	if p.Z > b.Max.Z {
		b.Max.Z = p.Z
	}
}

func (b *Bounds) Translate(p *Point, resultSize int) *Point {
	relX := (p.X - b.Min.X) / (b.Max.X - b.Min.X)
	relY := (p.Y - b.Min.Y) / (b.Max.Y - b.Min.Y)
	relZ := (p.Z - b.Min.Z) / (b.Max.Z - b.Min.Z)
	s := float64(resultSize - 1)
	return &Point{
		X: relX * s,
		Y: relY * s,
		Z: relZ * s,
	}
}
