package table

import (
	"go-vt100/tab"
)

type FixedCellTable struct {
	col        int
	row        int
	CellWidth  int
	CellHeight int
	Content    []byte
}

func NewFixedCellTable(row, col int, content string) *FixedCellTable {
	if len(content) == 0 {
		content = "fixed cell table"
	}
	return &FixedCellTable{
		col:        row,
		row:        col,
		CellWidth:  len(content),
		CellHeight: 1,
		Content:    []byte(content),
	}
}

func (t FixedCellTable) calculateTableWidth() int {
	return t.CellWidth*t.col + tab.Width()*(t.col+1) + 1
}

func (t FixedCellTable) calculateTableHeight() int {
	return t.CellHeight*t.row + tab.Width()*(t.row+1)
}

func (t FixedCellTable) calculateCellWidthInfo(colRelativeIndex int) (int, int, int) {
	cellWidthStartInCol := 1
	for index := 0; index != t.col; index++ {
		if cellWidthStartInCol <= colRelativeIndex && colRelativeIndex <= cellWidthStartInCol+t.CellWidth {
			return index, cellWidthStartInCol, t.CellWidth
		}
		cellWidthStartInCol += t.CellWidth + 1
	}
	return -1, -1, 0
}

func (t FixedCellTable) calculateCellHeightInfo(rowRelativeIndex int) (int, int, int) {
	cellHeightStartInRow := 1
	for index := 0; index != t.row; index++ {
		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1 {
			return index, cellHeightStartInRow, t.CellHeight
		}
		cellHeightStartInRow += t.CellHeight + 1
	}
	return -1, -1, 0
}

func (t FixedCellTable) calculateCellContentRune(cellX, cellY, contentColIndex, contentRowIndex int) rune {
	return rune(t.Content[contentRowIndex])
}
