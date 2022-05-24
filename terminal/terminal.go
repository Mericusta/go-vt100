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

type Terminal core.RenderContext

var terminal Terminal

func Stdout() Terminal {
	return terminal
}

func Destruct() {
	core.ResetAttribute()
	core.ClearScreen()
	core.CursorVisible()
}

func Context() core.RenderContext {
	return terminal
}

func DebugOutput(outFunc func(), conditionFunc func() bool) {
	if conditionFunc == nil || conditionFunc() {
		core.SaveScreen()
		core.ClearScreen()
		core.MoveCursorToLine(terminal.Height() / 2)
		outFunc()
		<-ControlSignal
		core.RestoreScreen()
	}
}
