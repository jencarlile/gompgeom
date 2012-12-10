package core

import "math"

type Polygon2 []*Point2

func (p Polygon2) Area() float64 {
	// http://softsurfer.com/Archive/algorithm_0101/algorithm_0101.htm
	var area float64
	nVert := len(p)
	for i := 0; i < nVert; i++ {
		area += p[(i+1)%nVert].X * (p[(i+2)%nVert].Y - p[i].Y)
	}

	return math.Abs(area) / 2.0
}
