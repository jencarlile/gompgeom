package gompgeom

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X, Y float64
}

func (p1 *Point) Add(p2 *Point) *Point {
	return &Point{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 *Point) Sub(p2 *Point) *Point {
	return &Point{p1.X - p2.X, p1.Y - p2.Y}
}

func (p1 *Point) Dot(p2 *Point) float64 {
	return (p1.X * p2.X) + (p1.Y * p2.Y)
}

func (p *Point) Length() float64 {
	return math.Sqrt(p.Dot(p))
}

func (p1 *Point) Distance(p2 *Point) float64 {
	return p1.Sub(p2).Length()
}

//     | a  b  c |
// D = | d  e  f |
//     | g  h  i |
func Determinant(a, b, c, d, e, f, g, h, i float64) float64 {
	return (a * e * i) + (b * f * g) + (c * d * h) - (c * e * g) - (b * d * i) - (a * f * h)
}

func ToRight(p1, p2, p3 *Point) bool {
	return Determinant(
		1, p1.X, p1.Y,
		1, p2.X, p2.Y,
		1, p3.X, p3.Y) < 0
}

type Points []*Point

func (pts Points) Len() int      { return len(pts) }
func (pts Points) Swap(i, j int) { pts[i], pts[j] = pts[j], pts[i] }

func printPoints(pt2s []*Point) {
	for _, p := range pt2s {
		fmt.Printf("(%v,%v) ", p.X, p.Y)
	}
	fmt.Println("")
}

type ByX struct{ Points }
type ByY struct{ Points }

func (s ByX) Less(i, j int) bool {
	if s.Points[i].X == s.Points[j].X {
		return s.Points[i].Y < s.Points[j].Y
	}
	return s.Points[i].X < s.Points[j].X
}

func (s ByY) Less(i, j int) bool {
	if s.Points[i].Y == s.Points[j].Y {
		return s.Points[i].X < s.Points[j].X
	}
	return s.Points[i].Y < s.Points[j].Y
}

func ConvexHull(points []*Point) Polygon {
	// TODO(jcarlile): Remove duplicates?
	ptsLen := len(points)
	if ptsLen < 3 {
		return Polygon{}
	}

	// Sort points by x-coordinate
	sort.Sort(ByX{points})

	// Find the upper hull
	upperHull := []*Point{points[0], points[1]}
	for idx := 2; idx < len(points); idx++ {
		for true {
			uhLen := len(upperHull)
			if ToRight(upperHull[uhLen-2], upperHull[uhLen-1], points[idx]) {
				upperHull = append(upperHull, points[idx])
				break
			} else {
				// delete last point in upperHull, recheck
				upperHull = append(upperHull[:uhLen-1])

				// Ensure there are always at least two points in the upper hull
				if len(upperHull) < 2 {
					upperHull = append(upperHull, points[idx])
					break
				}
			}
		}
	}

	// Find the lower hull
	lowerHull := []*Point{points[ptsLen-1], points[ptsLen-2]}
	for idx := ptsLen - 3; idx >= 0; idx-- {
		for true {
			lhLen := len(lowerHull)

			if ToRight(lowerHull[lhLen-2], lowerHull[lhLen-1], points[idx]) {
				lowerHull = append(lowerHull, points[idx])
				break
			} else {
				// delete the last point in lower hull
				lowerHull = append(lowerHull[:lhLen-1])
				// Ensure there are always at least two points in the lower hull
				if len(lowerHull) < 2 {
					lowerHull = append(lowerHull, points[idx])
					break
				}
			}
		}
	}

	// Remove first and last point from lower hull to avoid duplication of upper/lower hull intersection
	lowerHull = append(lowerHull[1 : len(lowerHull)-1])
	// Append lower hull to upper hull to form the convex hull.
	convexHull := append(upperHull, lowerHull...)
	return convexHull
}
