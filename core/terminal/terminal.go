package terminal

import (
	"os"
	"os/signal"
)

var ControlSignal chan os.Signal

func init() {
	ControlSignal = make(chan os.Signal)
	signal.Notify(ControlSignal, os.Interrupt)
}

type Terminal interface {
	Width() int
	Height() int
}

var terminal Terminal

func Stdout() Terminal {
	return terminal
}
