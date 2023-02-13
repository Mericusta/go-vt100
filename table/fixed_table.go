package table

import "github.com/Mericusta/go-vt100/core"

type FixedCellTable struct {
	table
	col        uint
	row        uint
	CellWidth  uint
	CellHeight uint
	Content    []byte
}

func NewFixedCellTable(row, col uint, content string, fc, bc core.Color) *FixedCellTable {
	l := uint(len(content))
	if l == 0 {
		content = "fixed cell table"
	}
	t := &FixedCellTable{
		table: table{
			fc: fc,
			bc: bc,
		},
		col:        col,
		row:        row,
		CellWidth:  l,
		CellHeight: 1,
		Content:    []byte(content),
	}
	// t.i = t
	return t
}

// func (t FixedCellTable) calculateTableWidth() uint {
// 	return uint(t.CellWidth*t.col) + character.TabWidth()*(t.col+1)
// }

// func (t FixedCellTable) calculateTableHeight() uint {
// 	return t.CellHeight*t.row + character.TabWidth()*(t.row+1)
// }

// func (t FixedCellTable) calculateCellWidthInfo(colRelativeIndex uint) (uint, uint, uint) {
// 	cellWidthStartInCol := 1
// 	for index := uint(0); index != t.col; index++ {
// 		if cellWidthStartInCol <= colRelativeIndex && colRelativeIndex <= cellWidthStartInCol+t.CellWidth {
// 			return index, cellWidthStartInCol, t.CellWidth
// 		}
// 		cellWidthStartInCol += t.CellWidth + 1
// 	}
// 	return -1, -1, 0
// }

// func (t FixedCellTable) calculateCellHeightInfo(rowRelativeIndex int) (int, int, int) {
// 	cellHeightStartInRow := 1
// 	for index := 0; index != t.row; index++ {
// 		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1 {
// 			return index, cellHeightStartInRow, t.CellHeight
// 		}
// 		cellHeightStartInRow += t.CellHeight + 1
// 	}
// 	return -1, -1, 0
// }

// func (t FixedCellTable) calculateCellContentRune(cellX, cellY, contentColIndex, contentRowIndex int) rune {
// 	return rune(t.Content[contentColIndex])
// }
