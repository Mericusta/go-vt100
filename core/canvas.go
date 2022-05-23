package core

type Canvas struct {
	S               Size
	cursorVisible   bool
	withBoundary    bool
	backgroundColor Color
	layerObjects    []Object
}

func NewCanvas(width, height uint) Canvas {
	return Canvas{
		S: Size{
			Width:  width,
			Height: height,
		},
	}
}

func NewCanvasWithBoundary(width, height uint) Canvas {
	c := NewCanvas(width+2, height+2)
	c.withBoundary = true
	return c
}

func (c *Canvas) Destruct() {
	c.layerObjects = nil
	ResetAttribute()
	ClearScreen()
	CursorVisible()
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

// AddLayerObject
// @param coordinate relative coordinates of the canvas origin
// @param d          something to draw
func (c *Canvas) AddLayerObject(coordinate Coordinate, d Drawable) {
	if c.withBoundary {
		// boundary is not the canvas content, so it can not coincide with boundaries
		c.layerObjects = append(c.layerObjects, NewObject(Coordinate{X: coordinate.X + 1, Y: coordinate.Y + 1}, d))
	} else {
		c.layerObjects = append(c.layerObjects, NewObject(coordinate, d))
	}
}

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
	for y := uint(1); y <= c.S.Height; y++ {
		for x := uint(1); x <= c.S.Width; x++ {
			MoveCursorToAndPrint(x, y, string(Space()))
		}
	}
	ClearBackgroundColor()
}

func (c Canvas) drawBoundary() {
	topLineY := uint(0)
	bottomLineY := c.S.Height - 1
	leftLineX := uint(0)
	rightLineX := c.S.Width - 1
	for x := uint(1); x < rightLineX; x++ {
		MoveCursorToAndPrint(x, topLineY, string(HL()))
		MoveCursorToAndPrint(x, bottomLineY, string(HL()))
	}
	for y := uint(1); y < bottomLineY; y++ {
		MoveCursorToAndPrint(leftLineX, y, string(VL()))
		MoveCursorToAndPrint(rightLineX, y, string(VL()))
	}
	MoveCursorToAndPrint(leftLineX, topLineY, string(TL()))
	MoveCursorToAndPrint(rightLineX, topLineY, string(TR()))
	MoveCursorToAndPrint(leftLineX, bottomLineY, string(BL()))
	MoveCursorToAndPrint(rightLineX, bottomLineY, string(BR()))
}

func TransformMatrixCoordinatesToArrayIndex(x, y, width, height int) int {
	return y*width + x
}

func TransformArrayIndexToMatrixCoordinates(index, width, height int) (int, int) {
	return index % width, index / width
}
