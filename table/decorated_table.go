package table

import (
	"fmt"
	"strings"
)

type DecoratedTable struct {
	*AdaptiveTable
	d *TableDecoration
}

type TableDecoration struct {
	WidthPadding  int
	HeightPadding int
}

func NewDecoratedTable(headSlice []string, lineContentSlice [][]string, decoration *TableDecoration) *DecoratedTable {
	t := &DecoratedTable{}
	t.AdaptiveTable = NewAdaptiveTable(headSlice, lineContentSlice)
	t.d = decoration
	for index := range t.colMaxWidthMap {
		t.colMaxWidthMap[index] += t.d.WidthPadding * 2
	}

	builder := strings.Builder{}
	for cellY, contentSlice := range t.contentMap {
		for cellX, content := range contentSlice {
			builder.Reset()
			for index := 0; index != t.d.WidthPadding; index++ {
				builder.WriteRune(' ')
			}
			builder.WriteString(content)
			for index := 0; index != t.d.WidthPadding; index++ {
				builder.WriteRune(' ')
			}
			t.contentMap[cellY][cellX] = builder.String()
		}
	}

	fmt.Printf("t.contentMap = %v\n", t.contentMap)

	return t
}
