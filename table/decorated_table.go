package table

import (
	"go-vt100/color"
	"go-vt100/tab"
)

type DecoratedTable struct {
	Table
	*AdaptiveCellTable
	*TableDecoration
}

type TableDecoration struct {
	WidthPadding  int
	HeightPadding int
}

func NewDecoratedTable(headSlice []string, lineContentSlice [][]string, decoration *TableDecoration, fc, bc color.Color) *DecoratedTable {
	t := &DecoratedTable{}
	t.AdaptiveCellTable = NewAdaptiveCellTable(headSlice, lineContentSlice, fc, bc)
	t.TableDecoration = decoration
	for index := range t.colMaxWidthMap {
		t.colMaxWidthMap[index] += t.WidthPadding * 2
	}
	t.i = t
	return t
}

func (t DecoratedTable) calculateTableHeight() int {
	return (1+t.HeightPadding*2)*t.row + tab.Width()*(t.row+1)
}

func (t DecoratedTable) calculateCellHeightInfo(rowRelativeIndex int) (int, int, int) {
	cellHeightStartInRow := 1
	for index, length := range t.rowMaxHeightMap {
		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1+t.HeightPadding*2 {
			return index, cellHeightStartInRow, length + t.HeightPadding*2
		}
		cellHeightStartInRow += length + 1 + t.HeightPadding*2
	}
	return -1, -1, 0
}

func (t DecoratedTable) calculateCellContentRune(cellX, cellY, contentColIndex, contentRowIndex int) rune {
	content := t.contentMap[cellY][cellX]
	contentLength := len(content)
	if contentColIndex < t.WidthPadding || t.WidthPadding+contentLength <= contentColIndex {
		return tab.Space()
	} else if contentRowIndex < t.HeightPadding || t.HeightPadding+1 <= contentRowIndex {
		return tab.Space()
	}
	return rune(content[contentColIndex-t.WidthPadding])
}
