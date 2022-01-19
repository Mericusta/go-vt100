package terminal

type Terminal interface {
	Width() int
	Height() int
}

var terminal Terminal

func Stdout() Terminal {
	return terminal
}
