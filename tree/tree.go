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
			fmt.Printf("%v", t.Value().Show())
			if len(t.Children()) > 0 {
				printBranch(t, 1)
			}
		}
		return true
	})
}

func printBranch(n Tree) [][]rune {
	childrenCount := len(n.Children())
	if childrenCount > 0 {
		return nil
	}
	// 2*space width + tab width
	branchWidth := ((tab.Width() * 2) + tab.Width())
	isOdd := childrenCount%2 == 1
	branchHeight := childrenCount
	var splitterRune rune
	if isOdd {
		// |b┌s|
		// |s┼s|
		// |b└s|
		splitterRune = tab.CT()
	} else {
		// |b┌s|
		// |s┤b|
		// |b└s|
		childrenCount++
		splitterRune = tab.RT()
	}

	totalPoints := branchWidth * branchHeight
	branchesRuneSlice := make([]rune, totalPoints)
	splitterColRelativeIndex := branchWidth / 2
	splitterRowRelativeIndex := branchHeight / 2
	// splitterIndex := canvas.TransformMatrixCoordinatesToArrayIndex(splitterColRelativeIndex, splitterRowRelativeIndex, branchWidth, branchHeight)
	// branchesRuneSlice[splitterIndex] = splitterRune
	for index := 0; index != totalPoints; index++ {
		colRelativeIndex, rowRelativeIndex := canvas.TransformArrayIndexToMatrixCoordinates(index, branchWidth, branchHeight)
		switch {
		case colRelativeIndex < splitterColRelativeIndex:
			branchesRuneSlice[index] = tab.Space()
		}
	}
	return nil
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
