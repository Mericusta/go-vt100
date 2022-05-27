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
