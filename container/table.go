package container

import (
	"github.com/Mericusta/go-vt100/character"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

// Table
type Table struct {
	ContainerContext
	objects        map[uint]map[uint]core.Object // row : col : object
	columnMaxWidth map[uint]uint                 // col : width
	rowMaxHeight   map[uint]uint                 // row : height
	rowCount       uint
}

func NewTable(headerSlice []core.Drawable, valueMap map[uint]map[uint]core.Drawable) Table {
	t := Table{
		objects:        make(map[uint]map[uint]core.Object),
		columnMaxWidth: make(map[uint]uint, len(headerSlice)),
		rowMaxHeight:   make(map[uint]uint, 1+len(valueMap)),
	}
	for i, h := range headerSlice {
		if t.columnMaxWidth[uint(i)+1] == 0 || t.columnMaxWidth[uint(i)+1] < h.Width() {
			t.columnMaxWidth[uint(i)+1] = h.Width()
		}
		if t.rowMaxHeight[0] == 0 || t.rowMaxHeight[0] < h.Height() {
			t.rowMaxHeight[0] = h.Height()
		}
	}
	t.rowMaxHeight[1] = 1
	t.rowCount = uint(1) // from 0, including header, at least one value row
	for row, colMap := range valueMap {
		if t.rowCount < row {
			t.rowCount = row
		}
		for col, d := range colMap {
			if t.columnMaxWidth[col] == 0 || t.columnMaxWidth[col] < d.Width() {
				t.columnMaxWidth[col] = d.Width()
			}
			if t.rowMaxHeight[row] == 0 || t.rowMaxHeight[row] < d.Height() {
				t.rowMaxHeight[row] = d.Height()
			}
		}
	}
	for row := uint(0); row <= t.rowCount; row++ {
		t.objects[uint(row)] = make(map[uint]core.Object, len(headerSlice))
		if t.rowMaxHeight[uint(row)] == 0 {
			t.rowMaxHeight[uint(row)] = 1
		}
	}
	t.resize()
	t.adjustBorder(headerSlice, valueMap)
	return t
}

func (t *Table) resize() {
	totalWidth := uint(0)
	totalHeight := uint(0)
	for _, width := range t.columnMaxWidth {
		totalWidth += uint(character.TabWidth()) + width
	}
	for _, height := range t.rowMaxHeight {
		totalHeight += uint(character.TabHeight()) + height
	}
	t.BasicContext = core.NewBasicContext(core.Size{
		Width:  totalWidth + uint(character.TabWidth()),
		Height: totalHeight + uint(character.TabHeight()),
	})
}

func (t *Table) adjustBorder(headerSlice []core.Drawable, valueMap map[uint]map[uint]core.Drawable) {
	colCount := uint(len(headerSlice))
	objY := 0
	for _row := uint(0); _row <= t.rowCount; _row++ {
		objX := 0
		for _col := uint(0); _col < colCount; _col++ {
			drawCanvas := NewCanvas(core.Size{
				Width:  t.columnMaxWidth[uint(_col)+1],
				Height: t.rowMaxHeight[uint(_row)],
			}, true)
			if _row == 0 {
				drawCanvas.AppendObjects(core.NewObject(
					core.Coordinate{},
					headerSlice[_col],
				))
			} else {
				if len(valueMap) > 0 {
					if colMap, has := valueMap[_row]; has && len(colMap) > 0 {
						if d, has := valueMap[_row][_col+1]; has {
							drawCanvas.AppendObjects(core.NewObject(core.Coordinate{}, d))
						}
					}
				}
			}
			switch {
			case _row == 0 && _row == t.rowCount:
				switch {
				case _col == 0 && _col == colCount-1:
				case _col == 0 && _col != colCount-1:
					drawCanvas.RightTop = shape.NewPoint(character.TT())
					drawCanvas.RightBottom = shape.NewPoint(character.BT())
				case _col != 0 && _col != colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.TT())
					drawCanvas.RightTop = shape.NewPoint(character.TT())
					drawCanvas.RightBottom = shape.NewPoint(character.BT())
					drawCanvas.LeftBottom = shape.NewPoint(character.BT())
				case _col != 0 && _col == colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.TT())
					drawCanvas.LeftBottom = shape.NewPoint(character.BT())
				}
			case _row == 0 && _row != t.rowCount:
				switch {
				case _col == 0 && _col != colCount-1:
					drawCanvas.RightTop = shape.NewPoint(character.TT())
					drawCanvas.RightBottom = shape.NewPoint(character.CT())
					drawCanvas.LeftBottom = shape.NewPoint(character.LT())
				case _col != 0 && _col != colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.TT())
					drawCanvas.RightTop = shape.NewPoint(character.TT())
					drawCanvas.RightBottom = shape.NewPoint(character.CT())
					drawCanvas.LeftBottom = shape.NewPoint(character.CT())
				case _col != 0 && _col == colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.TT())
					drawCanvas.RightBottom = shape.NewPoint(character.RT())
					drawCanvas.LeftBottom = shape.NewPoint(character.CT())
				}
			case _row != 0 && _row != t.rowCount:
				switch {
				case _col == 0 && _col != colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.LT())
					drawCanvas.RightTop = shape.NewPoint(character.CT())
					drawCanvas.RightBottom = shape.NewPoint(character.CT())
					drawCanvas.LeftBottom = shape.NewPoint(character.LT())
				case _col != 0 && _col != colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.CT())
					drawCanvas.RightTop = shape.NewPoint(character.CT())
					drawCanvas.RightBottom = shape.NewPoint(character.CT())
					drawCanvas.LeftBottom = shape.NewPoint(character.CT())
				case _col != 0 && _col == colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.CT())
					drawCanvas.RightTop = shape.NewPoint(character.RT())
					drawCanvas.RightBottom = shape.NewPoint(character.RT())
					drawCanvas.LeftBottom = shape.NewPoint(character.CT())
				}
			case _row != 0 && _row == t.rowCount:
				switch {
				case _col == 0 && _col != colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.LT())
					drawCanvas.RightTop = shape.NewPoint(character.CT())
					drawCanvas.RightBottom = shape.NewPoint(character.BT())
				case _col != 0 && _col != colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.CT())
					drawCanvas.RightTop = shape.NewPoint(character.CT())
					drawCanvas.RightBottom = shape.NewPoint(character.BT())
					drawCanvas.LeftBottom = shape.NewPoint(character.BT())
				case _col != 0 && _col == colCount-1:
					drawCanvas.LeftTop = shape.NewPoint(character.CT())
					drawCanvas.RightTop = shape.NewPoint(character.RT())
					drawCanvas.LeftBottom = shape.NewPoint(character.BT())
				}
			}
			t.objects[uint(_row)][uint(_col+1)] = core.NewObject(
				core.Coordinate{X: objX, Y: objY}, &drawCanvas,
			)
			objX += int(drawCanvas.Width()) - 1
		}
		objY += int(t.rowMaxHeight[_row] + character.TabWidth())
	}
}

func (t *Table) Draw(ctx core.RenderContext, c core.Coordinate) {
	t.SetCoordinate(c)
	coincidenceCtx, has := ctx.CoincidenceCheck(t)
	if !has {
		return
	}
	for _, colMap := range t.objects {
		for _, o := range colMap {
			o.D.Draw(coincidenceCtx, core.Coordinate{
				X: c.X + o.C.X,
				Y: c.Y + o.C.Y,
			})
		}
	}
}
