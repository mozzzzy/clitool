package list

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

type List struct {
	ChoiceStrs     []string
	CursorPosition int
}

/*
 * Constants and Package Scope Variables
 */

var CursorStr string = "❯ "
var CursorColorFg termbox.Attribute = termbox.ColorCyan
var CursorColorBg termbox.Attribute = termbox.ColorDefault

var NoCursorStr string = "  "
var NoCursorColorFg termbox.Attribute = termbox.ColorDefault
var NoCursorColorBg termbox.Attribute = termbox.ColorDefault

var ChoiceColorFg termbox.Attribute = termbox.ColorDefault
var ChoiceColorBg termbox.Attribute = termbox.ColorDefault

var NextLineStr string = "\n"
var NextLineColorFg termbox.Attribute = termbox.ColorDefault
var NextLineColorBg termbox.Attribute = termbox.ColorDefault

/*
 * Functions
 */

func New(choiceStrs []string) *List {
	list := new(List)
	list.ChoiceStrs = choiceStrs
	list.CursorPosition = 0
	return list
}

func (list List) Print(x int, y int) (int, int) {
	// For each choice
	for index, choiceStr := range list.ChoiceStrs {
		if index == list.CursorPosition {
			// If the choice is selected, print cursor
			x, y = common.PrintString(
				CursorStr, CursorColorFg, CursorColorBg, x, y)
		} else {
			// If the choice is not selected, print no cursor
			x, y = common.PrintString(
				NoCursorStr, NoCursorColorFg, NoCursorColorBg, x, y)
		}
		// Print choice
		x, y = common.PrintString(
			choiceStr, ChoiceColorFg, ChoiceColorBg, x, y)

		// If the choice is not last one, print next line
		if index != len(list.ChoiceStrs)-1 {
			x, y = common.GoNextLine(x, y)
		}
	}
	return x, y
}

func (list *List) Inquire(qStr string, x int, y int) (string, int, int) {
	// Save start point
	startX := x
	startY := y
	// Print question
	q := question.New(qStr)
	x, y = q.PrintQuestion(x, y)
	x, y = common.GoNextLine(x, y)

	// Print list
	x, y = list.Print(x, y)
	x, y = common.GoNextLine(x, y)

	// While typed Enter by keyboard
mainloop:
	for {
		// Get a key input
		key := common.GetKey()
		switch key {
		case termbox.KeyArrowUp:
			if list.CursorPosition > 0 {
				list.CursorPosition--
			}
		case termbox.KeyArrowDown:
			if list.CursorPosition < len(list.ChoiceStrs)-1 {
				list.CursorPosition++
			}
		case termbox.KeyEnter:
			q.AnswerStr = list.ChoiceStrs[list.CursorPosition]
			break mainloop
		}
		// Go back to start point
		x, y = startX, startY
		// Reprint question
		x, y = q.PrintQuestion(x, y)
		x, y = common.GoNextLine(x, y)
		// Reprint list
		x, y = list.Print(x, y)
		x, y = common.GoNextLine(x, y)
	}
	// Save end point
	endX, endY := x, y
	// Go back to start point
	x, y = startX, startY
	// Print answer
	x, y = q.PrintAnswer(x, y)
	// Go to end point
	x, y = endX, endY
	x, y = common.GoNextLine(x, y)

	return list.ChoiceStrs[list.CursorPosition], x, y
}

func (list *List) AddChoice(choiceStr string) {
	list.ChoiceStrs = append(list.ChoiceStrs, choiceStr)
}

func (list *List) RemoveChoice(rmChoiceStr string) {
	var newChoiceStrs []string
	for _, choiceStr := range list.ChoiceStrs {
		if choiceStr == rmChoiceStr {
			continue
		}
		newChoiceStrs = append(newChoiceStrs, choiceStr)
	}
	list.ChoiceStrs = newChoiceStrs
}
