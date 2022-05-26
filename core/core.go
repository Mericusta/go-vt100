package core

import (
	"math"
)

type Unit interface {
	Width() uint
	Height() uint
}

type RenderContext interface {
	Unit
	Coordinate() Coordinate
	// CoincidenceCheck
	// @param  RenderContext
	// @return RenderContext coincidence context
	//         bool          is coincidence
	CoincidenceCheck(RenderContext) (RenderContext, bool)
}

type BasicContext struct {
	c Coordinate
	s Size
}

func NewBasicContext(s Size) BasicContext {
	return BasicContext{s: s}
}

func (ctx *BasicContext) SetCoordinate(c Coordinate) {
	ctx.c = c
}

func (c BasicContext) Size() Size {
	return c.s
}

func (c BasicContext) Width() uint {
	return c.s.Width
}

func (c BasicContext) Height() uint {
	return c.s.Height
}

func (c BasicContext) Coordinate() Coordinate {
	return c.c
}

// ┌──────────┐
// │ Terminal │
// │     ┌───────────┐
// │     │ Container │
// └─────│           │
//       │ * Point   │
//       └───────────┘
// Point is in Container but not in terminal.
// Thats why Container also need coincidence check.
// @param ctx parent container ctx
func (c BasicContext) CoincidenceCheck(ctx RenderContext) (RenderContext, bool) {
	newCtx := BasicContext{}
	startAbsX := c.c.X
	endAbsX := startAbsX + int(c.s.Width)
	newCtx.s.Width = c.s.Width
	// DebugOutput(func() {
	// 	fmt.Printf("startAbsX = %v\n", startAbsX)
	// 	fmt.Printf("endAbsX = %v\n", endAbsX)
	// 	fmt.Printf("ctx.Coordinate().X = %v\n", ctx.Coordinate().X)
	// 	fmt.Printf("ctx.Coordinate().X+int(ctx.Width()) = %v\n", ctx.Coordinate().X+int(ctx.Width()))
	// }, nil)
	if endAbsX < ctx.Coordinate().X || ctx.Coordinate().X+int(ctx.Width()) < startAbsX {
		// outer left || outer right
		return newCtx, false
	} else {
		newCtx.c.X = int(math.Max(float64(startAbsX), float64(ctx.Coordinate().X)))
		if newCtx.c.X < 0 {
			newCtx.c.X = 0
		}
	}
	newCtx.s.Width = uint(math.Max(float64(endAbsX), float64(ctx.Coordinate().X+int(ctx.Width()))) - float64(newCtx.c.X))

	startAbsY := c.c.Y
	endAbsY := startAbsY + int(c.s.Height)
	newCtx.s.Height = c.s.Height
	if startAbsY < ctx.Coordinate().Y || endAbsY > ctx.Coordinate().Y+int(ctx.Height()) {
		// outer top || outer bottom
		return newCtx, false
	} else {
		newCtx.c.Y = int(math.Max(float64(startAbsY), float64(ctx.Coordinate().Y)))
		if newCtx.c.Y < 0 {
			newCtx.c.Y = 0
		}
	}
	newCtx.s.Height = uint(math.Max(float64(endAbsY), float64(ctx.Coordinate().Y+int(ctx.Height()))) - float64(newCtx.c.Y))
	return newCtx, true
}

type Drawable interface {
	Unit
	// Draw
	// @param RenderContext parent container's Size and Coordinate
	// @param Coordinate absolute coordinate
	Draw(RenderContext, Coordinate)
}

// Object hold drawable and its relative coordinate of parent
type Object struct {
	C Coordinate // relative coordinate
	D Drawable
}

func NewObject(c Coordinate, d Drawable) Object {
	return Object{c, d}
}
