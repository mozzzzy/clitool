package spinner

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/common"
)

/*
 * Types
 */

type Spinner struct {
	MessageStr string
	State      int
}

/*
 * Constants and Package Scope Variables
 */

var SpinnerElementStrs []string = []string{"⠙", "⠸", "⠴", "⠦", "⠇", "⠋"}
var SpinnerElementsColorFg termbox.Attribute = termbox.ColorCyan
var SpinnerElementsColorBg termbox.Attribute = termbox.ColorDefault

var PaddingSpinnerAndMessageStr string = " "
var PaddingSpinnerAndMessageColorFg termbox.Attribute = termbox.ColorDefault
var PaddingSpinnerAndMessageColorBg termbox.Attribute = termbox.ColorDefault

var MessageStrColorFg termbox.Attribute = termbox.ColorDefault
var MessageStrColorBg termbox.Attribute = termbox.ColorDefault

var FinishedStr string = "✔"
var FinishedColorFg termbox.Attribute = termbox.ColorGreen
var FinishedColorBg termbox.Attribute = termbox.ColorDefault

/*
 * Functions
 */

func New(messageStr string) *Spinner {
	spinner := new(Spinner)
	spinner.MessageStr = messageStr
	spinner.State = 0
	return spinner
}

func (spinner Spinner) Print(x int, y int) (int, int) {
	// Print a spinner element
	x, y = common.PrintString(
		SpinnerElementStrs[spinner.State], SpinnerElementsColorFg, SpinnerElementsColorBg, x, y)
	// Print padding between spinner and message
	x, y = common.PrintString(
		PaddingSpinnerAndMessageStr, PaddingSpinnerAndMessageColorFg, PaddingSpinnerAndMessageColorBg, x, y)
	// Print message
	x, y = common.PrintString(
		spinner.MessageStr, MessageStrColorFg, MessageStrColorBg, x, y)
	return x, y
}

func (spinner *Spinner) Spin(x int, y int, finished *bool) (int, int) {
	for *finished == false {
		spinner.Print(x, y)
		spinner.State = (spinner.State + 1) % len(SpinnerElementStrs)
		time.Sleep(100 * time.Millisecond)
	}
	common.PrintString(FinishedStr, FinishedColorFg, FinishedColorBg, x, y)
	x, y = common.GoNextLine(x, y)
	return x, y
}
