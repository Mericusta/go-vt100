package container

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

// Grids is a collection of distanced lines
// The area enclosed by the lines is canvas
// Canvas hold another shape instead of Canvas and Grids
type Grids struct {
	ContainerContext
	col           uint
	row           uint
	canvasMap     map[uint]map[uint]Canvas // row : col : canvas
	maxCanvasSize core.Size
}

func NewGrids(canvas map[uint]map[uint]Canvas) Grids {
	g := Grids{
		col:       1,
		row:       1,
		canvasMap: canvas,
		maxCanvasSize: core.Size{
			Width:  1,
			Height: 1,
		},
	}
	for y, xMap := range g.canvasMap {
		if g.row < y {
			g.row = y
		}
		for x, o := range xMap {
			if g.col < x {
				g.col = x
			}
			if g.maxCanvasSize.Width < o.Width() {
				g.maxCanvasSize.Width = o.Width()
			}
			if g.maxCanvasSize.Height < o.Height() {
				g.maxCanvasSize.Height = o.Height()
			}
		}
	}
	g.ContainerContext.s.Height = g.row*(g.maxCanvasSize.Height+border.TabWidth()) + border.TabWidth()
	g.ContainerContext.s.Width = g.col*(g.maxCanvasSize.Width+border.TabWidth()) + border.TabWidth()
	for _row := uint(1); _row <= g.row; _row++ {
		for _col := uint(1); _col <= g.col; _col++ {
			drawCanvas := NewCanvas(g.maxCanvasSize)
			if colMap, hasRow := g.canvasMap[_row]; hasRow {
				if c, hasCol := colMap[_col]; hasCol {
					drawCanvas = c
				}
			} else {
				g.canvasMap[_row] = make(map[uint]Canvas)
			}
			switch {
			case _row == 1 && _row != g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.RightTop = shape.NewPoint(border.TT())
					drawCanvas.RightBottom = shape.NewPoint(border.CT())
					drawCanvas.LeftBottom = shape.NewPoint(border.LT())
				case _col != 1 && _col != g.col:
					drawCanvas.LeftTop = shape.NewPoint(border.TT())
					drawCanvas.RightTop = shape.NewPoint(border.TT())
					drawCanvas.RightBottom = shape.NewPoint(border.CT())
					drawCanvas.LeftBottom = shape.NewPoint(border.CT())
				case _col != 1 && _col == g.col:
					drawCanvas.LeftTop = shape.NewPoint(border.TT())
					drawCanvas.RightBottom = shape.NewPoint(border.RT())
					drawCanvas.LeftBottom = shape.NewPoint(border.CT())
				}
			case _row != 1 && _row != g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.LeftTop = shape.NewPoint(border.LT())
					drawCanvas.RightTop = shape.NewPoint(border.CT())
					drawCanvas.RightBottom = shape.NewPoint(border.CT())
					drawCanvas.LeftBottom = shape.NewPoint(border.LT())
				case _col != 1 && _col != g.col:
					drawCanvas.LeftTop = shape.NewPoint(border.CT())
					drawCanvas.RightTop = shape.NewPoint(border.CT())
					drawCanvas.RightBottom = shape.NewPoint(border.CT())
					drawCanvas.LeftBottom = shape.NewPoint(border.CT())
				case _col != 1 && _col == g.col:
					drawCanvas.LeftTop = shape.NewPoint(border.CT())
					drawCanvas.RightTop = shape.NewPoint(border.RT())
					drawCanvas.RightBottom = shape.NewPoint(border.RT())
					drawCanvas.LeftBottom = shape.NewPoint(border.CT())
				}
			case _row != 1 && _row == g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.LeftTop = shape.NewPoint(border.LT())
					drawCanvas.RightTop = shape.NewPoint(border.CT())
					drawCanvas.RightBottom = shape.NewPoint(border.BT())
				case _col != 1 && _col != g.col:
					drawCanvas.LeftTop = shape.NewPoint(border.CT())
					drawCanvas.RightTop = shape.NewPoint(border.CT())
					drawCanvas.RightBottom = shape.NewPoint(border.BT())
					drawCanvas.LeftBottom = shape.NewPoint(border.BT())
				case _col != 1 && _col == g.col:
					drawCanvas.LeftTop = shape.NewPoint(border.CT())
					drawCanvas.RightTop = shape.NewPoint(border.RT())
					drawCanvas.LeftBottom = shape.NewPoint(border.BT())
				}
			}
			g.canvasMap[_row][_col] = drawCanvas
		}
	}
	return g
}

func (g *Grids) SetObjects(oMap map[uint]map[uint]Canvas) {
	for _y, xMap := range oMap {
		for _x, d := range xMap {
			if _, has := g.canvasMap[_y]; !has {
				g.canvasMap[_y] = make(map[uint]Canvas)
			}
			g.canvasMap[_y][_x] = d
		}
	}
}

func (g Grids) Draw(ctx core.RenderContext, coordinate core.Coordinate) {
	g.ContainerContext.c = coordinate
	for _row := uint(1); _row <= g.row; _row++ {
		for _col := uint(1); _col <= g.col; _col++ {
			c := g.canvasMap[_row][_col]
			// terminal.DebugOutput(func() {
			// 	fmt.Printf("c = %v, %v\n", c.RightTop, border.TR())
			// }, nil)
			c.Draw(ctx, core.Coordinate{
				X: coordinate.X + int((_col-1)*(border.TabWidth()+c.Width())) + int(border.TabWidth()),
				Y: coordinate.Y + int((_row-1)*(border.TabHeight()+c.Height())) + int(border.TabHeight()),
			})
		}
	}

	// objects
}

func (g Grids) Clear() {

}
