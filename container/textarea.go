package container

import (
	"strings"

	"github.com/Mericusta/go-vt100/border"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

// Textarea is a Canvas with multi line text
type Textarea struct {
	ContainerContext
	canvas Canvas
}

func NewTextarea(s string, d core.Direction) Textarea {
	lineTextSlice := strings.Split(s, string(border.EndLine()))
	objects := make([]core.Object, 0, len(s))
	var canvas Canvas
	if d == core.Horizontal {
		var maxWidth, height int
		for _, line := range lineTextSlice {
			lineWidth, maxHeight := 0, 0
			for _, r := range line {
				p := shape.NewPoint(r)
				objects = append(
					objects,
					core.Object{
						C: core.Coordinate{X: lineWidth, Y: height},
						D: p,
					},
				)
				lineWidth += int(p.Width())
				if maxHeight == 0 || maxHeight < int(p.Height()) {
					maxHeight = int(p.Height())
				}
			}
			if maxWidth == 0 || maxWidth < lineWidth {
				maxWidth = lineWidth
			}
			height += maxHeight
		}
		canvas = NewCanvas(core.Size{Width: uint(maxWidth), Height: uint(height)})
		canvas.AppendObjects(objects...)
	} else if d == core.Vertical {
		var maxHeight, width int
		for _, line := range lineTextSlice {
			lineHeight, maxWidth := 0, 0
			for _, r := range line {
				p := shape.NewPoint(r)
				objects = append(
					objects,
					core.Object{
						C: core.Coordinate{X: width, Y: lineHeight},
						D: p,
					},
				)
				lineHeight += int(p.Height())
				if maxWidth == 0 || maxWidth < int(p.Width()) {
					maxWidth = int(p.Width())
				}
			}
			if maxHeight == 0 || maxHeight < lineHeight {
				maxHeight = lineHeight
			}
			width += maxWidth
		}
		canvas = NewCanvas(core.Size{Width: uint(width), Height: uint(maxHeight)})
		canvas.AppendObjects(objects...)
	} else {
		panic("not supported direction")
	}
	return Textarea{
		ContainerContext: canvas.ContainerContext,
		canvas:           canvas,
	}
}

func (t *Textarea) Draw(ctx core.RenderContext, c core.Coordinate) {
	t.ContainerContext.c = c
	t.canvas.Draw(ctx, c)
}
