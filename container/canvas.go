package container

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

// Canvas is a rectangular area surrounded by standard tab borders
// Canvas area can hold some shapes or other containers instead of Canvas
type Canvas struct {
	ContainerContext
	objects        []core.Object
	LeftTop        shape.Point
	RightTop       shape.Point
	LeftBottom     shape.Point
	RightBottom    shape.Point
	HorizontalLine shape.Line
	VerticalLine   shape.Line
}

func NewCanvas(s core.Size) Canvas {
	return Canvas{
		ContainerContext: ContainerContext{s: s},
		LeftTop:          shape.NewPoint(border.TL()),
		RightTop:         shape.NewPoint(border.TR()),
		LeftBottom:       shape.NewPoint(border.BL()),
		RightBottom:      shape.NewPoint(border.BR()),
		HorizontalLine: shape.NewLine(
			shape.NewPoint(border.HL()),
			s.Width, core.Horizontal,
		),
		VerticalLine: shape.NewLine(
			shape.NewPoint(border.VL()),
			s.Height, core.Vertical,
		),
	}
}

func (c *Canvas) AppendObjects(o ...core.Object) {
	c.objects = append(c.objects, o...)
}

func (c *Canvas) Draw(ctx core.RenderContext, coordinate core.Coordinate) {
	c.ContainerContext.c = coordinate

	// objects
	for _, o := range c.objects {
		o.D.Draw(c, core.Coordinate{X: coordinate.X + o.C.X, Y: coordinate.Y + o.C.Y})
	}

	// border
	c.LeftTop.Draw(ctx, core.Coordinate{X: coordinate.X - 1, Y: coordinate.Y - 1})
	c.HorizontalLine.Draw(ctx, core.Coordinate{X: coordinate.X, Y: coordinate.Y - 1})
	c.RightTop.Draw(ctx, core.Coordinate{X: coordinate.X + int(c.Width()), Y: coordinate.Y - 1})
	c.VerticalLine.Draw(ctx, core.Coordinate{X: coordinate.X - 1, Y: coordinate.Y})
	c.VerticalLine.Draw(ctx, core.Coordinate{X: coordinate.X + int(c.Width()), Y: coordinate.Y})
	c.LeftBottom.Draw(ctx, core.Coordinate{X: coordinate.X - 1, Y: coordinate.Y + int(c.Height())})
	c.HorizontalLine.Draw(ctx, core.Coordinate{X: coordinate.X, Y: coordinate.Y + int(c.Height())})
	c.RightBottom.Draw(ctx, core.Coordinate{X: coordinate.X + int(c.Width()), Y: coordinate.Y + int(c.Height())})
}

func (c *Canvas) Clear() {
	// c.objects = []core.Object{core.NewObject(
	// 	core.Coordinate{X: 0, Y: 0},
	// 	shape.NewRectangle(
	// 		shape.NewLine(
	// 			shape.NewPoint(border.Space()),
	// 			c.Width(),
	// 			core.Horizontal,
	// 		),
	// 		c.Height(),
	// 	),
	// )}
	// c.Draw(c, c.Coordinate())
	// c.objects = nil

	r := shape.NewRectangle(
		shape.NewLine(
			shape.NewPoint(border.Space()),
			c.Width()+2*border.TabWidth(),
			core.Horizontal,
		),
		c.Height()+2*border.TabHeight(),
	)
	rCoordinate := core.Coordinate{X: c.c.X - 1, Y: c.c.Y - 1}
	clearCtx := shape.NewShapeContext(core.Size{Width: r.Width(), Height: r.Height()}, rCoordinate)
	r.Draw(clearCtx, rCoordinate)
	c.objects = nil
}

func (c *Canvas) resize(s core.Size) {
	c.s = s
	c.HorizontalLine = shape.NewLine(
		shape.NewPoint(border.HL()),
		s.Width, core.Horizontal,
	)
	c.VerticalLine = shape.NewLine(
		shape.NewPoint(border.VL()),
		s.Height, core.Vertical,
	)
}
