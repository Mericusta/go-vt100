package shape

import (
	"math"

	"github.com/Mericusta/go-vt100/core"
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
	if p.r > math.MaxInt8 {
		return 2
	}
	return 1
}

func (p Point) Height() uint {
	return 1
}
