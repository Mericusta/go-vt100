package container

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

// Canvas is a rectangular area surrounded by standard tab borders
// Canvas area can hold some shapes or other containers instead of Canvas
type Canvas struct {
	core.BasicContext
	objects        []core.Object
	LeftTop        shape.Point
	RightTop       shape.Point
	LeftBottom     shape.Point
	RightBottom    shape.Point
	HorizontalLine shape.Line
	VerticalLine   shape.Line
}

func NewCanvas(s core.Size) Canvas {
	c := Canvas{
		LeftTop:     shape.NewPoint(border.TL()),
		RightTop:    shape.NewPoint(border.TR()),
		LeftBottom:  shape.NewPoint(border.BL()),
		RightBottom: shape.NewPoint(border.BR()),
	}
	c.Resize(s)
	return c
}

func (c *Canvas) AppendObjects(o ...core.Object) {
	c.objects = append(c.objects, o...)
}

func (c *Canvas) Draw(ctx core.RenderContext, coordinate core.Coordinate) {
	c.SetCoordinate(coordinate)
	coincidenceCtx, has := ctx.CoincidenceCheck(c)
	if !has {
		return
	}

	// objects
	for _, o := range c.objects {
		o.D.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X + o.C.X, Y: coordinate.Y + o.C.Y})
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
	r := shape.NewRectangle(
		shape.NewLine(
			shape.NewPoint(border.Space()),
			c.Width()+2*border.TabWidth(),
			core.Horizontal,
		),
		c.Height()+2*border.TabHeight(),
	)
	clearCtx := shape.NewShapeContext(core.Size{
		Width:  r.Width() + border.TabWidth()*2,
		Height: r.Height() + border.TabHeight()*2,
	})
	clearCtx.SetCoordinate(core.Coordinate{X: c.Coordinate().X - 1, Y: c.Coordinate().Y - 1})
	r.Draw(clearCtx, core.Coordinate{X: c.Coordinate().X - 1, Y: c.Coordinate().Y - 1})
	c.objects = nil
}

func (c *Canvas) Resize(s core.Size) {
	c.HorizontalLine = shape.NewLine(
		shape.NewPoint(border.HL()),
		s.Width, core.Horizontal,
	)
	c.VerticalLine = shape.NewLine(
		shape.NewPoint(border.VL()),
		s.Height, core.Vertical,
	)
	c.BasicContext = core.NewBasicContext(core.Size{
		Width:  s.Width,
		Height: s.Height,
	})
}
