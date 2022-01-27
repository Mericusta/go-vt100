package tree

import (
	"fmt"
	"go-vt100/color"
	"go-vt100/coordinate"
	"go-vt100/size"
	"go-vt100/tab"
	"go-vt100/vt100"
	"strings"
)

type valueInterface interface {
	Show() string
}

type treeInterface interface {
	Value() valueInterface
	Children() []treeInterface
	Parent() treeInterface
	calculateChildTreeInfo() (int, int)
}

type Tree struct {
	s      size.Size
	i      treeInterface
	fc     color.Color
	bc     color.Color
	margin int
}

func (t Tree) Draw(x, y int) {
	vt100.SetForegroundColor(t.fc)
	vt100.SetBackgroundColor(t.bc)

	nodePosition := make(map[treeInterface]coordinate.Coordinate)
	nodePosition[t.i] = coordinate.Coordinate{
		X: x,
		Y: y,
	}
	bft(t.i, func(ti treeInterface) bool {
		pos, hasPos := nodePosition[ti]
		if !hasPos {
			panic(fmt.Sprintf("node %v not has position", ti.Value().Show()))
		}
		// utility.DebugPrintf(terminal.Stdout().Height()-1, "pos.X = %v, pos.Y = %v", pos.X, pos.Y)
		vt100.MoveCursorToAndPrint(pos.X, pos.Y, ti.Value().Show())
		childrenCount := len(ti.Children())
		if childrenCount > 0 {
			previousYOffset := 0
			for childIndex, child := range ti.Children() {
				// child position y = parent position y + 1 + child index + previous y offset
				childPosY := pos.Y + childIndex + 1 + previousYOffset
				xOffset, childTreeHeight := child.calculateChildTreeInfo()
				if childTreeHeight != -1 {
					previousYOffset += childTreeHeight
					if childIndex != childrenCount-1 {
						for offset := 1; offset <= childTreeHeight; offset++ {
							vt100.MoveCursorToAndPrint(pos.X, childPosY+offset, string(tab.VL()))
						}
					}
				}
				splitter := tab.LT()
				if childIndex == childrenCount-1 {
					splitter = tab.BL()
				}
				offsetContent := strings.Repeat(string(tab.HL()), xOffset)
				marginContent := strings.Repeat(string(tab.HL()), t.margin)
				childRowContent := fmt.Sprintf("%v%v%v ", string(splitter), offsetContent, marginContent)
				// utility.DebugPrintf(terminal.Stdout().Height()-1, "pos.X = %v, childPosY = %v, childRowContent = |%v|", pos.X, childPosY, childRowContent)
				vt100.MoveCursorToAndPrint(pos.X, childPosY, childRowContent)
				nodePosition[child] = coordinate.Coordinate{
					// child position x = parent position X + splitter width + space width
					X: pos.X + tab.Width() + len(offsetContent) + t.margin*tab.Width() + tab.SpaceWidth(),
					Y: childPosY,
				}
				// utility.DebugPrintf(terminal.Stdout().Height()-1, "child %v position %v, %v", child.Value().Show(), nodePosition[child].X, nodePosition[child].Y)
			}
			// <-terminal.ControlSignal
		}
		return true
	})

	vt100.ClearForegroundColor()
	vt100.ClearBackgroundColor()
}

func bft(root treeInterface, f func(treeInterface) bool) {
	nodeSlice := append(make([]treeInterface, 0), root)
	for len(nodeSlice) != 0 {
		nodeSlice = append(nodeSlice, nodeSlice[0].Children()...)
		if !f(nodeSlice[0]) {
			break
		}
		nodeSlice = nodeSlice[1:]
	}
	fmt.Println()
}
