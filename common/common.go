package common

/*
 * Module Dependencies
 */

import (
	"github.com/nsf/termbox-go"
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

func PrintString(
	str string, clrFg termbox.Attribute, clrBg termbox.Attribute, x int, y int,
) (int, int) {
	strRunes := []rune(str)
	for _, rne := range strRunes {
		termbox.SetCell(x, y, rne, clrFg, clrBg)
		x++
	}
	termbox.Flush()
	return x, y
}

func GetKey() (returnKey termbox.Key) {
	// Poll event
	switch ev := termbox.PollEvent(); ev.Type {
	// Keyboard is typed
	case termbox.EventKey:
		returnKey = ev.Key
	// Terminal is resized
	case termbox.EventResize:
		// TODO implement if needed
	}
	return returnKey
}

func GoNextLine(x int, y int) (int, int) {
	x = 0
	y++
	return x, y
}
