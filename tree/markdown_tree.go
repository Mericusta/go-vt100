package tree

import (
	"fmt"
	"go-vt100/utility"
	"regexp"
	"strings"
)

type MarkdownTopic struct {
	v string
}

func (m MarkdownTopic) Show() string {
	return strings.TrimSpace(m.v)
}

type MarkdownTree struct {
	tree
	depth int
}

func NewMarkdownTree(filename, rootTopic string, markdownDepthSpaceWidth, margin int) Tree {
	rootTopicRegexp := regexp.MustCompile(fmt.Sprintf(`^\s*-\s+%v\s*$`, rootTopic))
	topicRegexp := regexp.MustCompile(`^(?P<DEPTH>\s+)-\s+(?P<TOPIC>.*)$`)
	depthIndex := topicRegexp.SubexpIndex("DEPTH")
	topicIndex := topicRegexp.SubexpIndex("TOPIC")
	if depthIndex == -1 || topicIndex == -1 {
		panic("submatch DEPTH or TOPIC index is -1")
	}
	inTopicScope := false

	// rootNode := &MarkdownTree{tree: tree{v: &MarkdownTopic{v: rootTopic}}, depth: 0}
	var rootNode *MarkdownTree
	var currentLineParentNode treeInterface
	nodeDepthMap := make(map[treeInterface]int)

	utility.ReadFileLineOneByOne(filename, func(s string) bool {
		switch {
		case rootTopicRegexp.MatchString(s):
			inTopicScope = true
			rootNode = &MarkdownTree{
				tree:  tree{v: &MarkdownTopic{v: rootTopic}},
				depth: 0,
			}
			nodeDepthMap[rootNode] = 0
			currentLineParentNode = rootNode
			return true
		case inTopicScope:
			if !topicRegexp.MatchString(s) {
				return false
			}
			stringSubmatchSlice := topicRegexp.FindStringSubmatch(s)
			if depthIndex >= len(stringSubmatchSlice) {
				panic(fmt.Sprintf("not find sub match DEPTH at |%v|", s))
			}
			depth := len(stringSubmatchSlice[depthIndex]) / markdownDepthSpaceWidth
			if topicIndex >= len(stringSubmatchSlice) {
				panic(fmt.Sprintf("not find sub match TOPIC at |%v|", s))
			}

			switch {
			case currentLineParentNode.(*MarkdownTree).depth+1 == depth:
			case currentLineParentNode.(*MarkdownTree).depth+1 > depth:
				for currentLineParentNode.(*MarkdownTree).depth+1 > depth {
					currentLineParentNode = currentLineParentNode.Parent()
				}
			case currentLineParentNode.(*MarkdownTree).depth+1 < depth:
				parentNodeChildren := currentLineParentNode.Children()
				if childrenCount := len(parentNodeChildren); childrenCount > 0 {
					currentLineParentNode = parentNodeChildren[childrenCount-1]
				}
			}

			// fmt.Printf("topic |%v|, depth %v\n", stringSubmatchSlice[topicIndex], depth)
			currentLineNode := &MarkdownTree{
				tree: tree{
					v:      MarkdownTopic{v: stringSubmatchSlice[topicIndex]},
					parent: currentLineParentNode,
				},
				depth: depth,
			}
			nodeDepthMap[currentLineNode] = depth
			currentLineParentNode.AppendChildren([]treeInterface{currentLineNode})
			return true
		default:
			return true
		}
	})

	return Tree{
		i:      rootNode,
		margin: margin,
		// maxDepth:     treeMaxDepth,
		// maxWidth:     treeMaxWidth,
		nodeDepthMap: nodeDepthMap,
	}
}

func (t *MarkdownTree) calculateTreeInfo(parentDepth, nodeDepth, margin int) (int, int) {
	xOffset, treeHeight := 0, 0
	bft(t, func(ti treeInterface) bool {
		treeHeight++
		return true
	})
	return xOffset, treeHeight
}
