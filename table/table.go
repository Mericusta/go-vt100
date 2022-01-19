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
	pos        Position
}

type Position int

const (
	LINE Position = iota + 1
	CONTENT
)

func (t Table) Draw() {
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
					lineRuneSlice[drawWidthIndex] = ' '
				}
			}
		}
		fmt.Printf("%v\n", string(lineRuneSlice))
	}
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
