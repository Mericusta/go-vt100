package core

type Drawable interface {
	Draw(int, int, Size)
}

type Object struct {
	c Coordinate // relative coordinate
	d Drawable
}

func NewObject(c Coordinate, d Drawable) Object {
	return Object{c, d}
}

func (o Object) X() int { return o.c.X }

func (o Object) Y() int { return o.c.Y }

func (o Object) Draw(s Size) {
	o.d.Draw(o.c.X, o.c.Y, s)
}
