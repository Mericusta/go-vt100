package container

import (
	"github.com/Mericusta/go-vt100/core"
)

// About Object: Container contains some shapes, the shapes relative coordinate is defined by user, so they need Object.
// About Context: Container can be contained other container, so they must know its container Size and Coordinate.

// ContainerContext support container size and coordinate to objects while drawing
type ContainerContext struct {
	core.BasicContext
}

func NewContainerContext(s core.Size) ContainerContext {
	return ContainerContext{BasicContext: core.NewBasicContext(s)}
}
