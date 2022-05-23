package shape

import "github.com/Mericusta/go-vt100/core"

type Grid struct {
	// logic data
	col     int
	row     int
	content map[int]map[int][]byte // y : x : content
	// graphic data
	size            core.Size
	maxContentWidth int
}

func NewGrid(content map[int]map[int][]byte) *Grid {
	if len(content) == 0 {
		return nil
	}
	g := &Grid{
		col:             1,
		row:             1,
		content:         content,
		maxContentWidth: 1,
	}
	for y, xMap := range g.content {
		if g.row < y {
			g.row = y
		}
		for x, c := range xMap {
			if g.col < x {
				g.col = x
			}
			if g.maxContentWidth < len(c) {
				g.maxContentWidth = len(c)
			}
		}
	}
	g.size.Height = g.row*(1+core.TabWidth()) + core.TabWidth()
	g.size.Width = g.col*(g.maxContentWidth+core.TabWidth()) + core.TabWidth()
	return g
}

func (g Grid) Draw(x, y int) {

}
