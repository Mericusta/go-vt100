package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/Mericusta/go-vt100/character"
	"github.com/Mericusta/go-vt100/container"
	"github.com/Mericusta/go-vt100/core"
)

type OutputWidget struct {
	canvas           *container.Canvas
	canvasCoordinate core.Coordinate
	cmdChan          chan string
}

func NewOutputWidget(c *container.Canvas, cc core.Coordinate) *OutputWidget {
	return &OutputWidget{
		canvas:           c,
		canvasCoordinate: cc,
		cmdChan:          make(chan string, 8),
	}
}

func (w *OutputWidget) Run(ctx core.RenderContext) {
	w.canvas.Draw(ctx, w.canvasCoordinate)
	w.cmdChan <- "ok"
	for text := range w.cmdChan {
		switch text {
		case "exit":
			w.canvas.Clear()
			close(w.cmdChan)
			return
		default:
			w.canvas.ClearObjects()
			w.canvas.Clear()
			cmdText := container.NewTextarea(text, core.Horizontal)
			w.canvas.AppendObjects(core.Object{
				C: core.Coordinate{},
				D: &cmdText,
			})
			w.canvas.Draw(ctx, w.canvasCoordinate)
			w.cmdChan <- "ok"
		}
	}
}

func (w *OutputWidget) CmdChan() chan string {
	return w.cmdChan
}

type InputWidget struct {
	canvas             *container.Canvas
	canvasCoordinate   core.Coordinate
	cursorCoordinate   core.Coordinate
	cmdInputTextString string
	cmdStringChan      chan string
}

func NewInputWidget(c *container.Canvas, cc core.Coordinate, cmdStringChan chan string) *InputWidget {
	iw := &InputWidget{
		canvas:             c,
		canvasCoordinate:   cc,
		cmdInputTextString: "Input Command:",
		cmdStringChan:      cmdStringChan,
	}
	cmdInputText := container.NewTextarea(iw.cmdInputTextString, core.Horizontal)
	iw.canvas.AppendObjects(core.Object{
		C: core.Coordinate{}, D: &cmdInputText,
	})
	iw.cursorCoordinate = core.Coordinate{
		X: cc.X + int(character.TabWidth()+cmdInputText.Width()),
		Y: cc.Y + int(character.TabHeight()),
	}
	return iw
}

func (w *InputWidget) Run(ctx core.RenderContext) {
	s := <-w.cmdStringChan
	if s == "ok" {
		w.canvas.Draw(ctx, w.canvasCoordinate)
		w.ResetInputCursor()
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			inputText := input.Text()
			if inputText == "exit" {
				return
			}
			w.cmdStringChan <- inputText
			s := <-w.cmdStringChan
			if s == "ok" {
				core.MoveCursorToAndPrint(
					w.cursorCoordinate.X, w.cursorCoordinate.Y,
					strings.Repeat(character.SpaceString(), len(inputText)),
				)
				w.ResetInputCursor()
			}
		}
	}
}

func (w *InputWidget) ResetInputCursor() {
	core.MoveCursorToAndPrint(w.cursorCoordinate.X, w.cursorCoordinate.Y, "")
}

func main() {
	defer core.Destruct()
	core.ClearScreen()
	// core.CursorInvisible()

	outputCanvasWidth := (core.Stdout().Width() / 3) / 2
	outputCanvasHeight := outputCanvasWidth / 4
	outputCanvasSize := core.Size{
		Width:  outputCanvasWidth,
		Height: outputCanvasHeight,
	}
	outputCanvas := container.NewCanvas(outputCanvasSize, true)
	outputCanvasCoordinate := core.Origin()
	outputWidget := NewOutputWidget(&outputCanvas, outputCanvasCoordinate)
	go outputWidget.Run(core.Context())

	inputCanvasSize := core.Size{
		Width:  outputCanvasWidth,
		Height: character.TabHeight(),
	}
	inputCanvas := container.NewCanvas(inputCanvasSize, true)
	inputCanvasCoordinate := core.Coordinate{
		X: core.Origin().X,
		Y: int(outputCanvas.Height() + character.TabHeight()),
	}
	inputWidget := NewInputWidget(&inputCanvas, inputCanvasCoordinate, outputWidget.cmdChan)
	inputWidget.Run(core.Context())

	// TODO: concurrency cursor coordinate control
}
