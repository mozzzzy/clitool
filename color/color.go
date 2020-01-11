package color

/*
 * Module Dependencies
 */

import (
	"github.com/nsf/termbox-go"
)

/*
 * Types
 */
type Color struct {
	Termbox termbox.Attribute
	StdFg   string
	StdBg   string
}

/*
 * Constants and Package Scope Variables
 */

var (
	Default Color = Color{
		Termbox: termbox.ColorDefault,
		StdFg:   "\x1b[0m",
		StdBg:   "\x1b[0m",
	}
	Black Color = Color{
		Termbox: termbox.ColorBlack,
		StdFg:   "\x1b[30m",
		StdBg:   "\x1b[40m",
	}
	Red Color = Color{
		Termbox: termbox.ColorRed,
		StdFg:   "\x1b[31m",
		StdBg:   "\x1b[41m",
	}
	Green Color = Color{
		Termbox: termbox.ColorGreen,
		StdFg:   "\x1b[32m",
		StdBg:   "\x1b[42m",
	}
	Yellow Color = Color{
		Termbox: termbox.ColorYellow,
		StdFg:   "\x1b[33m",
		StdBg:   "\x1b[43m",
	}
	Blue Color = Color{
		Termbox: termbox.ColorBlue,
		StdFg:   "\x1b[34m",
		StdBg:   "\x1b[44m",
	}
	Magenta Color = Color{
		Termbox: termbox.ColorMagenta,
		StdFg:   "\x1b[35m",
		StdBg:   "\x1b[45m",
	}
	Cyan Color = Color{
		Termbox: termbox.ColorCyan,
		StdFg:   "\x1b[36m",
		StdBg:   "\x1b[46m",
	}
	White Color = Color{
		Termbox: termbox.ColorWhite,
		StdFg:   "\x1b[37m",
		StdBg:   "\x1b[47m",
	}
)

/*
 * Functions
 */
