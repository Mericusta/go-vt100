package shape

import (
	"github.com/Mericusta/go-vt100/core"
	"golang.org/x/text/width"
)

// Point is the unit shape, that is just one ANSI code
// Point height is fixed at 1
// Point width is its length in []byte
type Point struct {
	r rune
}

func NewPoint(r rune) Point {
	return Point{r}
}

func (p Point) Draw(x, y uint, s core.Size) {
	if x > s.Width || y > s.Height {
		return
	}
	core.MoveCursorToAndPrint(x, y, string(p.r))
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
