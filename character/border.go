package character

type BorderType int

const (
	Custom BorderType = iota
	StandardTab
)

// Space ' '
func Space() rune { return ' ' }

func SpaceWidth() uint { return 1 }

// EndLine
func EndLine() rune { return '\n' }

func EndLineWidth() uint { return 1 }
