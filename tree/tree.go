package tree

import "fmt"

type Tree struct {
	v       string
	subTree []*Tree
}

func NewTree() *Tree {
	nodeA := &Tree{v: "A: Steel magazine"}
	nodeB := &Tree{v: "B: Iron magazine"}
	nodeC := &Tree{v: "C: Steel plate"}
	nodeD := &Tree{v: "D: Copper plate"}
	nodeE := &Tree{v: "E: Iron plate"}
	nodeF := &Tree{v: "F: Copper ore"}
	nodeH := &Tree{v: "H: Iron ore"}
	nodeG := &Tree{v: "G: Coal"}

	nodeA.subTree = append(nodeA.subTree, nodeB, nodeC, nodeD) // steel magazine
	nodeB.subTree = append(nodeB.subTree, nodeE)               // iron magazine
	nodeC.subTree = append(nodeC.subTree, nodeE, nodeG)        // steel plate
	nodeD.subTree = append(nodeD.subTree, nodeF, nodeG)        // copper plate
	nodeE.subTree = append(nodeE.subTree, nodeH, nodeG)        // iron plate

	//       A               A
	//      /|\            / | \
	//     B C D          B  C  \
	//    / /| |\        /  /|   \
	//   E E G F G ->   E  E |    D
	//  /| |\          /| /| |   /|
	// H G H G        H GH G G  F G
	// ----------------------------
	// align to the bottom rules:
	// rule 1: same element
	// rule 2: no-subnode element
	// rule 3: the element which its subnode satisfied rule2

	return nodeA
}

func Draw(root *Tree) {
	treeMaxLevel := 0
	nodeMaxLevelMap := make(map[*Tree]int)
	nodeMaxLevelMap[root] = 0
	bft(root, func(t *Tree) bool {
		for _, subNode := range t.subTree {
			if level, hasLevel := nodeMaxLevelMap[subNode]; hasLevel && level < nodeMaxLevelMap[t]+1 {
				fmt.Printf("node %v, %p level align to bottom level %v\n", subNode.v, subNode, nodeMaxLevelMap[t]+1)
			}
			if treeMaxLevel < nodeMaxLevelMap[t]+1 {
				treeMaxLevel = nodeMaxLevelMap[t] + 1
			}
			nodeMaxLevelMap[subNode] = nodeMaxLevelMap[t] + 1
		}
		fmt.Printf("node: %v, %p, nodeMaxLevelMap = %v\n", t.v, t, nodeMaxLevelMap)
		return true
	})

	fmt.Printf("treeMaxLevel = %v\n", treeMaxLevel)
	for node, level := range nodeMaxLevelMap {
		fmt.Printf("level: %v, node: %v\n", level, node.v)
	}
}

func bft(root *Tree, f func(*Tree) bool) {
	nodeSlice := append(make([]*Tree, 0), root)
	for len(nodeSlice) != 0 {
		nodeSlice = append(nodeSlice, nodeSlice[0].subTree...)
		if !f(nodeSlice[0]) {
			break
		}
		nodeSlice = nodeSlice[1:]
	}
	fmt.Println()
}
