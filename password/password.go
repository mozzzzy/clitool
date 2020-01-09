package password

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

var PaddingQuestionAndAnswer string = " "
var PaddingQuestionAndAnswerColorFg termbox.Attribute = termbox.ColorDefault
var PaddingQuestionAndAnswerColorBg termbox.Attribute = termbox.ColorDefault

var PasswordRune rune = '*'
var PasswordColorFg termbox.Attribute = termbox.ColorDefault
var PasswordColorBg termbox.Attribute = termbox.ColorDefault

var AnswerRune rune = '*'
var AnswerColorFg termbox.Attribute = termbox.ColorCyan
var AnswerColorBg termbox.Attribute = termbox.ColorDefault

/*
 * Functions
 */

func removeRule(runes []rune, rmPosition int) []rune {
	var newRunes []rune
	for index, rne := range runes {
		if index == rmPosition {
			continue
		}
		newRunes = append(newRunes, rne)
	}
	return newRunes
}

func Inquire(qStr string, x int, y int) (string, int, int) {
	// Print question
	q := question.New(qStr)
	x, y = q.PrintQuestion(x, y)
	// Print padding
	x, y = common.PrintString(
		PaddingQuestionAndAnswer,
		PaddingQuestionAndAnswerColorFg,
		PaddingQuestionAndAnswerColorBg,
		x, y)
	// Save start of answer
	answerStartX := x
	// Print cursor
	termbox.SetCursor(x, y)
	termbox.Flush()

	var answerRunes []rune
	// While typed EnterKey by keyboard
mainloop:
	for {
		// Get an event password
		event := common.GetEventKey()
		switch event.Key {
		// If get enter key, break password loop
		case termbox.KeyEnter:
			break mainloop
		case termbox.KeyBackspace:
			fallthrough
		case termbox.KeyBackspace2:
			if x > answerStartX {
				x--
				answerRunes = removeRule(answerRunes, x-answerStartX)
				termbox.SetCell(x, y, ' ', PasswordColorFg, PasswordColorBg)
				// Move cursor
				termbox.SetCursor(x, y)
				termbox.Flush()
			}
			continue
		}
		answerRunes = append(answerRunes, event.Ch)
		termbox.SetCell(x, y, PasswordRune, PasswordColorFg, PasswordColorBg)
		termbox.Flush()
		x++
		// Move cursor
		termbox.SetCursor(x, y)
		termbox.Flush()
	}
	// Print answer
	x = answerStartX
	for count := 0; count < len(answerRunes); count++ {
		termbox.SetCell(x, y, AnswerRune, AnswerColorFg, AnswerColorBg)
		termbox.Flush()
		x++
	}
	// Go to next line
	x, y = common.GoNextLine(x, y)

	return string(answerRunes), x, y
}
