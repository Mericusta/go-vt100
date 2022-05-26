package container

import (
	"github.com/Mericusta/go-vt100/core"
)

// About Object: Container contains some shapes, the shapes relative coordinate is defined by user, so they need Object.
// About Context: Container can be contained other container, so they must know its container Size and Coordinate.

// ContainerContext support container size and coordinate to objects while drawing
type ContainerContext struct {
	c core.Coordinate
	s core.Size
}

func NewContainerContext(c core.Coordinate, s core.Size) ContainerContext {
	return ContainerContext{c: c, s: s}
}

func (c ContainerContext) Coordinate() core.Coordinate {
	return c.c
}

func (c ContainerContext) Width() uint {
	return c.s.Width
}

func (c ContainerContext) Height() uint {
	return c.s.Height
}
