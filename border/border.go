package border

type BorderType int

const (
	Custom BorderType = iota
	StandardTab
)

// Space ' '
func Space() rune { return ' ' }

func SpaceWidth() int { return 1 }

// EndLine
func EndLine() rune { return '\n' }

func EndLineWidth() uint { return 1 }

func TabWidth() uint { return 1 }

func TabHeight() uint { return 1 }
