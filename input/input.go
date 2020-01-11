package input

/*
 * Module Dependencies
 */

import (
	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/common"
	"github.com/mozzzzy/clitool/question"
)

/*
 * Types
 */

type Input struct {
	q             question.Question
	inputColorBg  color.Color
	inputColorFg  color.Color
	answerColorBg color.Color
	answerColorFg color.Color
	inputStr      string
	minX          int
	maxX          int
	minY          int
	maxY          int
	resolved      bool
}

/*
 * Constants and Package Scope Variables
 */

var (
	DEFAULT_INPUT_COLOR_BG  color.Color = color.Default
	DEFAULT_INPUT_COLOR_FG  color.Color = color.Default
	DEFAULT_ANSWER_COLOR_BG color.Color = color.Default
	DEFAULT_ANSWER_COLOR_FG color.Color = color.Cyan
)

/*
 * Functions
 */

func New(questionStr string) *Input {
	input := new(Input)
	input.SetQuestion(questionStr)
	input.inputColorBg = DEFAULT_INPUT_COLOR_BG
	input.inputColorFg = DEFAULT_INPUT_COLOR_FG
	input.answerColorBg = DEFAULT_ANSWER_COLOR_BG
	input.answerColorFg = DEFAULT_ANSWER_COLOR_FG
	return input
}

func (input *Input) GetMinX() int {
	return input.minX
}

func (input *Input) GetMaxX() int {
	return input.maxX
}

func (input *Input) GetMinY() int {
	return input.minY
}

func (input *Input) GetMaxY() int {
	return input.maxY
}

func (input *Input) SetMinX(minX int) {
	// Set minX to inner question for q.Print() method.
	input.q.SetMinX(minX)
	// Set minX
	input.minX = minX
}

func (input *Input) SetMinY(minY int) {
	// Set minY to inner question for q.Print() method.
	input.q.SetMinY(minY)
	// Set minY
	input.minY = minY
}

func (input *Input) SetQuestion(qStr string) {
	q := question.New(qStr)
	input.q = *q
}

func (input *Input) Print() {
	// NOTE if this input has alread printed, its maxX and maxY also has already set.
	//      But the case that its length is changed exists.
	//      So I clear old maxX and maxY here.
	//      These values are calculated again in following logic.
	input.maxX = 0
	input.maxY = 0

	// Set current input
	if !input.resolved {
		input.q.SetSuffix(input.inputColorFg, input.inputColorBg, input.inputStr)
	} else {
		input.q.SetSuffix(input.answerColorFg, input.answerColorBg, input.inputStr)
	}
	// Print question
	input.q.Print()
	input.updateMinMax(input.q.GetMaxX(), input.q.GetMaxY())
}

func (input *Input) Clear() {
	minX := input.GetMinX()
	minY := input.GetMinY()
	maxX := input.GetMaxX()
	maxY := input.GetMaxY()
	for yCursor := minY; yCursor <= maxY; yCursor++ {
		for xCursor := minX; xCursor <= maxX; xCursor++ {
			common.PrintString(" ", color.Default, color.Default, xCursor, yCursor)
		}
	}
	input.maxX = input.GetMinX()
	input.maxY = input.GetMinY()
}

func (input *Input) Inquire() interface{} {
	input.Print()

	// While typed EnterKey by keyboard
mainloop:
	for {
		// Get an event input
		event := common.GetEventKey()
		switch event.Key {
		// If get enter key, break input loop
		case termbox.KeyEnter:
			input.resolved = true
			break mainloop
		case termbox.KeyBackspace:
			fallthrough
		case termbox.KeyBackspace2:
			if len(input.inputStr) > 0 {
				input.inputStr = string(([]rune(input.inputStr))[:len(input.inputStr)-1])
			}
		default:
			input.inputStr += string(event.Ch)
		}
		input.Clear()
		input.Print()
	}
	input.Clear()
	input.Print()
	return string(input.inputStr)
}

func (input *Input) updateMinMax(x int, y int) {
	// Update maxX
	if input.maxX < x {
		input.maxX = x
	}
	// Update maxY
	if input.maxY < y {
		input.maxY = y
	}
	// Update minX
	if input.minX > x {
		input.minX = x
	}
	// Update minY
	if input.minY > y {
		input.minY = y
	}
}
