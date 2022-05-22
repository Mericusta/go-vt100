package core

type Canvas struct {
	S               Size
	cursorVisible   bool
	withBoundary    bool
	backgroundColor Color
	layerObjects    []Object
}

func NewCanvas(width, height int) Canvas {
	return Canvas{
		S: Size{
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
		CursorVisible()
	} else {
		CursorInvisible()
	}
}

func (c *Canvas) SetBackgroundColor(bc Color) {
	c.backgroundColor = bc
}

func (c *Canvas) AddLayerObject(x, y int, d Drawable) {
	if c.withBoundary {
		c.layerObjects = append(c.layerObjects, NewObject(x+1, y+1, d))
	} else {
		c.layerObjects = append(c.layerObjects, NewObject(x, y, d))
	}
}

// has some bug on vscode terminal with git - bash at sometime
func (c Canvas) Draw() {
	if c.cursorVisible {
		CursorVisible()
	} else {
		CursorInvisible()
	}

	ClearScreen()
	MoveCursorToHome()

	if c.backgroundColor != 0 {
		c.drawBackground()
	}

	if c.withBoundary {
		c.drawBoundary()
	}

	for _, object := range c.layerObjects {
		object.Draw(c.S)
	}

	MoveCursorToHome()
}

func (c *Canvas) Clear() {
	c.layerObjects = nil
	c.Draw()
}

func (c Canvas) drawBackground() {
	SetBackgroundColor(c.backgroundColor)
	for y := 1; y <= c.S.Height; y++ {
		for x := 1; x <= c.S.Width; x++ {
			MoveCursorToAndPrint(x, y, string(Space()))
		}
	}
	ClearBackgroundColor()
}

func (c Canvas) drawBoundary() {
	topLineY := 1
	bottomLineY := c.S.Height
	for x := 2; x < c.S.Width; x++ {
		MoveCursorToAndPrint(x, topLineY, string(HL()))
		MoveCursorToAndPrint(x, bottomLineY, string(HL()))
	}
	leftLineX := 1
	rightLineX := c.S.Width
	for y := 2; y < c.S.Height; y++ {
		MoveCursorToAndPrint(leftLineX, y, string(VL()))
		MoveCursorToAndPrint(rightLineX, y, string(VL()))
	}
	MoveCursorToAndPrint(1, 1, string(TL()))
	MoveCursorToAndPrint(rightLineX, 1, string(TR()))
	MoveCursorToAndPrint(1, bottomLineY, string(BL()))
	MoveCursorToAndPrint(rightLineX, bottomLineY, string(BR()))
}

func TransformMatrixCoordinatesToArrayIndex(x, y, width, height int) int {
	return y*width + x
}

func TransformArrayIndexToMatrixCoordinates(index, width, height int) (int, int) {
	return index % width, index / width
}
