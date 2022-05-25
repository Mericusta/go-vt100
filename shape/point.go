package shape

import (
	"github.com/Mericusta/go-vt100/core"
	"golang.org/x/text/width"
)

// Point is the unit shape, that is just one ANSI code
// Point height is fixed at 1
// Point width is depends on the width of its ANSI characters
type Point struct {
	r rune
}

func NewPoint(r rune) Point {
	return Point{r}
}

func (p Point) Draw(ctx core.RenderContext, c core.Coordinate) {
	startAbsX := c.X
	endAbsX := startAbsX + int(p.Width())
	if startAbsX < ctx.Coordinate().X || endAbsX > ctx.Coordinate().X+int(ctx.Width()) {
		// outer left || outer right
		return
	}
	startAbsY := c.Y
	endAbsY := startAbsY + int(p.Height())
	if startAbsY < ctx.Coordinate().Y || endAbsY > ctx.Coordinate().Y+int(ctx.Height()) {
		// outer top || outer bottom
		return
	}
	if startAbsX < 0 {
		startAbsX = 0
	}
	if startAbsY < 0 {
		startAbsY = 0
	}
	core.MoveCursorToAndPrint(uint(startAbsX), uint(startAbsY), string(p.r))
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
