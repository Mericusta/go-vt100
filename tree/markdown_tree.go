package tree

import (
	"bufio"
	"fmt"
	"io"
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

func NewMarkdownTree(f io.Reader, rootTopic string, markdownDepthSpaceWidth, margin int) Tree {
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

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		switch {
		case rootTopicRegexp.MatchString(scanner.Text()):
			inTopicScope = true
			rootNode = &MarkdownTree{
				tree:  tree{v: &MarkdownTopic{v: rootTopic}},
				depth: 0,
			}
			nodeDepthMap[rootNode] = 0
			currentLineParentNode = rootNode
		case inTopicScope:
			if !topicRegexp.MatchString(scanner.Text()) {
				break
			}
			stringSubmatchSlice := topicRegexp.FindStringSubmatch(scanner.Text())
			if depthIndex >= len(stringSubmatchSlice) {
				panic(fmt.Sprintf("not find sub match DEPTH at |%v|", scanner.Text()))
			}
			depth := len(stringSubmatchSlice[depthIndex]) / markdownDepthSpaceWidth
			if topicIndex >= len(stringSubmatchSlice) {
				panic(fmt.Sprintf("not find sub match TOPIC at |%v|", scanner.Text()))
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
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err().Error())
	}

	return Tree{
		i:      rootNode,
		margin: margin,
		// maxDepth:     treeMaxDepth,
		// maxWidth:     treeMaxWidth,
		nodeDepthMap: nodeDepthMap,
	}
}
