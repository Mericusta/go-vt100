package table

import (
	"fmt"
	"go-vt100/tab"
)

type Table struct {
	Col        int
	Row        int
	CellWidth  int
	CellHeight int
	Content    []byte
	pos        Position
}

type Position int

const (
	LINE Position = iota + 1
	CONTENT
)

func (t Table) drawWithNestedLoop() {
	tableWidth := t.CellWidth*t.Col + tab.Width()*(t.Col+1)
	tableHeight := t.CellHeight*t.Row + tab.Width()*(t.Row+1)

	for drawHeightIndex := 0; drawHeightIndex != tableHeight; drawHeightIndex++ {
		lineRuneSlice := make([]rune, tableWidth)
		for drawWidthIndex := 0; drawWidthIndex != tableWidth; drawWidthIndex++ {
			switch {
			case drawWidthIndex == 0:
				switch {
				case drawHeightIndex == 0:
					// left top
					lineRuneSlice[drawWidthIndex] = tab.TL()
				case drawHeightIndex == tableHeight-1:
					// left bottom
					lineRuneSlice[drawWidthIndex] = tab.BL()
				case drawHeightIndex%(t.CellHeight+1) == 0:
					// left T
					lineRuneSlice[drawWidthIndex] = tab.LT()
				default:
					lineRuneSlice[drawWidthIndex] = tab.VL()
				}
			case drawWidthIndex == tableWidth-1:
				switch {
				case drawHeightIndex == 0:
					// right top
					lineRuneSlice[drawWidthIndex] = tab.TR()
				case drawHeightIndex == tableHeight-1:
					// right bottom
					lineRuneSlice[drawWidthIndex] = tab.BR()
				case drawHeightIndex%(t.CellHeight+1) == 0:
					// right T
					lineRuneSlice[drawWidthIndex] = tab.RT()
				default:
					lineRuneSlice[drawWidthIndex] = tab.VL()
				}
			case drawWidthIndex%(t.CellWidth+1) == 0:
				switch {
				case drawHeightIndex == 0:
					// top T
					lineRuneSlice[drawWidthIndex] = tab.TT()
				case drawHeightIndex == tableHeight-1:
					// bottom T
					lineRuneSlice[drawWidthIndex] = tab.BT()
				case drawHeightIndex%(t.CellHeight+1) == 0:
					// center T
					lineRuneSlice[drawWidthIndex] = tab.CT()
				default:
					lineRuneSlice[drawWidthIndex] = tab.VL()
				}
			default:
				switch {
				case drawHeightIndex == 0 || drawHeightIndex == tableHeight-1 || drawHeightIndex%(t.CellHeight+1) == 0:
					lineRuneSlice[drawWidthIndex] = tab.HL()
				default:
					lineRuneSlice[drawWidthIndex] = rune(t.Content[drawWidthIndex%(t.CellWidth+1)-1])
				}
			}
		}
		fmt.Printf("%v\n", string(lineRuneSlice))
	}
}

func (t Table) drawWithOneLoop() {
	tableWidth := t.CellWidth*t.Col + tab.Width()*(t.Col+1) + 1
	tableHeight := t.CellHeight*t.Row + tab.Width()*(t.Row+1)
	totalPoints := tableWidth * tableHeight

	lineRuneSlice := make([]rune, totalPoints)
	for index := 0; index != totalPoints; index++ {
		colRelativeIndex := index % tableWidth
		rowRelativeIndex := index / tableWidth
		switch {
		case colRelativeIndex == 0:
			switch {
			case rowRelativeIndex == 0:
				// left top
				lineRuneSlice[index] = tab.TL()
			case rowRelativeIndex == tableHeight-1:
				// left bottom
				lineRuneSlice[index] = tab.BL()
			case rowRelativeIndex%(t.CellHeight+1) == 0:
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
			case rowRelativeIndex%(t.CellHeight+1) == 0:
				// right T
				lineRuneSlice[index] = tab.RT()
			default:
				lineRuneSlice[index] = tab.VL()
			}
			index++
			lineRuneSlice[index] = '\n'
		case colRelativeIndex%(t.CellWidth+1) == 0:
			switch {
			case rowRelativeIndex == 0:
				// top T
				lineRuneSlice[index] = tab.TT()
			case rowRelativeIndex == tableHeight-1:
				// bottom T
				lineRuneSlice[index] = tab.BT()
			case rowRelativeIndex%(t.CellHeight+1) == 0:
				// center T
				lineRuneSlice[index] = tab.CT()
			default:
				lineRuneSlice[index] = tab.VL()
			}
		default:
			switch {
			case rowRelativeIndex == 0 || rowRelativeIndex == tableHeight-1 || rowRelativeIndex%(t.CellHeight+1) == 0:
				lineRuneSlice[index] = tab.HL()
			default:
				lineRuneSlice[index] = rune(t.Content[colRelativeIndex%(t.CellWidth+1)-1])
			}
		}
	}
	fmt.Printf("%v", string(lineRuneSlice))
}

func (t Table) Draw() {
	t.drawWithOneLoop()
	// t.drawWithNestedLoop()
}

// cell height = 3
//    index cal
// TL 0
// VL 1
// VL 2
// VL 3
// LT 4    index % (cell height + 1) = 0
// VL 5
// VL 6
// VL 7
// LT 8    index % (cell height + 1) = 0
// VL 9
// VL 10
// VL 11
// BL 12
