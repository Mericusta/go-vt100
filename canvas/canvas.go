package canvas

import (
	"go-vt100/size"
	"go-vt100/tab"
	"go-vt100/vt100"
)

type Canvas struct {
	S              size.Size
	withBoundary   bool
	backgroundRune rune
}

func NewCanvas(width, height int) Canvas {
	return Canvas{
		S: size.Size{
			Width:  width,
			Height: height,
		},
	}
}

func NewCanvasWithBoundary(width, height int) Canvas {
	c := NewCanvas(width+2, height+2)
	c.withBoundary = true
	return c
}

func (c *Canvas) SetBackground(r rune) {
	c.backgroundRune = r
}

// has some bug on vscode terminal with git - bash at sometime
func (c Canvas) Draw() {
	vt100.ClearScreen()

	if c.backgroundRune != 0 {
		for y := 1; y <= c.S.Height; y++ {
			for x := 1; x <= c.S.Width; x++ {
				vt100.MoveCursorToAndPrint(x, y, string(c.backgroundRune))
			}
		}
	}

	if c.withBoundary {
		topLineY := 1
		bottomLineY := c.S.Height
		for x := 2; x < c.S.Width; x++ {
			vt100.MoveCursorToAndPrint(x, topLineY, string(tab.HL()))
			vt100.MoveCursorToAndPrint(x, bottomLineY, string(tab.HL()))
		}
		leftLineX := 1
		rightLineX := c.S.Width
		for y := 2; y < c.S.Height; y++ {
			vt100.MoveCursorToAndPrint(leftLineX, y, string(tab.VL()))
			vt100.MoveCursorToAndPrint(rightLineX, y, string(tab.VL()))
		}
		vt100.MoveCursorToAndPrint(1, 1, string(tab.TL()))
		vt100.MoveCursorToAndPrint(rightLineX, 1, string(tab.TR()))
		vt100.MoveCursorToAndPrint(1, bottomLineY, string(tab.BL()))
		vt100.MoveCursorToAndPrint(rightLineX, bottomLineY, string(tab.BR()))
	}

	vt100.MoveCursorToLine(c.S.Height)
}

func TransformMatrixCoordinatesToArrayIndex(x, y, width, height int) int {
	return y*width + x
}

func TransformArrayIndexToMatrixCoordinates(index, width, height int) (int, int) {
	return index % width, index / width
}
