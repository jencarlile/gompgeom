package gompgeom

import (
	. "launchpad.net/gocheck"
	"sort"
	"testing"
)

// Hook up gocheck into the gotest runner.
func Test(t *testing.T) { TestingT(t) }

type PointSuite struct{}

var _ = Suite(&PointSuite{})

func (s *PointSuite) TestToRight(c *C) {
	p1 := &Point{0, 0}
	p2 := &Point{10, 0}

	pLeft := &Point{12, 3}
	pRight := &Point{11, -1}
	pInline := &Point{12, 0}

	c.Check(ToRight(p1, p2, pLeft), Equals, false)
	c.Check(ToRight(p1, p2, pRight), Equals, true)
	c.Check(ToRight(p1, p2, pInline), Equals, false)
}

func (s *PointSuite) TestSortByX(c *C) {
	pts := []*Point{
		{10, 5},
		{2, 6},
		{3, 4},
		{12, 3},
		{7, 7},
	}

	sort.Sort(ByX{pts})
	expected := []*Point{
		{2, 6},
		{3, 4},
		{7, 7},
		{10, 5},
		{12, 3}}

	c.Check(pts, DeepEquals, expected)
}

func (s *PointSuite) TestConvexHull(c *C) {
	pts := []*Point{
		{8, 3},
		{5, 2},
		{4, 5},
		{2, 1},
		{1, 3},
		{0, 0},
		{3, -3},
	}
	expected := Polygon{
		{0, 0},
		{1, 3},
		{4, 5},
		{8, 3},
		{3, -3},
	}
	c.Check(ConvexHull(pts), DeepEquals, expected)

	pts = []*Point{
		{0, 0},
		{5, 5},
		{10, 0},
		{4, -4},
	}
	expected = Polygon{
		{0, 0},
		{5, 5},
		{10, 0},
		{4, -4},
	}
	c.Check(ConvexHull(pts), DeepEquals, expected)

	pts = []*Point{
		{0, 0},
		{5, 5},
		{10, 0},
		{5, -4},
	}
	expected = Polygon{
		{0, 0},
		{5, 5},
		{10, 0},
		{5, -4},
	}
	c.Check(ConvexHull(pts), DeepEquals, expected)

	pts = []*Point{
		{0, 0},
		{5, 5},
		{8, 2},
		{10, 0},
	}
	expected = Polygon{
		{0, 0},
		{5, 5},
		{10, 0},
	}
	c.Check(ConvexHull(pts), DeepEquals, expected)

	pts = []*Point{
		{0, 0},
		{10, 0},
	}
	expected = Polygon{}
	c.Check(ConvexHull(pts), DeepEquals, expected)
}
