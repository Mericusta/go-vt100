package object

import (
	"github.com/Mericusta/go-vt100/coordinate"
	"github.com/Mericusta/go-vt100/size"
)

type Drawable interface {
	Draw(int, int, size.Size)
}

type Object struct {
	c coordinate.Coordinate
	d Drawable
}

func NewObject(x, y int, d Drawable) Object {
	return Object{coordinate.Coordinate{X: x, Y: y}, d}
}

func (o Object) X() int { return o.c.X }

func (o Object) Y() int { return o.c.Y }

func (o Object) Draw(s size.Size) {
	o.d.Draw(o.c.X, o.c.Y, s)
}
