package table

import (
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
		cellWidthStartInCol += length + 1
	}
	return -1, -1, 0
}

func (t AdaptiveTable) calculateCellHeightEndIndex(rowRelativeIndex int) (int, int, int) {
	cellHeightStartInRow := 1
	for index, length := range t.rowMaxHeightMap {
		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1 {
			return index, cellHeightStartInRow, length
		}
		cellHeightStartInRow += length + 1
	}
	return -1, -1, 0
}

func (t AdaptiveTable) calculateCellContentRune(cellX, cellY, contentIndex int) rune {
	content := t.contentMap[cellY][cellX]
	if contentIndex < len(content) {
		return rune(content[contentIndex])
	}
	return ' '
}
