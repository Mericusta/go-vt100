package matrix

import (
	"go-vt100/object"
	"go-vt100/size"
)

type Matrix struct {
	object.Object
	S size.Size
}

func NewMatrix() {

}
