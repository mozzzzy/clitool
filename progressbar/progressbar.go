package progressbar

/*
 * Module Dependencies
 */

import (
	"strconv"
	"time"

	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/common"
)

/*
 * Types
 */

type Progressbar struct {
	MessageStr string
	Min        float64
	State      *float64
	Max        float64
}

/*
 * Constants and Package Scope Variables
 */

var ProgressbarWidth int = 50

var ProgressbarElementsStr string = "█"
var ProgressbarElementsColorFg termbox.Attribute = termbox.ColorCyan
var ProgressbarElementsColorBg termbox.Attribute = termbox.ColorDefault

var ProgressbarNoElementsStr string = " "
var ProgressbarNoElementsColorFg termbox.Attribute = termbox.ColorDefault
var ProgressbarNoElementsColorBg termbox.Attribute = termbox.ColorDefault

var PaddingMessageAndProgressbarStr string = " "
var PaddingMessageAndProgressbarColorFg termbox.Attribute = termbox.ColorDefault
var PaddingMessageAndProgressbarColorBg termbox.Attribute = termbox.ColorDefault

var MessageStrColorFg termbox.Attribute = termbox.ColorDefault
var MessageStrColorBg termbox.Attribute = termbox.ColorDefault

var PrefixStr string = "|"
var PrefixColorFg termbox.Attribute = termbox.ColorDefault
var PrefixColorBg termbox.Attribute = termbox.ColorDefault

var SuffixStr string = "|"
var SuffixColorFg termbox.Attribute = termbox.ColorDefault
var SuffixColorBg termbox.Attribute = termbox.ColorDefault

var PaddingSuffixAndPercentageStr string = " "
var PaddingSuffixAndPercentageColorFg termbox.Attribute = termbox.ColorDefault
var PaddingSuffixAndPercentageColorBg termbox.Attribute = termbox.ColorDefault

var ParcentageColorFg termbox.Attribute = termbox.ColorDefault
var ParcentageColorBg termbox.Attribute = termbox.ColorDefault

/*
 * Functions
 */

func New(messageStr string, min float64, state *float64, max float64) *Progressbar {
	progressbar := new(Progressbar)
	progressbar.MessageStr = messageStr
	progressbar.Min = min
	progressbar.State = state
	progressbar.Max = max
	return progressbar
}

func (progressbar Progressbar) Print(x int, y int) (int, int) {
	// Print message
	x, y = common.PrintString(
		progressbar.MessageStr, MessageStrColorFg, MessageStrColorBg, x, y)

	// Print padding between message and progress bar
	x, y = common.PrintString(
		PaddingMessageAndProgressbarStr,
		PaddingMessageAndProgressbarColorFg,
		PaddingMessageAndProgressbarColorBg,
		x, y)

	// Print prefix
	x, y = common.PrintString(PrefixStr, PrefixColorFg, PrefixColorBg, x, y)

	// Print progress bar elements
	var oneElementSize float64 = (progressbar.Max - progressbar.Min) / float64(ProgressbarWidth)
	var elementLen int = int(*(progressbar.State) / oneElementSize)
	for count := 0; count < elementLen; count++ {
		x, y = common.PrintString(ProgressbarElementsStr, ProgressbarElementsColorFg, ProgressbarElementsColorBg, x, y)
	}
	// Print progress bar no elements
	for count := elementLen; count < ProgressbarWidth; count++ {
		x, y = common.PrintString(
			ProgressbarNoElementsStr, ProgressbarNoElementsColorFg, ProgressbarNoElementsColorBg, x, y)
	}
	// Print suffix
	x, y = common.PrintString(SuffixStr, SuffixColorFg, SuffixColorBg, x, y)

	// Print padding between suffix and percentage
	x, y = common.PrintString(
		PaddingSuffixAndPercentageStr,
		PaddingSuffixAndPercentageColorFg,
		PaddingSuffixAndPercentageColorBg,
		x, y)

	// Print Percentage
	x, y = common.PrintString(
		strconv.Itoa(int(*(progressbar.State)/(progressbar.Max-progressbar.Min)*100.0))+"%",
		ParcentageColorFg,
		ParcentageColorBg,
		x, y)

	return x, y
}

func (progressbar *Progressbar) Show(x int, y int) (int, int) {
	for *(progressbar.State) <= progressbar.Max {
		progressbar.Print(x, y)
		time.Sleep(100 * time.Millisecond)
	}
	x, y = common.GoNextLine(x, y)
	return x, y
}
