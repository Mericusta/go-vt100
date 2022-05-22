package tree

type FactorioMaterial struct {
	v string
}

func (m FactorioMaterial) Show() string {
	return m.v
}

type FactorioTree struct {
	tree
	tag string
}

func NewFactorioTree(margin int) Tree {
	// load data
	nodeA0 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Steel magazine"}}, tag: "A0"}
	nodeB0 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Iron magazine"}}, tag: "B0"}
	nodeC0 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Steel plate"}}, tag: "C0"}
	nodeD0 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Copper plate"}}, tag: "D0"}
	nodeE0 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Iron plate"}}, tag: "E0"}
	nodeE1 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Iron plate"}}, tag: "E1"}
	nodeG0 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Coal"}}, tag: "G0"}
	nodeF0 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Copper ore"}}, tag: "F0"}
	nodeG1 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Coal"}}, tag: "G1"}
	nodeH0 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Iron ore"}}, tag: "H0"}
	nodeG2 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Coal"}}, tag: "G2"}
	nodeH1 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Iron ore"}}, tag: "H1"}
	nodeG3 := &FactorioTree{tree: tree{v: &FactorioMaterial{v: "Coal"}}, tag: "G3"}
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

	// right align tree
	treeMaxDepth, treeMaxWidth, nodeDepthMap := rightAlign(nodeA0)

	// middle align tree
	middleAlign(nodeA0)

	return Tree{
		i:            nodeA0,
		margin:       margin,
		maxDepth:     treeMaxDepth,
		maxWidth:     treeMaxWidth,
		nodeDepthMap: nodeDepthMap,
	}
}

func NewFactorioTreeWithRootNode(rootNode treeInterface, margin int) Tree {
	// right align tree
	treeMaxDepth, treeMaxWidth, nodeDepthMap := rightAlign(rootNode)

	return Tree{
		i:            rootNode,
		margin:       margin,
		maxDepth:     treeMaxDepth,
		maxWidth:     treeMaxWidth,
		nodeDepthMap: nodeDepthMap,
	}
}

// leafAlign, also called 'Horizontal Right Alignment'
// default horizontal alignment is 'Horizontal Left Alignment'
// tree node view changes as follows:
// |           A0       |    |           A0       |
// |        /  |  \     |    |        /  |  \     |
// |      B0   C0  D0   |    |      B0   C0  \    |
// |     /    / |  | \  | -> |     /    / |   \   |
// |   E0   E1 G0  F0 G1|    |   E0   E1  |    D0 |
// |  / |   | \         |    |  / |  / |  |   / | |
// |H0 G2   H1 G3       |    |H0 G2 H1 G3 G0 F0 G1|
// ------------------------------------------------
// render view changes as follows:
// |A0         |    |A0         |
// |├─ B0      |    |├─ B0      |
// |│  └─ E0   |    |│  └─ E0   |
// |│     ├─ H0|    |│     ├─ H0|
// |│     └─ G2|    |│     └─ G2|
// |├─ C0      |    |├─ C0      |
// |│  ├─ E1   | -> |│  ├─ E1   |
// |│  │  ├─ H1|    |│  │  ├─ H1|
// |│  │  └─ G3|    |│  │  └─ G3|
// |│  └─ G0   |    |│  └──── G0|
// |└─ D0      |    |└──── D0   |
// |   ├─ F0   |    |      ├─ F0|
// |   └─ G1   |    |      └─ G1|
// align to the bottom rules:
// rule 1: same element
// rule 2: no-subnode element
// rule 3: the element which its subnode satisfied rule2
func rightAlign(t treeInterface) (int, int, map[treeInterface]int) {
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

// middleAlign, also called 'Vertical Center Alignment'
// default vertical alignment is 'Vertical Top Alignment'
// |A0         |    |                  ┌─ H0|
// |├─ B0      |    |    ┌─ B0 ─── E0 ─┤    |
// |│  └─ E0   |    |    │             └─ G2|
// |│     ├─ H0|    |    │             ┌─ H1|
// |│     └─ G2|    |    │      ┌─ E1 ─┤    |
// |├─ C0      |    |A0 ─┼─ C0 ─┤      └─ G3|
// |│  ├─ E1   | -> |    │      └─ G0       |
// |│  │  ├─ H1|    |    │      ┌─ F0       |
// |│  │  └─ G3|    |    └─ D0 ─┤           |
// |│  └─ G0   |    |           └─ G1       |
// |└─ D0      |    |                       |
// |   ├─ F0   |    |                       |
// |   └─ G1   |    |                       |
func middleAlign(t treeInterface) {
	treeWidthMap := make(map[treeInterface]int)
	dft(t, func(ti treeInterface) bool {
		if len(ti.Children()) == 0 {
			treeWidthMap[ti] = 0
			for parent := ti.Parent(); parent != nil; parent = parent.Parent() {
				treeWidthMap[parent]++
			}
		}
		return true
	})

	// for node := range treeWidthMap {
	// 	if
	// }

	// utility.DebugPrintf("treeMaxWidth = %v", treeWidthMap[t])
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
