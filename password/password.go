package password

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

type Password struct {
	q             question.Question
	inputColorBg  color.Color
	inputColorFg  color.Color
	inputStr      string
	inputRune     rune
	answerColorBg color.Color
	answerColorFg color.Color
	answerStr     string
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
	DEFAULT_INPUT_RUNE      rune        = '*'
	DEFAULT_ANSWER_COLOR_BG color.Color = color.Default
	DEFAULT_ANSWER_COLOR_FG color.Color = color.Cyan
)

/*
 * Functions
 */

func New(questionStr string) *Password {
	password := new(Password)
	password.SetQuestion(questionStr)
	password.inputColorBg = DEFAULT_INPUT_COLOR_BG
	password.inputColorFg = DEFAULT_INPUT_COLOR_FG
	password.inputRune = DEFAULT_INPUT_RUNE
	password.answerColorBg = DEFAULT_ANSWER_COLOR_BG
	password.answerColorFg = DEFAULT_ANSWER_COLOR_FG
	return password
}

func (password *Password) GetMinX() int {
	return password.minX
}

func (password *Password) GetMaxX() int {
	return password.maxX
}

func (password *Password) GetMinY() int {
	return password.minY
}

func (password *Password) GetMaxY() int {
	return password.maxY
}

func (password *Password) SetMinX(minX int) {
	// Set minX to inner question for q.Print() method.
	password.q.SetMinX(minX)
	// Set minX
	password.minX = minX
}

func (password *Password) SetMinY(minY int) {
	// Set minY to inner question for q.Print() method.
	password.q.SetMinY(minY)
	// Set minY
	password.minY = minY
}

func (password *Password) SetQuestion(qStr string) {
	q := question.New(qStr)
	password.q = *q
}

func (password *Password) Print() {
	// NOTE if this password has alread printed, its maxX and maxY also has already set.
	//      But the case that its length is changed exists.
	//      So I clear old maxX and maxY here.
	//      These values are calculated again in following logic.
	password.maxX = 0
	password.maxY = 0

	// Set current password
	if !password.resolved {
		password.q.SetSuffix(
			password.inputColorFg,
			password.inputColorBg,
			password.inputStr)
	} else {
		password.q.SetSuffix(password.answerColorFg, password.answerColorBg, password.inputStr)
	}
	// Print question
	password.q.Print()
	password.updateMinMax(password.q.GetMaxX(), password.q.GetMaxY())
}

func (password *Password) Clear() {
	minX := password.GetMinX()
	minY := password.GetMinY()
	maxX := password.GetMaxX()
	maxY := password.GetMaxY()
	for yCursor := minY; yCursor <= maxY; yCursor++ {
		for xCursor := minX; xCursor <= maxX; xCursor++ {
			common.PrintString(" ", color.Default, color.Default, xCursor, yCursor)
		}
	}
	password.maxX = password.GetMinX()
	password.maxY = password.GetMinY()
}

func (password *Password) Inquire() interface{} {
	password.Print()

	// While typed EnterKey by keyboard
mainloop:
	for {
		// Get an event password
		event := common.GetEventKey()
		switch event.Key {
		// If get enter key, break password loop
		case termbox.KeyEnter:
			password.resolved = true
			break mainloop
		case termbox.KeyBackspace:
			fallthrough
		case termbox.KeyBackspace2:
			if len(password.answerStr) > 0 {
				password.answerStr = string(([]rune(password.answerStr))[:len(password.answerStr)-1])
				password.inputStr = string(([]rune(password.inputStr))[:len(password.inputStr)-1])
			}
		default:
			password.answerStr += string(event.Ch)
			password.inputStr += string(password.inputRune)
		}
		password.Clear()
		password.Print()
	}
	password.Clear()
	password.Print()
	return string(password.answerStr)
}

func (password *Password) updateMinMax(x int, y int) {
	// Update maxX
	if password.maxX < x {
		password.maxX = x
	}
	// Update maxY
	if password.maxY < y {
		password.maxY = y
	}
	// Update minX
	if password.minX > x {
		password.minX = x
	}
	// Update minY
	if password.minY > y {
		password.minY = y
	}
}
