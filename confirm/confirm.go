package confirm

/*
 * Module Dependencies
 */

import (
	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/common"
	"github.com/mozzzzy/clitool/question"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

var PaddingQuestionAndYesNo string = " "
var PaddingQuestionAndYesNoColorFg termbox.Attribute = termbox.ColorDefault
var PaddingQuestionAndYesNoColorBg termbox.Attribute = termbox.ColorDefault

var PrefixStr string = "("
var PrefixColorFg termbox.Attribute = termbox.ColorDefault
var PrefixColorBg termbox.Attribute = termbox.ColorDefault

var YesRune rune = 'y'
var YesColorFg termbox.Attribute = termbox.ColorDefault
var YesColorBg termbox.Attribute = termbox.ColorDefault

var SeparatorStr string = "/"
var SeparatorColorFg termbox.Attribute = termbox.ColorDefault
var SeparatorColorBg termbox.Attribute = termbox.ColorDefault

var NoRune rune = 'N'
var NoColorFg termbox.Attribute = termbox.ColorDefault
var NoColorBg termbox.Attribute = termbox.ColorDefault

var SuffixStr string = ")"
var SuffixColorFg termbox.Attribute = termbox.ColorDefault
var SuffixColorBg termbox.Attribute = termbox.ColorDefault

/*
 * Functions
 */

func Inquire(qStr string, x int, y int) (bool, int, int) {
	// Save start point
	startX := x
	startY := y
	// Print question
	q := question.New(qStr)
	x, y = q.PrintQuestion(x, y)
	// Print padding
	x, y = common.PrintString(
		PaddingQuestionAndYesNo,
		PaddingQuestionAndYesNoColorFg,
		PaddingQuestionAndYesNoColorBg,
		x, y)
	// Print candidates
	x, y = common.PrintString(PrefixStr, PrefixColorFg, PrefixColorBg, x, y)
	x, y = common.PrintString(string(YesRune), YesColorFg, YesColorBg, x, y)
	x, y = common.PrintString(SeparatorStr, SeparatorColorFg, SeparatorColorBg, x, y)
	x, y = common.PrintString(string(NoRune), NoColorFg, NoColorBg, x, y)
	x, y = common.PrintString(SuffixStr, SuffixColorFg, SuffixColorBg, x, y)

	var answerBool bool
	// While typed YesRune or NoRune by keyboard
mainloop:
	for {
		// Get an event input
		event := common.GetEventKey()
		switch event.Ch {
		case YesRune:
			answerBool = true
			break mainloop
		case NoRune:
			answerBool = false
			break mainloop
		}
	}
	// Save end point
	endX := x
	// Go back to start point
	x, y = startX, startY
	// Print answer
	if answerBool {
		q.AnswerStr = string(YesRune)
	} else {
		q.AnswerStr = string(NoRune)
	}
	x, y = q.PrintAnswer(x, y)
	for x <= endX {
		x, y = common.PrintString(" ", termbox.ColorDefault, termbox.ColorDefault, x, y)
	}
	// Go to next line
	x, y = common.GoNextLine(x, y)

	return answerBool, x, y
}
