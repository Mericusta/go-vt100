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
	withBoundry    bool
}

func NewCanvas(s core.Size, withBoundry bool) Canvas {
	c := Canvas{withBoundry: withBoundry}
	if withBoundry {
		c.LeftTop = shape.NewPoint(border.TL())
		c.RightTop = shape.NewPoint(border.TR())
		c.LeftBottom = shape.NewPoint(border.BL())
		c.RightBottom = shape.NewPoint(border.BR())
	}
	c.Resize(s)
	return c
}

// Resize calculate and change new canvas size, include border and context
// @param s new canvas size info
func (c *Canvas) Resize(s core.Size) {
	if c.withBoundry {
		c.HorizontalLine = shape.NewLine(
			shape.NewPoint(border.HL()),
			s.Width, core.Horizontal,
		)
		c.VerticalLine = shape.NewLine(
			shape.NewPoint(border.VL()),
			s.Height, core.Vertical,
		)
		c.BasicContext = core.NewBasicContext(core.Size{
			Width:  s.Width + c.VerticalLine.Width()*2,
			Height: s.Height + c.HorizontalLine.Height()*2,
		})
	} else {
		c.BasicContext = core.NewBasicContext(s)
	}
}

// AppendObjects append objects in canvas
func (c *Canvas) AppendObjects(o ...core.Object) {
	c.objects = append(c.objects, o...)
}

// ClearObjects pop all objects in canvas
func (c *Canvas) ClearObjects() {
	c.objects = nil
}

func (c *Canvas) Draw(ctx core.RenderContext, coordinate core.Coordinate) {
	c.SetCoordinate(coordinate)
	coincidenceCtx, has := ctx.CoincidenceCheck(c)
	if !has {
		return
	}

	// objects
	for _, o := range c.objects {
		o.D.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X + o.C.X + int(c.VerticalLine.Width()), Y: coordinate.Y + o.C.Y + int(c.HorizontalLine.Height())})
	}

	// border
	if c.withBoundry {
		c.LeftTop.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X, Y: coordinate.Y})
		c.HorizontalLine.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X + int(c.LeftTop.Width()), Y: coordinate.Y})
		c.RightTop.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X + int(c.LeftTop.Width()) + int(c.HorizontalLine.Width()), Y: coordinate.Y})
		c.VerticalLine.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X, Y: coordinate.Y + int(c.LeftTop.Height())})
		c.VerticalLine.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X + int(c.LeftTop.Width()) + int(c.HorizontalLine.Width()), Y: coordinate.Y + int(c.LeftTop.Height())})
		c.LeftBottom.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X, Y: coordinate.Y + int(c.LeftTop.Height()) + int(c.VerticalLine.Height())})
		c.HorizontalLine.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X + int(c.LeftBottom.Width()), Y: coordinate.Y + int(c.LeftTop.Height()) + int(c.VerticalLine.Height())})
		c.RightBottom.Draw(coincidenceCtx, core.Coordinate{X: coordinate.X + int(c.LeftTop.Width()) + int(c.HorizontalLine.Width()), Y: coordinate.Y + int(c.LeftTop.Height()) + int(c.VerticalLine.Height())})
	}
}

func (c Canvas) Size() core.Size {
	if c.withBoundry {
		return core.Size{
			Width:  c.Width() - c.VerticalLine.Width()*2,
			Height: c.Height() - c.HorizontalLine.Height()*2,
		}
	} else {
		return core.Size{
			Width:  c.Width(),
			Height: c.Height(),
		}
	}
}
