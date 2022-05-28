package container

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

// About Object: Container contains some shapes, the shapes relative coordinate is defined by user, so they need Object.
// About Context: Container can be contained other container, so they must know its container Size and Coordinate.

// ContainerContext support container size and coordinate to objects while drawing
type ContainerContext struct {
	core.BasicContext
}

func NewContainerContext(s core.Size) ContainerContext {
	return ContainerContext{BasicContext: core.NewBasicContext(s)}
}

func (ctx *ContainerContext) Clear() {
	r := shape.NewRectangle(
		shape.NewLine(
			shape.NewPoint(border.Space()),
			ctx.Width(),
			core.Horizontal,
		),
		ctx.Height(),
	)
	clearCtx := shape.NewShapeContext(core.Size{
		Width:  r.Width(),
		Height: r.Height(),
	})
	clearCtx.SetCoordinate(core.Coordinate{X: ctx.Coordinate().X, Y: ctx.Coordinate().Y})
	r.Draw(clearCtx, core.Coordinate{X: ctx.Coordinate().X, Y: ctx.Coordinate().Y})
}
