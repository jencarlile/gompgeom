package gompgeom

import "sort"

type Line struct {
	Pts []*Point
}

func NewLine(p1, p2 *Point) *Line {
	ln := []*Point{p1, p2}
	sort.Sort(ByY{ln})
	return &Line{ln}
}

type SegPoint struct {
	Pt  Point
	Seg []*Line
}
