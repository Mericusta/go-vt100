package container

import (
	"github.com/Mericusta/go-vt100/character"
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
			if g.maxCanvasSize.Width < (c.Width() - character.TabWidth()*2) {
				g.maxCanvasSize.Width = c.Width() - character.TabWidth()*2
			}
			if g.maxCanvasSize.Height < (c.Height() - character.TabHeight()*2) {
				g.maxCanvasSize.Height = c.Height() - character.TabHeight()*2
			}
		}
	}
	g.resize()
	g.adjustBorder(canvas)
	return g
}

func (g *Grids) resize() {
	g.BasicContext = core.NewBasicContext(core.Size{
		Width:  g.row*(g.maxCanvasSize.Width+character.TabWidth()) + character.TabWidth(),
		Height: g.col*(g.maxCanvasSize.Height+character.TabHeight()) + character.TabHeight(),
	})
}

func (g *Grids) adjustBorder(canvas map[uint]map[uint]Canvas) {
	emptyCanvas := NewCanvas(g.maxCanvasSize, true)
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
			drawCanvas.(*Canvas).Resize(g.maxCanvasSize)
			switch {
			case _row == 1 && _row != g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(character.TT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(character.LT())
				case _col != 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(character.TT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(character.TT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(character.CT())
				case _col != 1 && _col == g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(character.TT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(character.RT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(character.CT())
				}
			case _row != 1 && _row != g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(character.LT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(character.LT())
				case _col != 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(character.CT())
				case _col != 1 && _col == g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(character.RT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(character.RT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(character.CT())
				}
			case _row != 1 && _row == g.row:
				switch {
				case _col == 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(character.LT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(character.BT())
				case _col != 1 && _col != g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).RightBottom = shape.NewPoint(character.BT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(character.BT())
				case _col != 1 && _col == g.col:
					drawCanvas.(*Canvas).LeftTop = shape.NewPoint(character.CT())
					drawCanvas.(*Canvas).RightTop = shape.NewPoint(character.RT())
					drawCanvas.(*Canvas).LeftBottom = shape.NewPoint(character.BT())
				}
			}
			g.objects[_row][_col] = core.NewObject(
				core.Coordinate{
					X: int((_col - 1) * (drawCanvas.Width() - character.TabWidth())),
					Y: int((_row - 1) * (drawCanvas.Height() - character.TabHeight())),
				},
				drawCanvas,
			)
		}
	}
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
			if g.maxCanvasSize.Width < c.Size().Width {
				g.maxCanvasSize.Width = c.Size().Width
			}
			if g.maxCanvasSize.Height < c.Size().Height {
				g.maxCanvasSize.Height = c.Size().Height
			}
		}
	}
	g.resize()
	g.adjustBorder(canvas)
}

func (g *Grids) ClearGrid(gridMap map[uint]map[uint]struct{}) {
	for _y, xMap := range gridMap {
		for _x := range xMap {
			g.objects[_y][_x].D.(*Canvas).ClearObjects()
		}
	}
}

func (g *Grids) ClearAllGrid() {
	for _y, xMap := range g.objects {
		for _x := range xMap {
			g.objects[_y][_x].D.(*Canvas).ClearObjects()
		}
	}
}

func (g *Grids) Draw(ctx core.RenderContext, c core.Coordinate) {
	g.BasicContext.SetCoordinate(c)
	coincidenceCtx, has := ctx.CoincidenceCheck(g)
	if !has {
		return
	}
	for _row := uint(1); _row <= g.row; _row++ {
		for _col := uint(1); _col <= g.col; _col++ {
			o := g.objects[_row][_col]
			o.D.Draw(coincidenceCtx, core.Coordinate{
				X: c.X + o.C.X,
				Y: c.Y + o.C.Y,
			})
		}
	}
}
