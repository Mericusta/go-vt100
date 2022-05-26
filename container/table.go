package container

import (
	"github.com/Mericusta/go-vt100/core"
)

// Table
type Table struct {
	ContainerContext
	headerObjects  []core.Object
	valueObjects   [][]core.Object
	columnMaxWidth []uint
	rowMaxHeight   []uint
}

func NewTable(headerSlice []string, valueMap [][]core.Drawable) Table {
	t := Table{
		headerObjects:  make([]core.Object, len(headerSlice)),
		valueObjects:   make([][]core.Object, len(valueMap)),
		columnMaxWidth: make([]uint, len(headerSlice)),
		rowMaxHeight:   make([]uint, len(valueMap)),
	}
	// for i, hd := range headerSlice {
	// 	if t.columnMaxWidth[i] == 0 || t.columnMaxWidth[i] < hd.Width() {
	// 		t.columnMaxWidth[i] = hd.Width()
	// 	}
	// 	for j, vdSlice := range valueMap {
	// 		if i > len(vdSlice) {
	// 			continue
	// 		}
	// 		if t.columnMaxWidth[i] == 0 || t.columnMaxWidth[i] < vdSlice[i].Width() {
	// 			t.columnMaxWidth[i] = vdSlice[i].Width()
	// 		}
	// 		for _, vd := range vdSlice {
	// 			if t.rowMaxHeight[j] == 0 || t.rowMaxHeight[j] < vd.Height() {
	// 				t.rowMaxHeight[j] = vd.Height()
	// 			}
	// 		}
	// 	}
	// }
	// terminal.DebugOutput(func() {
	// 	fmt.Printf("t.columnMaxWidth = %v\n", t.columnMaxWidth)
	// 	fmt.Printf("t.rowMaxHeight = %v\n", t.rowMaxHeight)
	// }, nil)
	return t
}
