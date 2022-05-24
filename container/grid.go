package container

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
)

// Grids is a collection of distanced lines
// The area enclosed by the lines is called the Grid
// Grid area can hold another shape instead of Grids
type Grids struct {
	ContainerContext
	col           uint
	row           uint
	objectMap     map[uint]map[uint]core.Drawable // y : x : object
	maxObjectSize core.Size
}

func NewGrids(content map[uint]map[uint]core.Drawable) Grids {
	g := Grids{
		col:       1,
		row:       1,
		objectMap: content,
		maxObjectSize: core.Size{
			Width:  1,
			Height: 1,
		},
	}
	for y, xMap := range g.objectMap {
		if g.row < y {
			g.row = y
		}
		for x, o := range xMap {
			if g.col < x {
				g.col = x
			}
			if g.maxObjectSize.Width < o.Width() {
				g.maxObjectSize.Width = o.Width()
			}
			if g.maxObjectSize.Height < o.Height() {
				g.maxObjectSize.Height = o.Height()
			}
		}
	}
	g.ContainerContext.s.Height = g.row*(g.maxObjectSize.Height+border.TabWidth()) + border.TabWidth()
	g.ContainerContext.s.Width = g.col*(g.maxObjectSize.Width+border.TabWidth()) + border.TabWidth()
	return g
}

func (g *Grids) SetObjects(oMap map[uint]map[uint]core.Drawable) {
	for _y, xMap := range oMap {
		for _x, d := range xMap {
			if _, has := g.objectMap[_y]; !has {
				g.objectMap[_y] = make(map[uint]core.Drawable)
			}
			g.objectMap[_y][_x] = d
		}
	}
}

func (g Grids) Draw(ctx core.RenderContext, coordinate core.Coordinate) {
	g.ContainerContext.c = coordinate

	// border

	// objects
}

func (g Grids) Clear() {

}
