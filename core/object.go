package core

type Drawable interface {
	Draw(int, int, Size)
}

type Object struct {
	c Coordinate
	d Drawable
}

func NewObject(x, y int, d Drawable) Object {
	return Object{Coordinate{X: x, Y: y}, d}
}

func (o Object) X() int { return o.c.X }

func (o Object) Y() int { return o.c.Y }

func (o Object) Draw(s Size) {
	o.d.Draw(o.c.X, o.c.Y, s)
}
