package question

/*
 * Module Dependencies
 */

import (
	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/common"
)

/*
 * Types
 */

type Question struct {
	QuestionStr string
	AnswerStr   string
}

/*
 * Constants and Package Scope Variables
 */

var PrefixStr string = "? "
var PrefixColorFg termbox.Attribute = termbox.ColorGreen
var PrefixColorBg termbox.Attribute = termbox.ColorDefault

var QuestionColorFg termbox.Attribute = termbox.ColorDefault
var QuestionColorBg termbox.Attribute = termbox.ColorDefault

var PaddingQuestionAndAnswerStr string = " "
var PaddingQuestionAndAnswerColorFg termbox.Attribute = termbox.ColorDefault
var PaddingQuestionAndAnswerColorBg termbox.Attribute = termbox.ColorDefault

var AnswerColorFg termbox.Attribute = termbox.ColorCyan
var AnswerColorBg termbox.Attribute = termbox.ColorDefault

/*
 * Functions
 */

func New(questionStr string) *Question {
	question := new(Question)
	return question
}

func (question Question) PrintQuestion(x int, y int) (int, int) {
	// Print prefix
	x, y = common.PrintString(PrefixStr, PrefixColorFg, PrefixColorBg, x, y)
	// Print question
	x, y = common.PrintString(
		question.QuestionStr, QuestionColorFg, QuestionColorBg, x, y)
	return x, y
}

func (question Question) PrintAnswer(x int, y int) (int, int) {
	// Print question
	x, y = question.PrintQuestion(x, y)
	// Print padding between question and answer
	x, y = common.PrintString(
		PaddingQuestionAndAnswerStr,
		PaddingQuestionAndAnswerColorFg,
		PaddingQuestionAndAnswerColorBg,
		x,
		y,
	)
	// Print answer
	x, y = common.PrintString(
		question.AnswerStr, AnswerColorFg, AnswerColorBg, x, y)
	return x, y
}
