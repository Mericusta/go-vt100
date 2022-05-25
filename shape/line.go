package shape

import "github.com/Mericusta/go-vt100/core"

// Line is a collection of points in the horizontal or vertical direction
// the default origin is {0,0}
// line origin is relative coordinate
type Line struct {
	ShapeContext
	point     Point
	length    uint
	direction core.Direction
}

func NewLine(p Point, l uint, d core.Direction) Line {
	return Line{
		point:     p,
		length:    l,
		direction: d,
	}
}

func (l Line) Draw(ctx core.RenderContext, c core.Coordinate) {
	l.ShapeContext.c = c
	switch l.direction {
	case core.Horizontal:
		for _x := 0; _x < int(l.length*l.point.Width()); _x += int(l.point.Width()) {
			l.point.Draw(ctx, core.Coordinate{X: c.X + _x, Y: c.Y})
		}
	case core.Vertical:
		for _y := 0; _y < int(l.length*l.point.Height()); _y += int(l.point.Height()) {
			l.point.Draw(ctx, core.Coordinate{X: c.X, Y: c.Y + _y})
		}
	}
}

func (l Line) Width() uint {
	if l.direction == core.Vertical {
		return l.point.Width()
	}
	return l.length * l.point.Width()
}

func (l Line) Height() uint {
	if l.direction == core.Horizontal {
		return l.point.Height()
	}
	return l.length * l.point.Height()
}
