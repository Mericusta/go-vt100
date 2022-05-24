package core

type RenderContext interface {
	Width() uint
	Height() uint
	// Coordinate
	// @return absolute coordinate
	Coordinate() Coordinate
}

type Drawable interface {
	// Width() uint
	// Height() uint
	// Draw
	// @Coordinate absolute coordinate
	Draw(RenderContext, Coordinate)
}

// type RenderTree struct {
// 	nodes []RenderTree
// }

// func (t RenderTree) Append() {
// 	// t.N.Draw()
// }

// func (t RenderTree) Render() {
// 	// t.N.Draw()
// }

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
