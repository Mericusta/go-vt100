package table

import (
	"fmt"
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
	lineRuneSlice := make([]rune, totalPoints)
	for index := 0; index != totalPoints; index++ {
		colRelativeIndex := index % tableWidth
		rowRelativeIndex := index / tableWidth
		cellX, cellWidthStartIndex, cellWidth := t.calculateCellWidthInfo(colRelativeIndex)
		cellY, cellHeightStartIndex, cellHeight := t.calculateCellHeightInfo(rowRelativeIndex)
		switch {
		case colRelativeIndex == 0:
			switch {
			case rowRelativeIndex == 0:
				// left top
				lineRuneSlice[index] = tab.TL()
			case rowRelativeIndex == tableHeight-1:
				// left bottom
				lineRuneSlice[index] = tab.BL()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// left T
				lineRuneSlice[index] = tab.LT()
			default:
				lineRuneSlice[index] = tab.VL()
			}
		case colRelativeIndex == tableWidth-2:
			switch {
			case rowRelativeIndex == 0:
				// right top
				lineRuneSlice[index] = tab.TR()
			case rowRelativeIndex == tableHeight-1:
				// right bottom
				lineRuneSlice[index] = tab.BR()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// right T
				lineRuneSlice[index] = tab.RT()
			default:
				lineRuneSlice[index] = tab.VL()
			}
			index++
			lineRuneSlice[index] = '\n'
			// fmt.Printf("\n%v", string(lineRuneSlice))
		case colRelativeIndex == cellWidthStartIndex+cellWidth:
			switch {
			case rowRelativeIndex == 0:
				// top T
				lineRuneSlice[index] = tab.TT()
			case rowRelativeIndex == tableHeight-1:
				// bottom T
				lineRuneSlice[index] = tab.BT()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// center T
				lineRuneSlice[index] = tab.CT()
			default:
				lineRuneSlice[index] = tab.VL()
			}
		default:
			switch {
			case rowRelativeIndex == 0 || rowRelativeIndex == tableHeight-1 || rowRelativeIndex == cellHeightStartIndex+cellHeight:
				lineRuneSlice[index] = tab.HL()
			default:
				lineRuneSlice[index] = t.calculateCellContentRune(cellX, cellY, colRelativeIndex-cellWidthStartIndex, rowRelativeIndex-cellHeightStartIndex)
			}
		}
		// fmt.Printf("index %v, rune %v\n", index, string(lineRuneSlice[index]))
	}
	fmt.Printf("%v", string(lineRuneSlice))
}
