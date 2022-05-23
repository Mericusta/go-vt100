package shape

import "github.com/Mericusta/go-vt100/core"

// Line is a collection of points in the horizontal or vertical direction
// the default origin is {0,0}
// line origin is relative coordinate
type Line struct {
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

func (l Line) Draw(x, y uint, s core.Size) {
	if x > s.Width || y > s.Height {
		return
	}
	switch l.direction {
	case core.Horizontal:
		for _x := x; _x < x+l.length*l.point.Width(); _x += l.point.Width() {
			l.point.Draw(_x, y, s)
		}
	case core.Vertical:
		for _y := y; _y < y+l.length*l.point.Height(); _y += l.point.Height() {
			l.point.Draw(x, _y, s)
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
