package table

import (
	"fmt"

	"github.com/Mericusta/go-vt100/character"
	"github.com/Mericusta/go-vt100/core"
)

type Table interface {
	Draw()
}

// New
// @param1 header slice
// @param2 row:col:content cell map, coordinate begin from 1
// @return table interface
func New(header []string, data map[uint]map[uint]string) Table {
	d := make(map[uint]map[uint]string)
	cellMaxWidth, cellMaxHeight := character.TabWidth(), character.TabHeight()
	maxCol, maxRow := uint(1), uint(1)

	data[1] = make(map[uint]string)
	data[1][1] = character.SpaceString()
	for i, s := range header {
		if l := uint(len(s)); l >= cellMaxWidth {
			cellMaxWidth = l
			data[1][uint(i+1)] = s
		}
	}

	if l := uint(len(header)); l >= maxCol {
		maxCol = l
	}

	for r := range d {
		if r >= maxRow {
			maxRow = r
		}
	}

	for r := uint(1); r <= maxRow; r++ {
		for c := uint(1); c < maxCol; c++ {
			var s string
			if _, hasRow := d[r]; hasRow {
				s = d[r][c]
			}
			data[r+1][c] = s
		}
	}

	s := core.Size{
		Width:  maxCol*cellMaxWidth + (maxCol+1)*character.TabWidth(),
		Height: maxRow*cellMaxHeight + (maxRow+1)*character.TabHeight(),
	}

	c := cell{
		width:  cellMaxWidth,
		height: cellMaxHeight,
	}

	// fmt.Printf("data = %v\n", data)
	// fmt.Printf("s = %v\n", s)

	return &table{
		col: maxCol,
		row: maxRow,
		d:   d,
		s:   s,
		c:   c,
	}
}

type cell struct {
	width  uint
	height uint
	// calculateTableWidth() uint
	// calculateTableHeight() uint
	// calculateCellWidthInfo(uint) (uint, uint, uint)
	// calculateCellHeightInfo(uint) (uint, uint, uint)
	// calculateCellContentRune(uint, uint, uint, uint) rune
}

type table struct {
	col uint                     // number of table column
	row uint                     // number of table row
	d   map[uint]map[uint]string // table content, row:col:s map, coordinate begin from 1
	s   core.Size                // table size including border
	c   cell                     // cell
	fc  core.Color               // foreground color
	bc  core.Color               // background color
}

func (t *table) Draw() {
	tableWidth := t.s.Width
	tableHeight := t.s.Height
	totalPoints := t.s.Width * t.s.Height
	tableRuneSlice := make([]rune, totalPoints)
	for index := uint(0); index != totalPoints; index++ {
		colRelativeIndex, rowRelativeIndex := core.TransformArrayIndexToMatrixCoordinates(index, tableWidth, tableHeight)
		fmt.Printf("index %v, colRelativeIndex %v, rowRelativeIndex %v\n", index, colRelativeIndex, rowRelativeIndex)
		cellX, cellWidthStartIndex, cellWidth := calculateCellWidthInfo(t.c.width, t.col, colRelativeIndex)
		cellY, cellHeightStartIndex, cellHeight := calculateCellHeightInfo(t.c.height, t.row, rowRelativeIndex)
		fmt.Printf("cellX %v, cellWidthStartIndex %v, cellWidth %v\n", cellX, cellWidthStartIndex, cellWidth)
		fmt.Printf("cellY %v, cellHeightStartIndex %v, cellHeight %v\n", cellY, cellHeightStartIndex, cellHeight)
		switch {
		case colRelativeIndex == 0:
			switch {
			case rowRelativeIndex == 0:
				// left top
				tableRuneSlice[index] = character.TL()
			case rowRelativeIndex == tableHeight-1:
				// left bottom
				tableRuneSlice[index] = character.BL()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// left T
				tableRuneSlice[index] = character.LT()
			default:
				tableRuneSlice[index] = character.VL()
			}
		case colRelativeIndex == tableWidth-2:
			switch {
			case rowRelativeIndex == 0:
				// right top
				tableRuneSlice[index] = character.TR()
			case rowRelativeIndex == tableHeight-1:
				// right bottom
				tableRuneSlice[index] = character.BR()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// right T
				tableRuneSlice[index] = character.RT()
			default:
				tableRuneSlice[index] = character.VL()
			}
			index++
			tableRuneSlice[index] = character.EndLine()
			// fmt.Printf("\n%v", string(tableRuneSlice))
		case colRelativeIndex == cellWidthStartIndex+cellWidth:
			switch {
			case rowRelativeIndex == 0:
				// top T
				tableRuneSlice[index] = character.TT()
			case rowRelativeIndex == tableHeight-1:
				// bottom T
				tableRuneSlice[index] = character.BT()
			case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
				// center T
				tableRuneSlice[index] = character.CT()
			default:
				tableRuneSlice[index] = character.VL()
			}
		default:
			switch {
			case rowRelativeIndex == 0 || rowRelativeIndex == tableHeight-1 || rowRelativeIndex == cellHeightStartIndex+cellHeight:
				tableRuneSlice[index] = character.HL()
			default:
				// tableRuneSlice[index] = calculateCellContentRune(cellX, cellY, colRelativeIndex-cellWidthStartIndex, rowRelativeIndex-cellHeightStartIndex)
				tableRuneSlice[index] = rune('a' + colRelativeIndex - cellWidthStartIndex)
			}
		}
		// fmt.Printf("index %v, rune %v\n", index, string(tableRuneSlice[index]))
	}
	fmt.Printf("%v", string(tableRuneSlice))
}

func calculateCellWidthInfo(cellWidth, col, colRelativeIndex uint) (uint, uint, uint) {
	cellWidthStartInCol := uint(1)
	for index := uint(0); index != col; index++ {
		if cellWidthStartInCol <= colRelativeIndex && colRelativeIndex <= cellWidthStartInCol+cellWidth {
			return index, cellWidthStartInCol, cellWidth
		}
		cellWidthStartInCol += cellWidth + 1
	}
	return 0, 0, 0
}

func calculateCellHeightInfo(cellHeight, row, rowRelativeIndex uint) (uint, uint, uint) {
	cellHeightStartInRow := uint(1)
	for index := uint(0); index != row; index++ {
		if cellHeightStartInRow <= rowRelativeIndex && rowRelativeIndex <= cellHeightStartInRow+1 {
			return index, cellHeightStartInRow, cellHeight
		}
		cellHeightStartInRow += cellHeight + 1
	}
	return 0, 0, 0
}

func calculateCellContentRune(cellX, cellY, contentColIndex, contentRowIndex uint) rune {
	return 0
	// return rune(t.Content[contentColIndex])
}

// func (t Table) Draw(x, y uint, s core.Size) {
// 	t.s.Width = t.i.calculateTableWidth()
// 	t.s.Height = t.i.calculateTableHeight()
// 	core.SetForegroundColor(t.fc)
// 	core.SetBackgroundColor(t.bc)
// 	for _y := y; _y < y+t.s.Height && _y < s.Height; _y++ {
// 		for _x := x; _x < x+t.s.Width && _x < s.Width; _x++ {
// 			colRelativeIndex, rowRelativeIndex := _x-x, _y-y
// 			cellX, cellWidthStartIndex, cellWidth := t.i.calculateCellWidthInfo(colRelativeIndex)
// 			cellY, cellHeightStartIndex, cellHeight := t.i.calculateCellHeightInfo(rowRelativeIndex)
// 			switch {
// 			case colRelativeIndex == 0:
// 				switch {
// 				case rowRelativeIndex == 0:
// 					// left top
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.TL()))
// 				case rowRelativeIndex == t.s.Height-1:
// 					// left bottom
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.BL()))
// 				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
// 					// left T
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.LT()))
// 				default:
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.VL()))
// 				}
// 			case colRelativeIndex == t.s.Width-1:
// 				switch {
// 				case rowRelativeIndex == 0:
// 					// right top
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.TR()))
// 				case rowRelativeIndex == t.s.Height-1:
// 					// right bottom
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.BR()))
// 				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
// 					// right T
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.RT()))
// 				default:
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.VL()))
// 				}
// 			case colRelativeIndex == cellWidthStartIndex+cellWidth:
// 				switch {
// 				case rowRelativeIndex == 0:
// 					// top T
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.TT()))
// 				case rowRelativeIndex == t.s.Height-1:
// 					// bottom T
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.BT()))
// 				case rowRelativeIndex == (cellHeightStartIndex + cellHeight):
// 					// center T
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.CT()))
// 				default:
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.VL()))
// 				}
// 			default:
// 				switch {
// 				case rowRelativeIndex == 0 || rowRelativeIndex == t.s.Height-1 || rowRelativeIndex == cellHeightStartIndex+cellHeight:
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(character.HL()))
// 				default:
// 					core.MoveCursorToAndPrint(int(_x), int(_y), string(t.i.calculateCellContentRune(cellX, cellY, colRelativeIndex-cellWidthStartIndex, rowRelativeIndex-cellHeightStartIndex)))
// 					// core.MoveCursorToAndPrint(int(_x), int(_y), " ")
// 				}
// 			}
// 		}
// 	}
// 	core.ClearForegroundColor()
// 	core.ClearBackgroundColor()
// }
