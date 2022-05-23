package tree

import (
	"fmt"
	"strings"

	"github.com/Mericusta/go-vt100/core"
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

func (t *tree) calculateTreeInfo(parentDepth, nodeDepth, margin int) (int, int) {
	xOffset, treeHeight := 0, 0
	bft(t, func(ti treeInterface) bool {
		treeHeight++
		return true
	})

	// xOffset = (depth diff - 1) * (margin + splitter + space)
	// margin = 1
	// (2 - 0 - 1) * (1*1 + 1 + 1)
	// margin = 2
	// (2 - 0 - 1) * (2*1 + 1 + 1)
	// margin = 3
	// (2 - 0 - 1) * (3*1 + 1 + 1)
	xOffset = (nodeDepth - parentDepth - 1) * (margin*core.TabWidth() + core.TabWidth() + core.SpaceWidth())

	return xOffset, treeHeight
}

type Tree struct {
	s            core.Size
	i            treeInterface
	fc           core.Color
	bc           core.Color
	margin       int
	maxDepth     int
	maxWidth     int
	nodeDepthMap map[treeInterface]int
}

func (t Tree) RootNode() treeInterface {
	return t.i
}

func (t Tree) Draw(x, y int, s core.Size) {
	core.SetForegroundColor(t.fc)
	core.SetBackgroundColor(t.bc)

	nodePosition := make(map[treeInterface]core.Coordinate)
	nodePosition[t.i] = core.Coordinate{
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
			core.MoveCursorToAndPrint(pos.X, pos.Y, ti.Value().Show())
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
								core.MoveCursorToAndPrint(pos.X, childPosY+offset, string(core.VL()))
							}
						}
					}
				}
				splitter := core.LT()
				if childIndex == childrenCount-1 {
					splitter = core.BL()
				}
				// utility.DebugPrintf(terminal.Stdout().Height()-1, "pos.X = %v, childPosY = %v, xOffset = %v", pos.X, childPosY, xOffset)
				childRowContent := fmt.Sprintf("%v%v%v ", string(splitter), strings.Repeat(string(core.HL()), xOffset), strings.Repeat(string(core.HL()), t.margin))
				if pos.X <= s.Width && childPosY <= s.Height {
					core.MoveCursorToAndPrint(pos.X, childPosY, childRowContent)
				}
				nodePosition[child] = core.Coordinate{
					// child position x = parent position X + splitter width + space width
					X: pos.X + core.TabWidth() + xOffset + t.margin*core.TabWidth() + core.SpaceWidth(),
					Y: childPosY,
				}
				// utility.DebugPrintf(terminal.Stdout().Height()-1, "child %v position %v, %v", child.Value().Show(), nodePosition[child].X, nodePosition[child].Y)
			}
			// <-terminal.ControlSignal
		}
		return true
	})

	core.ClearForegroundColor()
	core.ClearBackgroundColor()
}

//            A0
//         /  |  \
//       B0   C0  D0
//      /    / |  | \
//    E0   E1 G0  F0 G1
//   / |   | \
// H0 G2   H1 G3
// --------------------
// A0 -> B0 -> C0 -> D0 ->
// E0 -> E1 -> G0 -> F0 ->
// G1 -> H0 -> G2 -> H1 -> G3
func bft(root treeInterface, f func(treeInterface) bool) {
	nodeSlice := append(make([]treeInterface, 0), root)
	for len(nodeSlice) != 0 {
		if !f(nodeSlice[0]) {
			break
		}
		nodeSlice = append(nodeSlice, nodeSlice[0].Children()...)
		nodeSlice = nodeSlice[1:]
	}
	fmt.Println()
}

//            A0
//         /  |  \
//       B0   C0  D0
//      /    / |  | \
//    E0   E1 G0  F0 G1
//   / |   | \
// H0 G2   H1 G3
// --------------------
// A0 -> B0 -> E0 -> H0 ->
// G2 -> C0 -> E1 -> H1 ->
// G3 -> G0 -> D0 -> F0 -> G1
func dft(root treeInterface, f func(treeInterface) bool) {
	nodeSlice := append(make([]treeInterface, 0), root)
	for len(nodeSlice) != 0 {
		if !f(nodeSlice[0]) {
			break
		}
		newHeadSlice := append([]treeInterface{}, nodeSlice[0].Children()...)
		nodeSlice = append(newHeadSlice, nodeSlice[1:]...)
	}
	fmt.Println()
}
