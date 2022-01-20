package table

import (
	"go-vt100/tab"
	"strings"
)

type DecoratedTable struct {
	*AdaptiveCellTable
	d *TableDecoration
}

type TableDecoration struct {
	WidthPadding  int
	HeightPadding int
}

func NewDecoratedTable(headSlice []string, lineContentSlice [][]string, decoration *TableDecoration) *DecoratedTable {
	t := &DecoratedTable{}
	t.AdaptiveCellTable = NewAdaptiveCellTable(headSlice, lineContentSlice)
	t.d = decoration
	for index := range t.colMaxWidthMap {
		t.colMaxWidthMap[index] += t.d.WidthPadding * 2
	}

	builder := strings.Builder{}
	for cellY, contentSlice := range t.contentMap {
		for cellX, content := range contentSlice {
			builder.Reset()
			for index := 0; index != t.d.WidthPadding; index++ {
				builder.WriteRune(' ')
			}
			builder.WriteString(content)
			for index := 0; index != t.d.WidthPadding; index++ {
				builder.WriteRune(' ')
			}
			t.contentMap[cellY][cellX] = builder.String()
		}
	}

	return t
}

func (t DecoratedTable) calculateTableHeight() int {
	return (1+t.d.HeightPadding*2)*t.row + tab.Width()*(t.row+1)
}

func (t DecoratedTable) calculateCellHeightInfo(rowRelativeIndex int) (int, int, int) {
	cellHeightStartInRow := 1
	for index, length := range t.rowMaxHeightMap {
		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1+t.d.HeightPadding*2 {
			return index, cellHeightStartInRow, length
		}
		cellHeightStartInRow += length + 1 + t.d.HeightPadding*2
	}
	return -1, -1, 0
}
