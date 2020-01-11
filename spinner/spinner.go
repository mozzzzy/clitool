package spinner

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/message"
)

/*
 * Types
 */

type Spinner struct {
	message                     message.Message
	elementStrs                 []string
	elementColorBg              color.Color
	elementColorFg              color.Color
	elementCursor               int
	successStr                  string
	successColorBg              color.Color
	successColorFg              color.Color
	failureStr                  string
	failureColorBg              color.Color
	failureColorFg              color.Color
	paddingSpinnerAndMessageStr string
	resolved                    bool
	success                     bool
}

/*
 * Constants and Package Scope Variables
 */

var (
	DEFAULT_MESSAGE_COLOR_BG color.Color = color.Default
	DEFAULT_MESSAGE_COLOR_FG color.Color = color.Default

	DEFAULT_ELEMENT_STRS     []string    = []string{"⠙", "⠸", "⠴", "⠦", "⠇", "⠋"}
	DEFAULT_ELEMENT_COLOR_BG color.Color = color.Default
	DEFAULT_ELEMENT_COLOR_FG color.Color = color.Cyan

	DEFAULT_SUCCESS_STR      string      = "✔"
	DEFAULT_SUCCESS_COLOR_BG color.Color = color.Default
	DEFAULT_SUCCESS_COLOR_FG color.Color = color.Green

	DEFAULT_FAILURE_STR      string      = "✖"
	DEFAULT_FAILURE_COLOR_BG color.Color = color.Default
	DEFAULT_FAILURE_COLOR_FG color.Color = color.Red

	DEFAULT_PADDING_SPINNER_AND_MESSAGE_STR string = " "
)

/*
 * Functions
 */

func New(messageStr string) *Spinner {
	spinner := new(Spinner)

	spinner.message.SetMessage(
		DEFAULT_MESSAGE_COLOR_FG,
		DEFAULT_MESSAGE_COLOR_BG,
		DEFAULT_PADDING_SPINNER_AND_MESSAGE_STR+messageStr)

	spinner.elementStrs = DEFAULT_ELEMENT_STRS
	spinner.message.SetPrefixColor(DEFAULT_ELEMENT_COLOR_FG, DEFAULT_ELEMENT_COLOR_BG)

	spinner.successStr = DEFAULT_SUCCESS_STR
	spinner.successColorBg = DEFAULT_SUCCESS_COLOR_BG
	spinner.successColorFg = DEFAULT_SUCCESS_COLOR_FG

	spinner.failureStr = DEFAULT_FAILURE_STR
	spinner.failureColorBg = DEFAULT_FAILURE_COLOR_BG
	spinner.failureColorFg = DEFAULT_FAILURE_COLOR_FG

	spinner.paddingSpinnerAndMessageStr = DEFAULT_PADDING_SPINNER_AND_MESSAGE_STR

	return spinner
}

func (spinner *Spinner) GetMinX() int {
	return spinner.message.GetMinX()
}

func (spinner *Spinner) GetMaxX() int {
	return spinner.message.GetMaxX()
}

func (spinner *Spinner) GetMinY() int {
	return spinner.message.GetMinY()
}

func (spinner *Spinner) GetMaxY() int {
	return spinner.message.GetMaxY()
}

func (spinner *Spinner) Print() {
	if !spinner.resolved {
		// Set spinner element
		spinner.message.SetPrefixStr(spinner.elementStrs[spinner.elementCursor])
	} else {
		if spinner.success {
			// Set success
			spinner.message.SetPrefix(
				spinner.successColorFg, spinner.successColorBg, spinner.successStr)
		} else {
			// Set failure
			spinner.message.SetPrefix(
				spinner.failureColorFg, spinner.failureColorBg, spinner.failureStr)
		}
	}
	spinner.message.Print()
}

func (spinner *Spinner) SetMinX(minX int) {
	spinner.message.SetMinX(minX)
}

func (spinner *Spinner) SetMinY(minY int) {
	spinner.message.SetMinY(minY)
}

func (spinner *Spinner) Run() {
	for spinner.resolved == false {
		spinner.Print()
		spinner.elementCursor = (spinner.elementCursor + 1) % len(spinner.elementStrs)
		time.Sleep(100 * time.Millisecond)
	}
	spinner.Print()
}

func (spinner *Spinner) Resolve(success bool) {
	spinner.success = success
	spinner.resolved = true
}
