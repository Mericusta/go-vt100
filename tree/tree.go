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
	AppendChildren([]treeInterface)
	Parent() treeInterface
	calculateTreeInfo(int, int, int) (int, int)
}

type tree struct {
	v        valueInterface
	parent   treeInterface
	children []treeInterface
}

func (t *tree) Value() valueInterface {
	return t.v
}

func (t *tree) Children() []treeInterface {
	return t.children
}

func (t *tree) AppendChildren(children []treeInterface) {
	t.children = append(t.children, children...)
}

func (t *tree) Parent() treeInterface {
	return t.parent
}

type Tree struct {
	s            size.Size
	i            treeInterface
	fc           color.Color
	bc           color.Color
	margin       int
	maxDepth     int
	maxWidth     int
	nodeDepthMap map[treeInterface]int
}

func (t Tree) Draw(x, y int, s size.Size) {
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
		if pos.X <= s.Width && pos.Y <= s.Height {
			vt100.MoveCursorToAndPrint(pos.X, pos.Y, ti.Value().Show())
		}
		nodeDepth, hasDepth := t.nodeDepthMap[ti]
		if !hasDepth {
			panic(fmt.Sprintf("node %v not has depth", ti.Value().Show()))
		}
		childrenCount := len(ti.Children())
		if childrenCount > 0 {
			previousYOffset := 0
			for childIndex, child := range ti.Children() {
				// child position y = parent position y + 1 + child index + previous y offset
				childPosY := pos.Y + childIndex + 1 + previousYOffset
				childDepth, hasChildDepth := t.nodeDepthMap[child]
				if !hasChildDepth {
					panic(fmt.Sprintf("node %v child %v not has depth", ti.Value().Show(), child.Value().Show()))
				}
				xOffset, childTreeHeight := child.calculateTreeInfo(nodeDepth, childDepth, t.margin)
				if childTreeHeight > 1 {
					// grandson tree height
					previousYOffset += childTreeHeight - 1
					// print grandson VL except the last child
					if childIndex != childrenCount-1 {
						for offset := 0; offset <= childTreeHeight-1; offset++ {
							if pos.X <= s.Width && childPosY+offset <= s.Height {
								vt100.MoveCursorToAndPrint(pos.X, childPosY+offset, string(tab.VL()))
							}
						}
					}
				}
				splitter := tab.LT()
				if childIndex == childrenCount-1 {
					splitter = tab.BL()
				}
				// utility.DebugPrintf(terminal.Stdout().Height()-1, "pos.X = %v, childPosY = %v, childRowContent = |%v|", pos.X, childPosY, xOffset)
				childRowContent := fmt.Sprintf("%v%v%v ", string(splitter), strings.Repeat(string(tab.HL()), xOffset), strings.Repeat(string(tab.HL()), t.margin))
				if pos.X <= s.Width && childPosY <= s.Height {
					vt100.MoveCursorToAndPrint(pos.X, childPosY, childRowContent)
				}
				nodePosition[child] = coordinate.Coordinate{
					// child position x = parent position X + splitter width + space width
					X: pos.X + tab.Width() + xOffset + t.margin*tab.Width() + tab.SpaceWidth(),
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
