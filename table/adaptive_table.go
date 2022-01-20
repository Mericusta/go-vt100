package table

import (
	"fmt"
	"go-vt100/tab"
)

type AdaptiveTable struct {
	Col             int
	Row             int
	colMaxWidthMap  []int
	rowMaxHeightMap []int
	contentMap      [][]string
}

func NewAdaptiveTable(headSlice []string, lineContentSlice [][]string) *AdaptiveTable {
	t := &AdaptiveTable{}
	t.Col = len(headSlice)
	t.Row = 1 + len(lineContentSlice)
	t.contentMap = make([][]string, t.Row)
	t.rowMaxHeightMap = make([]int, t.Row)
	t.colMaxWidthMap = make([]int, t.Col)
	for index, head := range headSlice {
		t.colMaxWidthMap[index] = len(head)
	}
	t.contentMap[0] = headSlice
	t.rowMaxHeightMap[0] = 1
	for rowIndex, lineContent := range lineContentSlice {
		for colindex, content := range lineContent {
			if colindex >= t.Col {
				continue
			}
			if contentLength := len(content); t.colMaxWidthMap[colindex] < contentLength {
				t.colMaxWidthMap[colindex] = contentLength
			}
		}
		t.contentMap[rowIndex+1] = lineContent
		t.rowMaxHeightMap[rowIndex+1] = 1
	}
	return t
}

func (t AdaptiveTable) drawWithOneLoop() {
	tableWidth := t.calculateTableWidth()
	tableHeight := t.calculateTableHeight()
	totalPoints := tableWidth * tableHeight
	lineRuneSlice := make([]rune, totalPoints)
	for index := 0; index != totalPoints; index++ {
		colRelativeIndex := index % tableWidth
		rowRelativeIndex := index / tableWidth
		cellX, cellWidthStartIndex, cellWidth := t.calculateCellWidthEndIndex(colRelativeIndex)
		cellY, cellHeightStartIndex, cellHeight := t.calculateCellHeightEndIndex(rowRelativeIndex)
		// fmt.Printf("cellX = %v, cellWidthStartIndex = %v, cellWidth = %v\n", cellX, cellWidthStartIndex, cellWidth)
		// fmt.Printf("cellY = %v, cellHeightStartIndex = %v, cellHeight = %v\n", cellY, cellHeightStartIndex, cellHeight)
		// fmt.Printf("colRelativeIndex = %v, rowRelativeIndex = %v\n", colRelativeIndex, rowRelativeIndex)
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
				content := t.contentMap[cellY][cellX]
				if colRelativeIndex-cellWidthStartIndex >= len(content) {
					lineRuneSlice[index] = ' '
				} else {
					// fmt.Printf("content := t.contentMap[cellY][cellX] = %v, contentRune = %v\n", content, string(content[colRelativeIndex-cellWidthStartIndex]))
					lineRuneSlice[index] = rune(content[colRelativeIndex-cellWidthStartIndex])
				}
			}
		}
	}
	fmt.Printf("%v", string(lineRuneSlice))
}

func (t AdaptiveTable) calculateTableWidth() int {
	tableWidth := 0
	for _, columnWidth := range t.colMaxWidthMap {
		tableWidth += columnWidth
	}
	return tableWidth + tab.Width()*(t.Col+1) + 1
}

func (t AdaptiveTable) calculateTableHeight() int {
	return 1*t.Row + tab.Width()*(t.Row+1)
}

func (t AdaptiveTable) calculateCellWidthEndIndex(colRelativeIndex int) (int, int, int) {
	cellWidthStartInCol := 1
	for index, length := range t.colMaxWidthMap {
		if cellWidthStartInCol <= colRelativeIndex && colRelativeIndex <= cellWidthStartInCol+length {
			return index, cellWidthStartInCol, length
		}
		cellWidthStartInCol += (length + 1)
	}
	return -1, -1, 0
}

func (t AdaptiveTable) calculateCellHeightEndIndex(rowRelativeIndex int) (int, int, int) {
	cellHeightStartInRow := 1
	for index, length := range t.rowMaxHeightMap {
		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1 {
			return index, cellHeightStartInRow, length
		}
		cellHeightStartInRow += (length + 1)
	}
	return -1, -1, 0
}

func (t AdaptiveTable) Draw() {
	t.drawWithOneLoop()
}
