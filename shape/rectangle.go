package shape

import "github.com/Mericusta/go-vt100/core"

// Rectangle is a collection of lines
type Rectangle struct {
	ShapeContext
	line  Line
	count uint
}

func NewRectangle(l Line, c uint) Rectangle {
	r := Rectangle{
		line:  l,
		count: c,
	}
	r.BasicContext = core.NewBasicContext(core.Size{
		Width:  r.Width(),
		Height: r.Height(),
	})
	return r
}

func (r Rectangle) Draw(ctx core.RenderContext, c core.Coordinate) {
	r.SetCoordinate(c)
	coincidenceCtx, has := r.CoincidenceCheck(ctx)
	if !has {
		return
	}
	switch r.line.direction {
	case core.Horizontal:
		for _y := 0; _y < int(r.count*r.line.Height()); _y += int(r.line.Height()) {
			r.line.Draw(coincidenceCtx, core.Coordinate{X: c.X, Y: c.Y + _y})
		}
	case core.Vertical:
		for _x := 0; _x < int(r.count*r.line.Width()); _x += int(r.line.Width()) {
			r.line.Draw(coincidenceCtx, core.Coordinate{X: c.X + _x, Y: c.Y})
		}
	}
}

func (l Rectangle) Width() uint {
	if l.line.direction == core.Horizontal {
		return l.line.Width()
	}
	return l.count * l.line.Width()
}

func (l Rectangle) Height() uint {
	if l.line.direction == core.Vertical {
		return l.line.Height()
	}
	return l.count * l.line.Height()
}
