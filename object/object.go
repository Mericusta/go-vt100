package object

import (
	"go-vt100/coordinate"
	"go-vt100/size"
)

type Object struct {
	C coordinate.Coordinate
	S size.Size
}
