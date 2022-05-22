package table

import (
	"github.com/Mericusta/go-vt100/core"
)

type AdaptiveCellTable struct {
	Table
	col             int
	row             int
	colMaxWidthMap  []int
	rowMaxHeightMap []int
	contentMap      [][]string
}

func NewAdaptiveCellTable(headSlice []string, lineContentSlice [][]string, fc, bc core.Color) *AdaptiveCellTable {
	t := &AdaptiveCellTable{
		Table: Table{
			fc: fc,
			bc: bc,
		},
	}
	t.col = len(headSlice)
	t.row = 1 + len(lineContentSlice)
	t.contentMap = make([][]string, t.row)
	t.rowMaxHeightMap = make([]int, t.row)
	t.colMaxWidthMap = make([]int, t.col)
	for index, head := range headSlice {
		t.colMaxWidthMap[index] = len(head)
	}
	t.contentMap[0] = headSlice
	t.rowMaxHeightMap[0] = 1
	for rowIndex, lineContent := range lineContentSlice {
		for coliIndex, content := range lineContent {
			if coliIndex >= t.col {
				continue
			}
			if contentLength := len(content); t.colMaxWidthMap[coliIndex] < contentLength {
				t.colMaxWidthMap[coliIndex] = contentLength
			}
		}
		t.contentMap[rowIndex+1] = lineContent
		t.rowMaxHeightMap[rowIndex+1] = 1
	}
	t.i = t
	return t
}

func (t AdaptiveCellTable) calculateTableWidth() int {
	tableWidth := 0
	for _, columnWidth := range t.colMaxWidthMap {
		tableWidth += columnWidth
	}
	return tableWidth + core.Width()*(t.col+1)
}

func (t AdaptiveCellTable) calculateTableHeight() int {
	return 1*t.row + core.Width()*(t.row+1)
}

func (t AdaptiveCellTable) calculateCellWidthInfo(colRelativeIndex int) (int, int, int) {
	cellWidthStartInCol := 1
	for index, length := range t.colMaxWidthMap {
		if cellWidthStartInCol <= colRelativeIndex && colRelativeIndex <= cellWidthStartInCol+length {
			return index, cellWidthStartInCol, length
		}
		cellWidthStartInCol += length + 1
	}
	return -1, -1, 0
}

func (t AdaptiveCellTable) calculateCellHeightInfo(rowRelativeIndex int) (int, int, int) {
	cellHeightStartInRow := 1
	for index, length := range t.rowMaxHeightMap {
		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1 {
			return index, cellHeightStartInRow, length
		}
		cellHeightStartInRow += length + 1
	}
	return -1, -1, 0
}

func (t AdaptiveCellTable) calculateCellContentRune(cellX, cellY, contentColIndex, contentRowIndex int) rune {
	content := t.contentMap[cellY][cellX]
	if contentColIndex < len(content) {
		return rune(content[contentColIndex])
	}
	return core.Space()
}
