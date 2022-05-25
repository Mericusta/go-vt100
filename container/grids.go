package container

import (
	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

// Grids is a collection of distanced lines
// The area enclosed by the lines is canvas
// Canvas hold another shape instead of Canvas and Grids
// Grids auto expanded itself to max column and row
// Grid auto adapts itself to max canvas size
type Grids struct {
	ContainerContext
	col           uint
	row           uint
	objects       map[uint]map[uint]core.Object // row : col : object
	maxCanvasSize core.Size
}

func NewGrids(canvas map[uint]map[uint]Canvas) Grids {
	g := Grids{
		col:     1,
		row:     1,
		objects: make(map[uint]map[uint]core.Object),
		maxCanvasSize: core.Size{
			Width:  1,
			Height: 1,
		},
	}
	for row, colMap := range canvas {
		if g.row < row {
			g.row = row
		}
		if g.objects[row] == nil {
			g.objects[row] = make(map[uint]core.Object)
		}
		for col, c := range colMap {
			if g.col < col {
				g.col = col
			}
			if g.maxCanvasSize.Width < c.Width() {
				g.maxCanvasSize.Width = c.Width()
			}
			if g.maxCanvasSize.Height < c.Height() {
				g.maxCanvasSize.Height = c.Height()
			}
		}
	}
	g.resize()
	g.adjustCanvasBorder(canvas)
	return g
}

func (g *Grids) SetCanvas(canvas map[uint]map[uint]Canvas) {
	for _y, xMap := range canvas {
		for _x, c := range xMap {
			if _, has := g.objects[_y]; !has {
				return
			}
			if _, has := g.objects[_y][_x]; !has {
				return
			}
			if g.maxCanvasSize.Width < c.Width() {
				g.maxCanvasSize.Width = c.Width()
			}
			if g.maxCanvasSize.Height < c.Height() {
				g.maxCanvasSize.Height = c.Height()
			}
		}
	}
	g.resize()
	g.adjustCanvasBorder(canvas)
}

func (g *Grids) Draw(ctx core.RenderContext, coordinate core.Coordinate) {
	g.ContainerContext.c = coordinate
	for _row := uint(1); _row <= g.row; _row++ {
		for _col := uint(1); _col <= g.col; _col++ {
			o := g.objects[_row][_col]
			o.D.Draw(ctx, core.Coordinate{
				X: coordinate.X + o.C.X,
				Y: coordinate.Y + o.C.Y,
			})
		}
	}
}

func (g *Grids) Clear() {
	for _, colMap := range g.objects {
		for _, o := range colMap {
			o.D.(*Canvas).Clear()
		}
	}
}

func (g *Grids) resize() {
	g.ContainerContext.s.Height = g.row*(g.maxCanvasSize.Height+border.TabWidth()) + border.TabWidth()
	g.ContainerContext.s.Width = g.col*(g.maxCanvasSize.Width+border.TabWidth()) + border.TabWidth()
}

func (g *Grids) adjustCanvasBorder(canvas map[uint]map[uint]Canvas) {
	emptyCanvas := NewCanvas(g.maxCanvasSize)
	for _row := uint(1); _row <= g.row; _row++ {
		for _col := uint(1); _col <= g.col; _col++ {
			var drawCanvas core.Drawable
			if colMap, hasRow := g.objects[_row]; hasRow {
				if o, hasCol := colMap[_col]; hasCol {
					drawCanvas = o.D
				}
			}
			if colMap, hasRow := canvas[_row]; hasRow {
				if c, hasCol := colMap[_col]; hasCol {
					drawCanvas = &c
				}
			}
			if drawCanvas == nil {
				drawCanvas = &emptyCanvas
			}
			switch {
			case _row == 1 && _row != g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(border.TT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(border.LT())
				case _col != 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(border.TT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(border.TT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(border.CT())
				case _col != 1 && _col == g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(border.TT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(border.RT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(border.CT())
				}
			case _row != 1 && _row != g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(border.LT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(border.LT())
				case _col != 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(border.CT())
				case _col != 1 && _col == g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(border.RT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(border.RT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(border.CT())
				}
			case _row != 1 && _row == g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(border.LT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(border.BT())
				case _col != 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(border.BT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(border.BT())
				case _col != 1 && _col == g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(border.CT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(border.RT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(border.BT())
				}
			}
			drawCanvas.(*Canvas).resize(g.maxCanvasSize)
			g.objects[_row][_col] = core.NewObject(
				core.Coordinate{
					X: int((_col-1)*(border.TabWidth()+drawCanvas.Width()) + border.TabWidth()),
					Y: int((_row-1)*(border.TabWidth()+drawCanvas.Height()) + border.TabHeight()),
				},
				drawCanvas,
			)
		}
	}
}
