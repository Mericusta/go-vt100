package tree

import "go-vt100/tab"

type FactorioMaterial struct {
	v string
}

func (m FactorioMaterial) Show() string {
	return m.v
}

type FactorioTree struct {
	v        valueInterface
	tag      string
	parent   treeInterface
	children []treeInterface
}

func NewFactorioTree(margin int) Tree {
	// load data
	nodeA0 := &FactorioTree{v: &FactorioMaterial{v: "Steel magazine"}, tag: "A0"}
	nodeB0 := &FactorioTree{v: &FactorioMaterial{v: "Iron magazine"}, tag: "B0"}
	nodeC0 := &FactorioTree{v: &FactorioMaterial{v: "Steel plate"}, tag: "C0"}
	nodeD0 := &FactorioTree{v: &FactorioMaterial{v: "Copper plate"}, tag: "D0"}
	nodeE0 := &FactorioTree{v: &FactorioMaterial{v: "Iron plate"}, tag: "E0"}
	nodeE1 := &FactorioTree{v: &FactorioMaterial{v: "Iron plate"}, tag: "E1"}
	nodeG0 := &FactorioTree{v: &FactorioMaterial{v: "Coal"}, tag: "G0"}
	nodeF0 := &FactorioTree{v: &FactorioMaterial{v: "Copper ore"}, tag: "F0"}
	nodeG1 := &FactorioTree{v: &FactorioMaterial{v: "Coal"}, tag: "G1"}
	nodeH0 := &FactorioTree{v: &FactorioMaterial{v: "Iron ore"}, tag: "H0"}
	nodeG2 := &FactorioTree{v: &FactorioMaterial{v: "Coal"}, tag: "G2"}
	nodeH1 := &FactorioTree{v: &FactorioMaterial{v: "Iron ore"}, tag: "H1"}
	nodeG3 := &FactorioTree{v: &FactorioMaterial{v: "Coal"}, tag: "G3"}
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

	// align tree
	treeMaxDepth, treeMaxWidth, nodeDepthMap := nodeA0.align()

	return Tree{
		i:            nodeA0,
		margin:       margin,
		maxDepth:     treeMaxDepth,
		maxWidth:     treeMaxWidth,
		nodeDepthMap: nodeDepthMap,
	}
}

func (t *FactorioTree) Value() valueInterface {
	return t.v
}

func (t *FactorioTree) Children() []treeInterface {
	return t.children
}

func (t *FactorioTree) Parent() treeInterface {
	return t.parent
}

func (t *FactorioTree) calculateTreeWidth() int {
	return -1
}

func (t *FactorioTree) calculateTreeHeight() int {
	return -1
}

func (t *FactorioTree) align() (int, int, map[treeInterface]int) {
	treeMaxDepth := 0
	nodeDepthMap := make(map[treeInterface]int)
	nodeDepthMap[t] = 0
	noSubNodeSlice := make([]treeInterface, 0)
	bft(t, func(ti treeInterface) bool {
		if len(ti.Children()) == 0 {
			noSubNodeSlice = append(noSubNodeSlice, ti)
		} else {
			for _, subNode := range ti.Children() {
				if treeMaxDepth < nodeDepthMap[ti]+1 {
					treeMaxDepth = nodeDepthMap[ti] + 1
				}
				nodeDepthMap[subNode] = nodeDepthMap[ti] + 1
			}
		}
		return true
	})

	// leaf node falldown to max depth
	treeMaxWidth := len(noSubNodeSlice)
	falldownNodes := make([]treeInterface, 0, treeMaxWidth)
	for _, noSubNode := range noSubNodeSlice {
		if depth := nodeDepthMap[noSubNode]; depth < treeMaxDepth {
			nodeDepthMap[noSubNode] = treeMaxDepth
			falldownNodes = append(falldownNodes, noSubNode)
		}
	}

	// non-leaf node search and falldown
	for len(falldownNodes) != 0 {
		if parentNodeDepth, has := nodeDepthMap[falldownNodes[0].Parent()]; has && parentNodeDepth+1 < nodeDepthMap[falldownNodes[0]] {
			parentFallDown := true
			for _, siblingNode := range falldownNodes[0].Parent().Children() {
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
				nodeDepthMap[falldownNodes[0].Parent()] = nodeDepthMap[falldownNodes[0]] - 1
				falldownNodes = append(falldownNodes, falldownNodes[0].Parent())
			}
		}
		falldownNodes = falldownNodes[1:]
	}

	// fmt.Printf("treeInterface max depth: %v\n", treeMaxDepth)
	// fmt.Printf("treeInterface max width: %v\n", treeMaxWidth)
	// bft(t, func(t Tree) bool {
	// 	fmt.Printf("depth %v, treeInterface node %v, %v\n", nodeDepthMap[t], t.(*FactorioTree).tag, t.Value().Show())
	// 	return true
	// })

	return treeMaxDepth, treeMaxWidth, nodeDepthMap
}

// |                  ┌─ H0|
// |    ┌─ B0 ─── E0 ─┤    |
// |    │             └─ G2|
// |    │             ┌─ H1|
// |    │      ┌─ E1 ─┤    |
// |A0 ─┼─ C0 ─┤      └─ G3|
// |    │      └──────── G0|
// |    │             ┌─ F0|
// |    └──────── D0 ─┤    |
// |                  └─ G1|

// margin = 1
// |A0         |
// |├─ B0      | -> B0 tree height
// |│  └─ E0   | -> B0 tree height
// |│     ├─ H0| -> B0 tree height
// |│     └─ G2| -> B0 tree height
// |├─ C0      |
// |│  ├─ E1   |
// |│  │  ├─ H1|
// |│  │  └─ G3|
// |│  └──── G0|
// |└──── D0   |
// |      ├─ F0|
// |      └─ G1|

// margin = 2
// |A0            |
// |├── B0        | -> B0 tree height
// |│   └── E0    | -> B0 tree height
// |│       ├── H0| -> B0 tree height
// |│       └── G2| -> B0 tree height
// |├── C0        |
// |│   ├── E1    |
// |│   │   ├── H1|
// |│   │   └── G3|
// |│   └────── G0|
// |└────── D0    |
// |        ├── F0|
// |        └── G1|

// margin = 3
// |A0               |
// |├─── B0          | -> B0 tree height
// |│    └─── E0     | -> B0 tree height
// |│         ├─── H0| -> B0 tree height
// |│         └─── G2| -> B0 tree height
// |├─── C0          |
// |│    ├─── E1     |
// |│    │    ├─── H1|
// |│    │    └─── G3|
// |│    └──────── G0|
// |└──────── D0     |
// |          ├─── F0|
// |          └─── G1|

func (t *FactorioTree) calculateTreeInfo(parentDepth, nodeDepth, margin int) (int, int) {
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
	xOffset = (nodeDepth - parentDepth - 1) * (margin*tab.Width() + tab.Width() + tab.SpaceWidth())

	return xOffset, treeHeight
}
