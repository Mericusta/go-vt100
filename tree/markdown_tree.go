package tree

import (
	"fmt"
	"go-vt100/utility"
	"regexp"
)

type MarkdownTree struct {
	tree
}

func NewMarkdownTree(filename, rootTopic string, depthSpaceWidth int) Tree {
	rootTopicRegexp := regexp.MustCompile(fmt.Sprintf(`^\s*-\s+%v\s*$`, rootTopic))
	topicRegexp := regexp.MustCompile(`^(?P<DEPTH>\s+)-\s+(?P<TOPIC>.*)$`)
	depthIndex := topicRegexp.SubexpIndex("DEPTH")
	topicIndex := topicRegexp.SubexpIndex("TOPIC")
	if depthIndex == -1 || topicIndex == -1 {
		panic("submatch DEPTH or TOPIC index is -1")
	}
	inTopicScope := false

	utility.ReadFileLineOneByOne(filename, func(s string) bool {
		switch {
		case rootTopicRegexp.MatchString(s):
			inTopicScope = true
			fmt.Printf("root topic |%v|\n", rootTopic)
			return true
		case inTopicScope:
			if !topicRegexp.MatchString(s) {
				fmt.Printf("not match |%v|\n", s)
				return false
			}
			stringSubmatchSlice := topicRegexp.FindStringSubmatch(s)
			if depthIndex >= len(stringSubmatchSlice) {
				fmt.Printf("not find sub match DEPTH at |%v|", s)
				return false
			}
			depth := len(stringSubmatchSlice[depthIndex]) / depthSpaceWidth
			if topicIndex >= len(stringSubmatchSlice) {
				fmt.Printf("not find sub match TOPIC at |%v|", s)
				return false
			}
			fmt.Printf("topic |%v|, depth %v\n", stringSubmatchSlice[topicIndex], depth)
			return true
		default:
			return true
		}
	})

	return Tree{
		// i:            nil,
		// margin:       margin,
		// maxDepth:     treeMaxDepth,
		// maxWidth:     treeMaxWidth,
		// nodeDepthMap: nodeDepthMap,
	}
}
