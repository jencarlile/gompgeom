package core

import (
	"fmt"
	"math"
	"sort"
)

type Point2 struct {
	X, Y float64
}

type Point3 struct {
	X, Y, Z float64
}

func (p *Point2) Point3() *Point3 {
	return &Point3{p.X, p.Y, 0}
}

func (p1 *Point2) Add(p2 *Point2) *Point2 {
	return &Point2{p1.X + p2.X, p1.Y + p2.Y}
}

func (p1 *Point2) Sub(p2 *Point2) *Point2 {
	return &Point2{p1.X - p2.X, p1.Y - p2.Y}
}

func (p1 *Point2) Dot(p2 *Point2) float64 {
	return (p1.X * p2.X) + (p1.Y * p2.Y)
}

func (p *Point2) Length() float64 {
	return math.Sqrt(p.Dot(p))
}

func (p1 *Point2) Distance(p2 *Point2) float64 {
	return p1.Sub(p2).Length()
}

//     | a  b  c |
// D = | d  e  f |
//     | g  h  i |
func Determinant(a, b, c, d, e, f, g, h, i float64) float64 {
	return (a * e * i) + (b * f * g) + (c * d * h) - (c * e * g) - (b * d * i) - (a * f * h)
}

func ToRight(p1, p2, p3 *Point2) bool {
	return Determinant(
		1, p1.X, p1.Y,
		1, p2.X, p2.Y,
		1, p3.X, p3.Y) < 0
}

type Point2s []*Point2

func (pt2s Point2s) Len() int      { return len(pt2s) }
func (pt2s Point2s) Swap(i, j int) { pt2s[i], pt2s[j] = pt2s[j], pt2s[i] }
func printPoint2s(pt2s []*Point2) {
	for _, p := range pt2s {
		fmt.Printf("(%v,%v) ", p.X, p.Y)
	}
	fmt.Println("")
}

type ByX struct{ Point2s }

func (s ByX) Less(i, j int) bool {
	if s.Point2s[i].X == s.Point2s[j].X {
		return s.Point2s[i].Y < s.Point2s[j].Y
	}
	return s.Point2s[i].X < s.Point2s[j].X
}

func ConvexHull(points []*Point2) Polygon2 {
	// TODO(jcarlile): Remove duplicates?
	ptsLen := len(points)
	if ptsLen < 3 {
		return Polygon2{}
	}

	// Sort points by x-coordinate
	sort.Sort(ByX{points})

	// Find the upper hull
	upperHull := []*Point2{points[0], points[1]}
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
	lowerHull := []*Point2{points[ptsLen-1], points[ptsLen-2]}
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
