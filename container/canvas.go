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
	horizontalLine shape.Line
	verticalLine   shape.Line
}

func NewCanvas(s core.Size) Canvas {
	return Canvas{
		ContainerContext: ContainerContext{s: s},
		LeftTop:          shape.NewPoint(border.TL()),
		RightTop:         shape.NewPoint(border.TR()),
		LeftBottom:       shape.NewPoint(border.BL()),
		RightBottom:      shape.NewPoint(border.BR()),
		horizontalLine: shape.NewLine(
			shape.NewPoint(border.HL()),
			s.Width, core.Horizontal,
		),
		verticalLine: shape.NewLine(
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
	c.horizontalLine.Draw(ctx, core.Coordinate{X: coordinate.X, Y: coordinate.Y - 1})
	c.RightTop.Draw(ctx, core.Coordinate{X: coordinate.X + int(c.Width()), Y: coordinate.Y - 1})
	c.verticalLine.Draw(ctx, core.Coordinate{X: coordinate.X - 1, Y: coordinate.Y})
	c.verticalLine.Draw(ctx, core.Coordinate{X: coordinate.X + int(c.Width()), Y: coordinate.Y})
	c.LeftBottom.Draw(ctx, core.Coordinate{X: coordinate.X - 1, Y: coordinate.Y + int(c.Height())})
	c.horizontalLine.Draw(ctx, core.Coordinate{X: coordinate.X, Y: coordinate.Y + int(c.Height())})
	c.RightBottom.Draw(ctx, core.Coordinate{X: coordinate.X + int(c.Width()), Y: coordinate.Y + int(c.Height())})
}

func (c *Canvas) Clear() {
	c.objects = nil
	r := shape.NewRectangle(shape.NewLine(shape.NewPoint(border.Space()), c.Width()+2, core.Horizontal), c.Height()+2)
	r.Draw(c, core.Coordinate{
		X: c.c.X - 1,
		Y: c.c.Y - 1,
	})
}
