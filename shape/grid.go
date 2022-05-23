package shape

import "github.com/Mericusta/go-vt100/core"

// Grids is a collection of distanced lines
// The area enclosed by the lines is called the Grid
// Grid area can hold another shape instead of Grids
type Grids struct {
	col             uint
	row             uint
	line            Line
	contents        map[uint]map[uint]Shape // y : x : content
	size            core.Size
	maxContentWidth uint
}

func NewGrid(content map[uint]map[uint]Shape) Grids {
	g := Grids{
		col:             1,
		row:             1,
		contents:        content,
		maxContentWidth: 1,
	}
	for y, xMap := range g.contents {
		if g.row < y {
			g.row = y
		}
		for x, c := range xMap {
			if g.col < x {
				g.col = x
			}
			if g.maxContentWidth < c.Width() {
				g.maxContentWidth = c.Width()
			}
		}
	}
	g.size.Height = g.row*(1+core.TabWidth()) + core.TabWidth()
	g.size.Width = g.col*(g.maxContentWidth+core.TabWidth()) + core.TabWidth()
	return g
}

func (g Grids) Draw(x, y int) {

}

func (g Grids) Width() uint {
	return 0
}

func (g Grids) Height() uint {
	return 0
}
