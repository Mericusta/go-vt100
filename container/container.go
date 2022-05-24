package container

import "github.com/Mericusta/go-vt100/core"

type ContainerContext struct {
	s core.Size
	c core.Coordinate
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
