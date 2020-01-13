package common

/*
 * Module Dependencies
 */

import (
	"os"

	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/color"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func ExitByCtlC() {
	// Poll event
	switch ev := termbox.PollEvent(); ev.Type {
	// Keyboard is typed
	case termbox.EventKey:
		if ev.Key == termbox.KeyCtrlC {
			termbox.Close()
			os.Exit(0)
		}
	}
}

func PrintString(
	str string, clrFg color.Color, clrBg color.Color, x int, y int,
) (int, int) {
	strRunes := []rune(str)
	for _, rne := range strRunes {
		switch rne {
		case '\n':
			x, y = 0, y+1
		default:
			termbox.SetCell(x, y, rne, clrFg.Termbox, clrBg.Termbox)
			x++
		}
	}
	termbox.Flush()
	return x, y
}

func GetKey() (returnKey termbox.Key) {
	// Poll event
	switch ev := termbox.PollEvent(); ev.Type {
	// Keyboard is typed
	case termbox.EventKey:
		if ev.Key == termbox.KeyCtrlC {
			termbox.Close()
			os.Exit(0)
		}
		returnKey = ev.Key
	// Terminal is resized
	case termbox.EventResize:
		// TODO implement if needed
	}
	return returnKey
}

func GetEventKey() (returnEvent termbox.Event) {
	// Poll event
	switch ev := termbox.PollEvent(); ev.Type {
	// Keyboard is typed
	case termbox.EventKey:
		if ev.Key == termbox.KeyCtrlC {
			termbox.Close()
			os.Exit(0)
		}
		returnEvent = ev
	}
	return returnEvent
}

func GoNextLine(x int, y int) (int, int) {
	x = 0
	y++
	return x, y
}
