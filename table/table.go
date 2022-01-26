package table

import (
	"fmt"
	"go-vt100/canvas"
	"go-vt100/color"
	"go-vt100/size"
	"go-vt100/tab"
	"go-vt100/vt100"
)

type tableInterface interface {
	calculateTableWidth() int
	calculateTableHeight() int
	calculateCellWidthInfo(int) (int, int, int)
	calculateCellHeightInfo(int) (int, int, int)
	calculateCellContentRune(int, int, int, int) rune
}

type Table struct {
	s  size.Size
	i  tableInterface
	fc color.Color
	bc color.Color
}

func (t Table) Draw(x, y int) {
	vt100.SetForegroundColor(t.fc)
	vt100.SetBackgroundColor(t.bc)
	for _y := y; _y < y+t.s.Height; _y++ {
		for _x := x; _x < x+t.s.Width; _x++ {
			colRelativeIndex, rowRelativeIndex := _x-x, _y-y
			cellX, cellWidthStartIndex, cellWidth := t.i.calculateCellWidthInfo(colRelativeIndex)
			cellY, cellHeightStartIndex, cellHeight := t.i.calculateCellHeightInfo(rowRelativeIndex)
			switch {
			case colRelativeIndex == 0:
				switch {
				case rowRelativeIndex == 0:
					// left top
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.TL()))
				case rowRelativeIndex == t.s.Height-1:
					// left bottom
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.BL()))
				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
					// left T
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.LT()))
				default:
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.VL()))
				}
			case colRelativeIndex == t.s.Width-1:
				switch {
				case rowRelativeIndex == 0:
					// right top
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.TR()))
				case rowRelativeIndex == t.s.Height-1:
					// right bottom
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.BR()))
				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
					// right T
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.RT()))
				default:
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.VL()))
				}
			case colRelativeIndex == cellWidthStartIndex+cellWidth:
				switch {
				case rowRelativeIndex == 0:
					// top T
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.TT()))
				case rowRelativeIndex == t.s.Height-1:
					// bottom T
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.BT()))
				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
					// center T
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.CT()))
				default:
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.VL()))
				}
			default:
				switch {
				case rowRelativeIndex == 0 || rowRelativeIndex == t.s.Height-1 || rowRelativeIndex == cellHeightStartIndex+cellHeight:
					vt100.MoveCursorToAndPrint(_x, _y, string(tab.HL()))
				default:
					vt100.MoveCursorToAndPrint(_x, _y, string(t.i.calculateCellContentRune(cellX, cellY, colRelativeIndex-cellWidthStartIndex, rowRelativeIndex-cellHeightStartIndex)))
					// vt100.MoveCursorToAndPrint(_x, _y, " ")
				}
			}
		}
	}
	vt100.ClearForegroundColor()
	vt100.ClearBackgroundColor()
}

func Draw(t tableInterface) {
	tableWidth := t.calculateTableWidth()
	tableHeight := t.calculateTableHeight()
	totalPoints := tableWidth * tableHeight
	tableRuneSlice := make([]rune, totalPoints)
	for index := 0; index != totalPoints; index++ {
		colRelativeIndex, rowRelativeIndex := canvas.TransformArrayIndexToMatrixCoordinates(index, tableWidth, tableHeight)
		cellX, cellWidthStartIndex, cellWidth := t.calculateCellWidthInfo(colRelativeIndex)
		cellY, cellHeightStartIndex, cellHeight := t.calculateCellHeightInfo(rowRelativeIndex)
		switch {
		case colRelativeIndex == 0:
			switch {
			case rowRelativeIndex == 0:
				// left top
				tableRuneSlice[index] = tab.TL()
			case rowRelativeIndex == tableHeight-1:
				// left bottom
				tableRuneSlice[index] = tab.BL()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// left T
				tableRuneSlice[index] = tab.LT()
			default:
				tableRuneSlice[index] = tab.VL()
			}
		case colRelativeIndex == tableWidth-2:
			switch {
			case rowRelativeIndex == 0:
				// right top
				tableRuneSlice[index] = tab.TR()
			case rowRelativeIndex == tableHeight-1:
				// right bottom
				tableRuneSlice[index] = tab.BR()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// right T
				tableRuneSlice[index] = tab.RT()
			default:
				tableRuneSlice[index] = tab.VL()
			}
			index++
			tableRuneSlice[index] = tab.EndLine()
			// fmt.Printf("\n%v", string(tableRuneSlice))
		case colRelativeIndex == cellWidthStartIndex+cellWidth:
			switch {
			case rowRelativeIndex == 0:
				// top T
				tableRuneSlice[index] = tab.TT()
			case rowRelativeIndex == tableHeight-1:
				// bottom T
				tableRuneSlice[index] = tab.BT()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// center T
				tableRuneSlice[index] = tab.CT()
			default:
				tableRuneSlice[index] = tab.VL()
			}
		default:
			switch {
			case rowRelativeIndex == 0 || rowRelativeIndex == tableHeight-1 || rowRelativeIndex == cellHeightStartIndex+cellHeight:
				tableRuneSlice[index] = tab.HL()
			default:
				tableRuneSlice[index] = t.calculateCellContentRune(cellX, cellY, colRelativeIndex-cellWidthStartIndex, rowRelativeIndex-cellHeightStartIndex)
			}
		}
		// fmt.Printf("index %v, rune %v\n", index, string(tableRuneSlice[index]))
	}
	fmt.Printf("%v", string(tableRuneSlice))
}
