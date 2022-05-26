package shape

import "github.com/Mericusta/go-vt100/core"

// About Object: Ths Shape is basic rendering unit, so they don't need Object.
// About Context: The Shape is basic rendering unit, so they must know its container Size and Coordinate.

// ShapeContext support shape container size and coordinate to objects while drawing
type ShapeContext struct {
	core.BasicContext
}

func NewShapeContext(s core.Size) ShapeContext {
	return ShapeContext{BasicContext: core.NewBasicContext(s)}
}

func (ctx *ShapeContext) SetCoordinate(c core.Coordinate) {
	x, y := c.X, c.Y
	if x < 0 {
		x = 0
	}
	if y < 0 {
		y = 0
	}
	ctx.BasicContext.SetCoordinate(core.Coordinate{X: x, Y: y})
}
