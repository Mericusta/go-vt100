package core

type Drawable interface {
	Draw(uint, uint, Size)
	Width() uint
	Height() uint
}

type Object struct {
	c Coordinate // absolute coordinate of terminal origin
	d Drawable
}

func NewObject(c Coordinate, d Drawable) Object {
	return Object{c, d}
}

func (o Object) X() uint { return o.c.X }

func (o Object) Y() uint { return o.c.Y }

func (o Object) Draw(s Size) {
	o.d.Draw(o.c.X, o.c.Y, s)
}
