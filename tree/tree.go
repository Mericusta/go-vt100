package tree

import (
	"fmt"
	"go-vt100/canvas"
	"go-vt100/tab"
)

type Value interface {
	Show() string
}

type Tree interface {
	Value() Value
	Children() []Tree
	Parent() Tree
}

func Draw(root Tree) {
	treeMaxDepth, treeMaxWidth, nodeDepthMap := align(root)
	fmt.Printf("treeMaxDepth = %v, treeMaxWidth = %v\n", treeMaxDepth, treeMaxWidth)
	bft(root, func(t Tree) bool {
		if nodeDepthMap[t] < 1 {
			fmt.Printf("%v, node tree:\n", t.Value().Show())
			if len(t.Children()) > 0 {
				printNode(t, 1)
			}
		}
		return true
	})
}

func printNode(n Tree, margin int) []rune {
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
	branchWidth := tab.Width()*2 + tab.Width()*margin*2 + tab.Width() + tab.Width()
	branchHeight := childrenCount
	isOdd := childrenCount%2 == 1
	var splitterRune rune
	if isOdd {
		// |                ┌─ Iron magazine|
		// |Steel magazine ─┼─ Steel plate  |
		// |                └─ Copper plate |
		splitterRune = tab.CT()
	} else {
		// |      ┌─ child|
		// |root ─│       |
		// |      └─ child|
		childrenCount++
		splitterRune = tab.RT()
	}

	nodeWidth := rootCellWidth + branchWidth + childCellMaxWidth
	nodeHeight := childrenCount
	totalPoints := nodeWidth * nodeHeight
	nodeRuneSlice := make([]rune, totalPoints)
	splitterColRelativeIndex := branchWidth/2 + rootCellWidth
	splitterRowRelativeIndex := branchHeight / 2
	fmt.Printf("totalPoints = %v\n", totalPoints)
	fmt.Printf("splitterColRelativeIndex = %v\n", splitterColRelativeIndex)
	fmt.Printf("splitterRowRelativeIndex = %v\n", splitterRowRelativeIndex)
	for index := 0; index != totalPoints; index++ {
		colRelativeIndex, rowRelativeIndex := canvas.TransformArrayIndexToMatrixCoordinates(index, nodeWidth, nodeHeight)
		switch {
		case colRelativeIndex < splitterColRelativeIndex-tab.Width()-tab.Width():
			if rowRelativeIndex == splitterRowRelativeIndex && colRelativeIndex-rootCellWidthStartIndex < rootCellWidth {
				nodeRuneSlice[index] = rune(rootCellContent[colRelativeIndex-rootCellWidthStartIndex])
			} else {
				nodeRuneSlice[index] = tab.Space()
			}
		case colRelativeIndex == splitterColRelativeIndex-1:
			if rowRelativeIndex == splitterRowRelativeIndex {
				nodeRuneSlice[index] = tab.HL()
			} else {
				nodeRuneSlice[index] = tab.Space()
			}
		case colRelativeIndex == splitterColRelativeIndex:
			switch {
			case rowRelativeIndex == 0:
				nodeRuneSlice[index] = tab.TL()
			case rowRelativeIndex == splitterRowRelativeIndex:
				nodeRuneSlice[index] = splitterRune
			case rowRelativeIndex == nodeHeight-1:
				nodeRuneSlice[index] = tab.BL()
			default:
				nodeRuneSlice[index] = tab.LT()
			}
		case colRelativeIndex == splitterColRelativeIndex+1:
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
		case colRelativeIndex == nodeWidth-1:
			nodeRuneSlice[index] = tab.EndLine()
		case colRelativeIndex > splitterColRelativeIndex+1:
			if rowRelativeIndex == splitterRowRelativeIndex && isOdd {
				nodeRuneSlice[index] = tab.Space()
			} else {
				childIndex := rowRelativeIndex
				if rowRelativeIndex > splitterRowRelativeIndex && !isOdd {
					childIndex--
				}
				childCellContent := n.Children()[childIndex].Value().Show()
				fmt.Printf("child index = %v, child content = %v\n", childIndex, childCellContent)
				if cellContentIndex := colRelativeIndex - splitterColRelativeIndex; cellContentIndex < len(childCellContent) {
					nodeRuneSlice[index] = rune(childCellContent[cellContentIndex])
				} else {
					nodeRuneSlice[index] = tab.Space()
				}
			}
		}
	}

	fmt.Printf("%v", string(nodeRuneSlice))
	return nodeRuneSlice
}

func bft(root Tree, f func(Tree) bool) {
	nodeSlice := append(make([]Tree, 0), root)
	for len(nodeSlice) != 0 {
		nodeSlice = append(nodeSlice, nodeSlice[0].Children()...)
		if !f(nodeSlice[0]) {
			break
		}
		nodeSlice = nodeSlice[1:]
	}
	fmt.Println()
}
