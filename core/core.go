package core

type Unit interface {
	Width() uint
	Height() uint
}

type RenderContext interface {
	Unit
	// Coordinate
	// @return absolute coordinate
	Coordinate() Coordinate
}

type Drawable interface {
	Unit
	// Draw
	// @Coordinate absolute coordinate
	Draw(RenderContext, Coordinate)
}

// Object hold drawable and its relative coordinate of parent
type Object struct {
	C Coordinate // relative coordinate
	D Drawable
}

func NewObject(c Coordinate, d Drawable) Object {
	return Object{c, d}
}

// func (o Object) X() int { return o.c.X }

// func (o Object) Y() int { return o.c.Y }

// func (o Object) Draw() {
// 	o.d.Draw(o.c)
// }
