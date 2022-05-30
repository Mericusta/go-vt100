package container

import (
	"github.com/Mericusta/go-vt100/character"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

type Tree struct {
	ContainerContext
	currentObject      core.Object
	subObjects         []core.Object
	borderCanvasObject core.Object
}

func NewTree(d core.Drawable, sds []core.Drawable) Tree {
	t := Tree{currentObject: core.NewObject(core.Coordinate{}, d)}
	subDrawableHeightSlice := make([]uint, len(sds))
	treeWidth := d.Width()
	treeHeight := d.Height()
	for i, sd := range sds {
		t.subObjects = append(t.subObjects, core.NewObject(core.Coordinate{
			X: int(character.TabWidth()*2 + character.SpaceWidth()),
			Y: int(treeHeight),
		}, sd))
		// root node
		// └ sub node
		if treeWidth < character.TabWidth()+character.SpaceWidth()+sd.Width() {
			treeWidth = character.TabWidth() + character.SpaceWidth() + sd.Width()
		}
		subDrawableHeightSlice[i] = sd.Height()
		treeHeight += sd.Height()
	}
	borderCanvas := NewCanvas(core.Size{
		Width:  character.TabWidth()*2 + character.SpaceWidth(),
		Height: treeHeight - d.Height(),
	}, false)

	splitterHeight := uint(0)
	for i, subDrawableHeight := range subDrawableHeightSlice {
		var splitterPoint shape.Point
		var splitterLine shape.Line = shape.NewLine(shape.NewPoint(character.HL()), 1, core.Horizontal)
		var splitterSpace shape.Point = shape.NewPoint(character.Space())
		if i == len(subDrawableHeightSlice)-1 {
			splitterPoint = shape.NewPoint(character.BL())
		} else {
			splitterPoint = shape.NewPoint(character.LT())
		}
		borderCanvas.AppendObjects(
			core.NewObject(core.Coordinate{X: 0, Y: int(splitterHeight)}, splitterPoint),
			core.NewObject(core.Coordinate{X: int(splitterPoint.Width()), Y: int(splitterHeight)}, splitterLine),
			core.NewObject(core.Coordinate{X: int(splitterPoint.Width() + splitterLine.Width()), Y: int(splitterHeight)}, splitterSpace),
		)
		if i > 0 && i != len(subDrawableHeightSlice)-1 {
			if subDrawableHeightSlice[i] > character.TabHeight() {
				splitterLineMaxHeight := splitterPoint.Height()
				if splitterLineMaxHeight < splitterLine.Height() {
					splitterLineMaxHeight = splitterLine.Height()
				}
				if splitterLineMaxHeight < splitterSpace.Height() {
					splitterLineMaxHeight = splitterSpace.Height()
				}
				borderCanvas.AppendObjects(core.NewObject(
					core.Coordinate{X: 0, Y: int(splitterHeight + splitterLineMaxHeight)},
					shape.NewLine(shape.NewPoint(character.VL()), uint(subDrawableHeight-splitterLineMaxHeight), core.Vertical)),
				)
			}
		}
		splitterHeight += subDrawableHeight
	}
	t.borderCanvasObject = core.NewObject(core.Coordinate{
		X: 0, Y: int(d.Height()),
	}, &borderCanvas)
	t.BasicContext = core.NewBasicContext(core.Size{
		Width:  treeWidth,
		Height: treeHeight,
	})
	return t
}

func (t *Tree) Draw(ctx core.RenderContext, c core.Coordinate) {
	t.SetCoordinate(c)
	coincidenceCtx, has := ctx.CoincidenceCheck(t)
	if !has {
		return
	}

	t.currentObject.D.Draw(coincidenceCtx, core.Coordinate{
		X: c.X + t.currentObject.C.X,
		Y: c.Y + t.currentObject.C.Y,
	})
	t.borderCanvasObject.D.Draw(coincidenceCtx, core.Coordinate{
		X: c.X + t.borderCanvasObject.C.X,
		Y: c.Y + t.borderCanvasObject.C.Y,
	})
	for _, so := range t.subObjects {
		so.D.Draw(coincidenceCtx, core.Coordinate{
			X: c.X + so.C.X,
			Y: c.Y + so.C.Y,
		})
	}
}
