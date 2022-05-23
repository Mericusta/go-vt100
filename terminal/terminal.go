package terminal

import (
	"os"
	"os/signal"

	"github.com/Mericusta/go-vt100/core"
)

var ControlSignal chan os.Signal

func init() {
	ControlSignal = make(chan os.Signal)
	signal.Notify(ControlSignal, os.Interrupt)
}

type Terminal interface {
	Width() uint
	Height() uint
}

var terminal Terminal

func Stdout() Terminal {
	return terminal
}

func Destruct() {
	core.ResetAttribute()
	core.ClearScreen()
	core.CursorVisible()
}
