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

func Origin() Coordinate {
	return Coordinate{X: 1, Y: 1}
}

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

func DebugOutput(outFunc func(), conditionFunc ...func() bool) {
	if conditionFunc == nil || func() bool {
		total := true
		for _, f := range conditionFunc {
			total = total && f()
		}
		return total
	}() {
		CursorVisible()
		SaveScreen()
		ClearScreen()
		MoveCursorToLine(int(terminal.Height() / 2))
		outFunc()
		<-ControlSignal
		RestoreScreen()
		CursorInvisible()
	} else {
		panic("here")
	}
}
