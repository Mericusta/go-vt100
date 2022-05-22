package core

// ┌────┬────┐
// │    │    │
// ├────┼────┤
// │    │    │
// └────┴────┘

// CT '┼'
func CT() rune { return '┼' }

// HL '─'
func HL() rune { return '─' }

// VL '│'
func VL() rune { return '│' }

// TL '┌'
func TL() rune { return '┌' }

// TR '┐'
func TR() rune { return '┐' }

// BL '└'
func BL() rune { return '└' }

// BR '┘'
func BR() rune { return '┘' }

// TT '┬'
func TT() rune { return '┬' }

// BT '┴'
func BT() rune { return '┴' }

// LT '├'
func LT() rune { return '├' }

// RT '┤'
func RT() rune { return '┤' }

// Space ' '
func Space() rune { return ' ' }

func SpaceWidth() int { return 1 }

// EndLine
func EndLine() rune { return '\n' }

func EndLineWidth() int { return 1 }

func Width() int { return 1 }

func Height() int { return 1 }
