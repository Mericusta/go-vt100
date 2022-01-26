package object

import (
	"go-vt100/coordinate"
)

type Drawable interface {
	Draw(int, int)
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

func (o Object) Draw() {
	o.d.Draw(o.c.X, o.c.Y)
}
