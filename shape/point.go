package shape

import (
	"github.com/Mericusta/go-vt100/core"
	"golang.org/x/text/width"
)

// Point is the unit shape, that is just one ANSI code
// Point height is fixed at 1
// Point width is depends on the width of its ANSI characters
type Point struct {
	ShapeContext
	r rune
}

func NewPoint(r rune) Point {
	p := Point{r: r}
	p.BasicContext = core.NewBasicContext(core.Size{
		Width:  p.Width(),
		Height: p.Height(),
	})
	return p
}

func (p Point) Draw(ctx core.RenderContext, c core.Coordinate) {
	if c.X < 0 {
		return
	}
	if c.Y < 0 {
		return
	}
	p.SetCoordinate(c)
	coincidenceCtx, has := p.CoincidenceCheck(ctx)
	if !has {
		return
	}
	core.MoveCursorToAndPrint(uint(coincidenceCtx.Coordinate().X), uint(coincidenceCtx.Coordinate().Y), string(p.r))
}

func (p Point) Width() uint {
	property := width.LookupRune(p.r)
	switch property.Kind() {
	case width.EastAsianWide, width.EastAsianFullwidth, width.Neutral:
		return 2
	case width.EastAsianNarrow, width.EastAsianHalfwidth, width.EastAsianAmbiguous:
		return 1
	default:
		return 0
	}
}

func (p Point) Height() uint {
	return 1
}
