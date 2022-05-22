package table

import (
	"fmt"

	"github.com/Mericusta/go-vt100/core"
)

type tableViewInterface interface {
	calculateTableWidth() int
	calculateTableHeight() int
	calculateCellWidthInfo(int) (int, int, int)
	calculateCellHeightInfo(int) (int, int, int)
	calculateCellContentRune(int, int, int, int) rune
}

type Table struct {
	s  core.Size
	i  tableViewInterface
	fc core.Color
	bc core.Color
}

func (t Table) Draw(x, y int, s core.Size) {
	t.s.Width = t.i.calculateTableWidth()
	t.s.Height = t.i.calculateTableHeight()
	core.SetForegroundColor(t.fc)
	core.SetBackgroundColor(t.bc)
	for _y := y; _y < y+t.s.Height && _y < s.Height; _y++ {
		for _x := x; _x < x+t.s.Width && _x < s.Width; _x++ {
			colRelativeIndex, rowRelativeIndex := _x-x, _y-y
			cellX, cellWidthStartIndex, cellWidth := t.i.calculateCellWidthInfo(colRelativeIndex)
			cellY, cellHeightStartIndex, cellHeight := t.i.calculateCellHeightInfo(rowRelativeIndex)
			switch {
			case colRelativeIndex == 0:
				switch {
				case rowRelativeIndex == 0:
					// left top
					core.MoveCursorToAndPrint(_x, _y, string(core.TL()))
				case rowRelativeIndex == t.s.Height-1:
					// left bottom
					core.MoveCursorToAndPrint(_x, _y, string(core.BL()))
				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
					// left T
					core.MoveCursorToAndPrint(_x, _y, string(core.LT()))
				default:
					core.MoveCursorToAndPrint(_x, _y, string(core.VL()))
				}
			case colRelativeIndex == t.s.Width-1:
				switch {
				case rowRelativeIndex == 0:
					// right top
					core.MoveCursorToAndPrint(_x, _y, string(core.TR()))
				case rowRelativeIndex == t.s.Height-1:
					// right bottom
					core.MoveCursorToAndPrint(_x, _y, string(core.BR()))
				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
					// right T
					core.MoveCursorToAndPrint(_x, _y, string(core.RT()))
				default:
					core.MoveCursorToAndPrint(_x, _y, string(core.VL()))
				}
			case colRelativeIndex == cellWidthStartIndex+cellWidth:
				switch {
				case rowRelativeIndex == 0:
					// top T
					core.MoveCursorToAndPrint(_x, _y, string(core.TT()))
				case rowRelativeIndex == t.s.Height-1:
					// bottom T
					core.MoveCursorToAndPrint(_x, _y, string(core.BT()))
				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
					// center T
					core.MoveCursorToAndPrint(_x, _y, string(core.CT()))
				default:
					core.MoveCursorToAndPrint(_x, _y, string(core.VL()))
				}
			default:
				switch {
				case rowRelativeIndex == 0 || rowRelativeIndex == t.s.Height-1 || rowRelativeIndex == cellHeightStartIndex+cellHeight:
					core.MoveCursorToAndPrint(_x, _y, string(core.HL()))
				default:
					core.MoveCursorToAndPrint(_x, _y, string(t.i.calculateCellContentRune(cellX, cellY, colRelativeIndex-cellWidthStartIndex, rowRelativeIndex-cellHeightStartIndex)))
					// core.MoveCursorToAndPrint(_x, _y, " ")
				}
			}
		}
	}
	core.ClearForegroundColor()
	core.ClearBackgroundColor()
}

func Draw(t tableViewInterface) {
	tableWidth := t.calculateTableWidth()
	tableHeight := t.calculateTableHeight()
	totalPoints := tableWidth * tableHeight
	tableRuneSlice := make([]rune, totalPoints)
	for index := 0; index != totalPoints; index++ {
		colRelativeIndex, rowRelativeIndex := core.TransformArrayIndexToMatrixCoordinates(index, tableWidth, tableHeight)
		cellX, cellWidthStartIndex, cellWidth := t.calculateCellWidthInfo(colRelativeIndex)
		cellY, cellHeightStartIndex, cellHeight := t.calculateCellHeightInfo(rowRelativeIndex)
		switch {
		case colRelativeIndex == 0:
			switch {
			case rowRelativeIndex == 0:
				// left top
				tableRuneSlice[index] = core.TL()
			case rowRelativeIndex == tableHeight-1:
				// left bottom
				tableRuneSlice[index] = core.BL()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// left T
				tableRuneSlice[index] = core.LT()
			default:
				tableRuneSlice[index] = core.VL()
			}
		case colRelativeIndex == tableWidth-2:
			switch {
			case rowRelativeIndex == 0:
				// right top
				tableRuneSlice[index] = core.TR()
			case rowRelativeIndex == tableHeight-1:
				// right bottom
				tableRuneSlice[index] = core.BR()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// right T
				tableRuneSlice[index] = core.RT()
			default:
				tableRuneSlice[index] = core.VL()
			}
			index++
			tableRuneSlice[index] = core.EndLine()
			// fmt.Printf("\n%v", string(tableRuneSlice))
		case colRelativeIndex == cellWidthStartIndex+cellWidth:
			switch {
			case rowRelativeIndex == 0:
				// top T
				tableRuneSlice[index] = core.TT()
			case rowRelativeIndex == tableHeight-1:
				// bottom T
				tableRuneSlice[index] = core.BT()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// center T
				tableRuneSlice[index] = core.CT()
			default:
				tableRuneSlice[index] = core.VL()
			}
		default:
			switch {
			case rowRelativeIndex == 0 || rowRelativeIndex == tableHeight-1 || rowRelativeIndex == cellHeightStartIndex+cellHeight:
				tableRuneSlice[index] = core.HL()
			default:
				tableRuneSlice[index] = t.calculateCellContentRune(cellX, cellY, colRelativeIndex-cellWidthStartIndex, rowRelativeIndex-cellHeightStartIndex)
			}
		}
		// fmt.Printf("index %v, rune %v\n", index, string(tableRuneSlice[index]))
	}
	fmt.Printf("%v", string(tableRuneSlice))
}
