package shape

import "github.com/Mericusta/go-vt100/core"

// ShapeContext support size and coordinate to objects while drawing
type ShapeContext struct {
	s core.Size
	c core.Coordinate
}

func NewShapeContext(s core.Size, c core.Coordinate) ShapeContext {
	return ShapeContext{s: s, c: c}
}

func (c ShapeContext) Width() uint {
	return c.s.Width
}

func (c ShapeContext) Height() uint {
	return c.s.Height
}

func (c ShapeContext) Coordinate() core.Coordinate {
	return c.c
}
