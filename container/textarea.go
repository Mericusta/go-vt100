package container

import (
	"strings"

	"github.com/Mericusta/go-vt100/character"
	"github.com/Mericusta/go-vt100/core"
	"github.com/Mericusta/go-vt100/shape"
)

// Textarea is a collection of Points
type Textarea struct {
	ContainerContext
	objects [][]core.Object
}

func NewTextarea(s string, d core.Direction) Textarea {
	lineTextSlice := strings.Split(s, string(character.EndLine()))
	t := Textarea{
		objects: make([][]core.Object, 0, len(s)),
	}
	if d == core.Horizontal {
		var maxWidth, height int
		for _, line := range lineTextSlice {
			lineWidth, maxHeight := 0, 0
			lineObjects := make([]core.Object, 0, len(line))
			for _, r := range line {
				p := shape.NewPoint(r)
				lineObjects = append(
					lineObjects,
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
			t.objects = append(t.objects, lineObjects)
		}
		t.BasicContext = core.NewBasicContext(core.Size{
			Width:  uint(maxWidth),
			Height: uint(height),
		})
	} else if d == core.Vertical {
		var maxHeight, width int
		for _, line := range lineTextSlice {
			lineHeight, maxWidth := 0, 0
			lineObjects := make([]core.Object, 0, len(line))
			for _, r := range line {
				p := shape.NewPoint(r)
				lineObjects = append(
					lineObjects,
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
			t.objects = append(t.objects, lineObjects)
		}
		t.BasicContext = core.NewBasicContext(core.Size{
			Width:  uint(width),
			Height: uint(maxHeight),
		})
	} else {
		panic("not supported direction")
	}
	return t
}

func (t *Textarea) Draw(ctx core.RenderContext, c core.Coordinate) {
	t.SetCoordinate(c)
	coincidenceCtx, has := ctx.CoincidenceCheck(t)
	if !has {
		return
	}

	for _, os := range t.objects {
		for _, o := range os {
			// <-core.ControlSignal
			o.D.Draw(coincidenceCtx, core.Coordinate{X: c.X + o.C.X, Y: c.Y + o.C.Y})
		}
	}
}
