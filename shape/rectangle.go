package shape

import (
	"github.com/Mericusta/go-vt100/core"
)

// Rectangle is a collection of lines
type Rectangle struct {
	line  Line
	count uint
}

func NewRectangle(l Line, c uint) Rectangle {
	return Rectangle{
		line:  l,
		count: c,
	}
}

func (r Rectangle) Draw(x, y uint, s core.Size) {
	if x > s.Width || y > s.Height {
		return
	}
	switch r.line.direction {
	case core.Horizontal:
		for _y := y; _y < y+r.count*r.line.Height(); _y += r.line.Height() {
			r.line.Draw(x, _y, s)
		}
	case core.Vertical:
		for _x := x; _x < x+r.count*r.line.Width(); _x += r.line.Width() {
			r.line.Draw(_x, y, s)
		}
	}
}

func (l Rectangle) Width() uint {
	if l.line.direction == core.Horizontal {
		return l.line.Width()
	}
	return l.count * l.line.Width()
}

func (l Rectangle) Height() uint {
	if l.line.direction == core.Vertical {
		return l.line.Height()
	}
	return l.count * l.line.Height()
}
