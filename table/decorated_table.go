package table

// import (
// 	"github.com/Mericusta/go-vt100/character"
// 	"github.com/Mericusta/go-vt100/core"
// )

// type DecoratedTable struct {
// 	*AdaptiveCellTable
// 	*TableDecoration
// }

// type TableDecoration struct {
// 	CellWidthPadding  int
// 	CellHeightPadding int
// }

// func NewDecoratedTable(headSlice []string, lineContentSlice [][]string, decoration *TableDecoration, fc, bc core.Color) *DecoratedTable {
// 	t := &DecoratedTable{}
// 	t.AdaptiveCellTable = NewAdaptiveCellTable(headSlice, lineContentSlice, fc, bc)
// 	t.TableDecoration = decoration
// 	for index := range t.colMaxWidthMap {
// 		t.colMaxWidthMap[index] += t.CellWidthPadding * 2
// 	}
// 	t.i = t
// 	return t
// }

// func (t DecoratedTable) calculateTableHeight() int {
// 	return (1+t.CellHeightPadding*2)*t.row + int(character.TabWidth())*(t.row+1)
// }

// func (t DecoratedTable) calculateCellHeightInfo(rowRelativeIndex int) (int, int, int) {
// 	cellHeightStartInRow := 1
// 	for index, length := range t.rowMaxHeightMap {
// 		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1+t.CellHeightPadding*2 {
// 			return index, cellHeightStartInRow, length + t.CellHeightPadding*2
// 		}
// 		cellHeightStartInRow += length + 1 + t.CellHeightPadding*2
// 	}
// 	return -1, -1, 0
// }

// func (t DecoratedTable) calculateCellContentRune(cellX, cellY, contentColIndex, contentRowIndex int) rune {
// 	content := t.contentMap[cellY][cellX]
// 	contentLength := len(content)
// 	if contentColIndex < t.CellWidthPadding || t.CellWidthPadding+contentLength <= contentColIndex {
// 		return character.Space()
// 	} else if contentRowIndex < t.CellHeightPadding || t.CellHeightPadding+1 <= contentRowIndex {
// 		return character.Space()
// 	}
// 	return rune(content[contentColIndex-t.CellWidthPadding])
// }
