package canvas

import (
	"github.com/Mericusta/go-vt100/color"
	"github.com/Mericusta/go-vt100/object"
	"github.com/Mericusta/go-vt100/size"
	"github.com/Mericusta/go-vt100/tab"
	"github.com/Mericusta/go-vt100/vt100"
)

type Canvas struct {
	S               size.Size
	cursorVisible   bool
	withBoundary    bool
	backgroundColor color.Color
	layerObjects    []object.Object
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

func (c *Canvas) SetCursorVisible(visible bool) {
	c.cursorVisible = visible
	if c.cursorVisible {
		vt100.CursorVisible()
	} else {
		vt100.CursorInvisible()
	}
}

func (c *Canvas) SetBackgroundColor(bc color.Color) {
	c.backgroundColor = bc
}

func (c *Canvas) AddLayerObject(x, y int, d object.Drawable) {
	if c.withBoundary {
		c.layerObjects = append(c.layerObjects, object.NewObject(x+1, y+1, d))
	} else {
		c.layerObjects = append(c.layerObjects, object.NewObject(x, y, d))
	}
}

// has some bug on vscode terminal with git - bash at sometime
func (c Canvas) Draw() {
	if c.cursorVisible {
		vt100.CursorVisible()
	} else {
		vt100.CursorInvisible()
	}

	vt100.ClearScreen()
	vt100.MoveCursorToHome()

	if c.backgroundColor != 0 {
		c.drawBackground()
	}

	if c.withBoundary {
		c.drawBoundary()
	}

	for _, object := range c.layerObjects {
		object.Draw(c.S)
	}

	vt100.MoveCursorToHome()
}

func (c *Canvas) Clear() {
	c.layerObjects = nil
	c.Draw()
}

func (c Canvas) drawBackground() {
	vt100.SetBackgroundColor(c.backgroundColor)
	for y := 1; y <= c.S.Height; y++ {
		for x := 1; x <= c.S.Width; x++ {
			vt100.MoveCursorToAndPrint(x, y, string(tab.Space()))
		}
	}
	vt100.ClearBackgroundColor()
}

func (c Canvas) drawBoundary() {
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

func TransformMatrixCoordinatesToArrayIndex(x, y, width, height int) int {
	return y*width + x
}

func TransformArrayIndexToMatrixCoordinates(index, width, height int) (int, int) {
	return index % width, index / width
}
