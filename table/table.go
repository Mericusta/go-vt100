package table

import (
	"fmt"
	"go-vt100/canvas"
	"go-vt100/tab"
)

type Table interface {
	calculateTableWidth() int
	calculateTableHeight() int
	calculateCellWidthInfo(int) (int, int, int)
	calculateCellHeightInfo(int) (int, int, int)
	calculateCellContentRune(int, int, int, int) rune
}

func Draw(t Table) {
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
