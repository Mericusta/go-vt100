package container

import "github.com/Mericusta/go-vt100/core"

// ContainerContext support container size and coordinate to objects while drawing
type ContainerContext struct {
	s core.Size
	c core.Coordinate
}

func NewContainerContext(s core.Size, c core.Coordinate) ContainerContext {
	return ContainerContext{s: s, c: c}
}

func (c ContainerContext) Width() uint {
	return c.s.Width
}

func (c ContainerContext) Height() uint {
	return c.s.Height
}

func (c ContainerContext) Coordinate() core.Coordinate {
	return c.c
}
