package tree

import (
	"fmt"
)

type Tree struct {
	v        string
	parent   *Tree
	children []*Tree
}

func NewTree() *Tree {
	nodeA0 := &Tree{v: "A0: Steel magazine"}
	nodeB0 := &Tree{v: "B0: Iron magazine"}
	nodeC0 := &Tree{v: "C0: Steel plate"}
	nodeD0 := &Tree{v: "D0: Copper plate"}
	nodeE0 := &Tree{v: "E0: Iron plate"}
	nodeE1 := &Tree{v: "E1: Iron plate"}
	nodeG0 := &Tree{v: "G0: Coal"}
	nodeF0 := &Tree{v: "F0: Copper ore"}
	nodeG1 := &Tree{v: "G1: Coal"}
	nodeH0 := &Tree{v: "H0: Iron ore"}
	nodeG2 := &Tree{v: "G2: Coal"}
	nodeH1 := &Tree{v: "H1: Iron ore"}
	nodeG3 := &Tree{v: "G3: Coal"}

	nodeA0.children = append(nodeA0.children, nodeB0, nodeC0, nodeD0)    // steel magazine
	nodeB0.parent, nodeC0.parent, nodeD0.parent = nodeA0, nodeA0, nodeA0 // steel magazine
	nodeB0.children = append(nodeB0.children, nodeE0)                    // iron magazine
	nodeE0.parent = nodeB0                                               // iron magazine
	nodeC0.children = append(nodeC0.children, nodeE1, nodeG0)            // steel plate
	nodeE1.parent, nodeG0.parent = nodeC0, nodeC0                        // steel plate
	nodeD0.children = append(nodeD0.children, nodeF0, nodeG1)            // copper plate
	nodeF0.parent, nodeG1.parent = nodeD0, nodeD0                        // copper plate
	nodeE0.children = append(nodeE0.children, nodeH0, nodeG2)            // iron plate
	nodeH0.parent, nodeG2.parent = nodeE0, nodeE0                        // iron plate
	nodeE1.children = append(nodeE1.children, nodeH1, nodeG3)            // iron plate
	nodeH1.parent, nodeG3.parent = nodeE1, nodeE1                        // iron plate

	//            A0                      A0
	//         /  |  \                 /  |  \
	//       B0   C0  D0             B0   C0  \
	//      /    / |  | \           /    / |   \
	//    E0   E1 G0  F0 G1 ->    E0   E1  |    D0
	//   / |   | \               / |  / |  |   / |
	// H0 G2   H1 G3           H0 G2 H1 G3 G0 F0 G1
	// ----------------------------
	// align to the bottom rules:
	// rule 1: same element
	// rule 2: no-subnode element
	// rule 3: the element which its subnode satisfied rule2

	return nodeA0
}

func Draw(root *Tree) {
	align(root)
}

func align(root *Tree) {
	treeMaxDepth := 0
	nodeDepthMap := make(map[*Tree]int)
	nodeDepthMap[root] = 0
	noSubNodeSlice := make([]*Tree, 0)
	bft(root, func(t *Tree) bool {
		if len(t.children) == 0 {
			noSubNodeSlice = append(noSubNodeSlice, t)
		} else {
			for _, subNode := range t.children {
				if treeMaxDepth < nodeDepthMap[t]+1 {
					treeMaxDepth = nodeDepthMap[t] + 1
				}
				nodeDepthMap[subNode] = nodeDepthMap[t] + 1
			}
		}
		fmt.Printf("node %v, %p, nodeDepthMap = %v\n", t.v, t, nodeDepthMap)
		return true
	})

	// leaf node falldown to max depth
	falldownNodes := make([]*Tree, 0, len(noSubNodeSlice))
	for _, noSubNode := range noSubNodeSlice {
		if depth := nodeDepthMap[noSubNode]; depth < treeMaxDepth {
			nodeDepthMap[noSubNode] = treeMaxDepth
			falldownNodes = append(falldownNodes, noSubNode)
		}
	}

	// non-leaf node search and falldown
	for len(falldownNodes) != 0 {
		if parentNodeDepth, has := nodeDepthMap[falldownNodes[0].parent]; has && parentNodeDepth+1 < nodeDepthMap[falldownNodes[0]] {
			parentFallDown := true
			for _, siblingNode := range falldownNodes[0].parent.children {
				if siblingNode == falldownNodes[0] {
					continue
				} else {
					if parentNodeDepth+1 == nodeDepthMap[siblingNode] {
						parentFallDown = false
						break
					}
				}
			}
			if parentFallDown {
				nodeDepthMap[falldownNodes[0].parent] = nodeDepthMap[falldownNodes[0]] - 1
				falldownNodes = append(falldownNodes, falldownNodes[0].parent)
			}
		}
		falldownNodes = falldownNodes[1:]
	}

	fmt.Printf("treeMaxDepth = %v\n", treeMaxDepth)
	for node, depth := range nodeDepthMap {
		fmt.Printf("depth: %v, node: %v\n", depth, node.v)
	}
}

func bft(root *Tree, f func(*Tree) bool) {
	nodeSlice := append(make([]*Tree, 0), root)
	for len(nodeSlice) != 0 {
		nodeSlice = append(nodeSlice, nodeSlice[0].children...)
		if !f(nodeSlice[0]) {
			break
		}
		nodeSlice = nodeSlice[1:]
	}
	fmt.Println()
}
