package core

import (
	"os"
	"os/signal"
)

var ControlSignal chan os.Signal

func init() {
	ControlSignal = make(chan os.Signal)
	signal.Notify(ControlSignal, os.Interrupt)
}

type Terminal Unit

var terminal Terminal

func Stdout() Terminal {
	return terminal
}

func Destruct() {
	ResetAttribute()
	ClearScreen()
	CursorVisible()
}

func Context() RenderContext {
	return NewBasicContext(Size{Width: terminal.Width(), Height: terminal.Height()})
}

func DebugOutput(outFunc func(), conditionFunc func() bool) {
	if conditionFunc == nil || conditionFunc() {
		CursorVisible()
		SaveScreen()
		ClearScreen()
		MoveCursorToLine(terminal.Height() / 2)
		outFunc()
		<-ControlSignal
		RestoreScreen()
		CursorInvisible()
	} else {
		panic("here")
	}
}
