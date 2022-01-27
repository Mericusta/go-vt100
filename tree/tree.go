package tree

import (
	"fmt"
	"go-vt100/canvas"
	"go-vt100/color"
	"go-vt100/coordinate"
	"go-vt100/size"
	"go-vt100/tab"
	"go-vt100/terminal"
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
}

type Tree struct {
	s      size.Size
	i      treeInterface
	fc     color.Color
	bc     color.Color
	margin int
}

func (t Tree) Draw(x, y int) {
	// treeMaxDepth, treeMaxWidth, nodeDepthMap := align(t.i)
	// t.s.Width = t.calculateTreeWidth()
	// t.s.Height = t.calculateTreeHeight()
	// nodeCanvansMap := make(map[treeInterface][]rune)
	// fmt.Printf("treeMaxDepth = %v, treeMaxWidth = %v\n", treeMaxDepth, treeMaxWidth)
	// bft(root, func(t treeInterface) bool {
	// 	// fmt.Printf("%v, node treeInterface:\n", t.Value().Show())
	// 	if len(t.Children()) > 0 {
	// 		nodeCanvansMap[t] = drawNodeCanvas(t, 1)
	// 	}
	// 	return true
	// })

	// mergeNodeCanvas(root, nodeDepthMap, nodeCanvansMap)

	// depthNodeMap := make(map[int][]treeInterface)
	// for node, depth := range nodeDepthMap {
	// 	depthNodeMap[depth] = append(depthNodeMap[depth], node)
	// }

	// printTree(root, treeMaxDepth, treeMaxWidth, depthNodeMap)

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
		vt100.MoveCursorToAndPrint(pos.X, pos.Y, ti.Value().Show())
		childrenCount := len(ti.Children())
		if childrenCount > 0 {
			yOffset := 0
			for childIndex, child := range ti.Children() {
				splitter := tab.LT()
				if childIndex == childrenCount-1 {
					splitter = tab.BL()
				}
				for index := 0; index <= yOffset; index++ {
					vt100.MoveCursorToAndPrint(pos.X, pos.Y+childIndex+1+index, string(tab.VL()))
				}
				// vt100.MoveCursorToAndPrint(pos.X, pos.Y+childIndex+1+yOffset, fmt.Sprintf("%v%v ", string(splitter), strings.Repeat(string(tab.HL()), t.margin)))
				vt100.MoveCursorToAndPrint(pos.X, pos.Y+1+childIndex+yOffset, fmt.Sprintf("%v%v %v", string(splitter), strings.Repeat(string(tab.HL()), t.margin), child.Value().Show()))
				nodePosition[child] = coordinate.Coordinate{
					X: pos.X + tab.Width() + t.margin*tab.Width() + tab.Width(),
					Y: pos.Y + 1 + childIndex + yOffset,
				}
				bft(child, func(ti treeInterface) bool {
					yOffset++
					return true
				})
				<-terminal.ControlSignal
				vt100.SaveScreen()
				fmt.Printf("child %v, y offset %v", child.Value().Show(), yOffset)
				<-terminal.ControlSignal
				vt100.RestoreScreen()
			}
		}
		return false
	})

	vt100.ClearForegroundColor()
	vt100.ClearBackgroundColor()
}

func mergeNodeCanvas(node treeInterface, nodeDepthMap map[treeInterface]int, nodeCanvasMap map[treeInterface][]rune) []rune {
	nodeCanvas := nodeCanvasMap[node]
	if len(node.Children()) == 0 {
		return nil
	}
	for _, child := range node.Children() {
		childCanvas := mergeNodeCanvas(child, nodeDepthMap, nodeCanvasMap)
		if childCanvas == nil {
			fmt.Printf("%v is leaf node\n", child.Value().Show())
			continue
		}
		fmt.Printf("%s", string(childCanvas))
	}

	return nodeCanvas
}

func drawNodeCanvas(n treeInterface, margin int) []rune {
	childrenCount := 0
	rootCellContent := n.Value().Show()
	rootCellWidthStartIndex := 0
	rootCellWidth := len(n.Value().Show())
	childCellMaxWidth := 0
	for _, child := range n.Children() {
		if childCellMaxWidth < len(child.Value().Show()) {
			childCellMaxWidth = len(child.Value().Show())
		}
		childrenCount++
	}
	if childrenCount == 0 {
		return nil
	}

	// |  ┌─ |    |  ┌─ |
	// | ─┼─ | or | ─│  |
	// |  └─ |    |  └─ |
	// space + margin + splitter + margin + space
	branchWidth := tab.Width()*2 + tab.Width()*margin*2 + tab.Width()
	branchHeight := childrenCount
	isOdd := childrenCount%2 == 1
	var splitterByte rune
	if isOdd {
		if childrenCount == 1 {
			splitterByte = tab.HL()
		} else {
			// |                ┌─ Iron magazine|
			// |Steel magazine ─┼─ Steel plate  |
			// |                └─ Copper plate |
			splitterByte = tab.CT()
		}
	} else {
		// |      ┌─ child|
		// |root ─│       |
		// |      └─ child|
		childrenCount++
		splitterByte = tab.RT()
	}

	// root cell + branch cell + child cell + \n
	nodeWidth := rootCellWidth + branchWidth + childCellMaxWidth + tab.Width()
	nodeHeight := childrenCount
	totalPoints := nodeWidth * nodeHeight
	nodeRuneSlice := make([]rune, totalPoints)
	splitterColRelativeIndex := branchWidth/2 + rootCellWidth
	splitterRowRelativeIndex := branchHeight / 2
	for index := 0; index != totalPoints; index++ {
		colRelativeIndex, rowRelativeIndex := canvas.TransformArrayIndexToMatrixCoordinates(index, nodeWidth, nodeHeight)
		switch {
		// root cell
		case colRelativeIndex < splitterColRelativeIndex-margin*tab.Width()-tab.Width():
			if rowRelativeIndex == splitterRowRelativeIndex && colRelativeIndex-rootCellWidthStartIndex < rootCellWidth {
				nodeRuneSlice[index] = rune(rootCellContent[colRelativeIndex-rootCellWidthStartIndex])
			} else {
				nodeRuneSlice[index] = tab.Space()
			}
		// root space
		case colRelativeIndex == splitterColRelativeIndex-margin*tab.Width()-tab.Width():
			nodeRuneSlice[index] = tab.Space()
		// root margin
		case colRelativeIndex < splitterColRelativeIndex:
			if rowRelativeIndex == splitterRowRelativeIndex {
				nodeRuneSlice[index] = tab.HL()
			} else {
				nodeRuneSlice[index] = tab.Space()
			}
		// splitter
		case colRelativeIndex == splitterColRelativeIndex:
			switch {
			case rowRelativeIndex == 0 && childrenCount != 1:
				nodeRuneSlice[index] = tab.TL()
			case rowRelativeIndex == splitterRowRelativeIndex:
				nodeRuneSlice[index] = splitterByte
			case rowRelativeIndex == nodeHeight-1 && childrenCount != 1:
				nodeRuneSlice[index] = tab.BL()
			default:
				nodeRuneSlice[index] = tab.LT()
			}
		// child margin
		case colRelativeIndex <= splitterColRelativeIndex+margin*tab.Width():
			switch {
			case rowRelativeIndex == splitterRowRelativeIndex:
				if isOdd {
					nodeRuneSlice[index] = tab.HL()
				} else {
					nodeRuneSlice[index] = tab.Space()
				}
			default:
				nodeRuneSlice[index] = tab.HL()
			}
		// child space
		case colRelativeIndex == splitterColRelativeIndex+margin*tab.Width()+tab.Width():
			nodeRuneSlice[index] = tab.Space()
		// child cell
		case colRelativeIndex < nodeWidth-1:
			if rowRelativeIndex == splitterRowRelativeIndex && !isOdd {
				nodeRuneSlice[index] = tab.Space()
			} else {
				childIndex := rowRelativeIndex
				if rowRelativeIndex >= splitterRowRelativeIndex && !isOdd {
					childIndex--
				}
				childCellContent := n.Children()[childIndex].Value().Show()
				if cellContentIndex := colRelativeIndex - rootCellWidth - branchWidth; cellContentIndex < len(childCellContent) {
					nodeRuneSlice[index] = rune(childCellContent[cellContentIndex])
				} else {
					nodeRuneSlice[index] = tab.Space()
				}
			}
		// end line
		case colRelativeIndex == nodeWidth-1:
			nodeRuneSlice[index] = tab.EndLine()
		}
	}

	// fmt.Printf("%v", string(nodeRuneSlice))
	return nodeRuneSlice
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
