package checkbox

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

type Checkbox struct {
	ChoiceStrs     []string
	CursorPosition int
	ChosePositions []int
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

var ChoseStr string = "⬢ "
var ChoseColorFg termbox.Attribute = termbox.ColorCyan
var ChoseColorBg termbox.Attribute = termbox.ColorDefault

var NoChoseStr string = "⬡ "
var NoChoseColorFg termbox.Attribute = termbox.ColorCyan
var NoChoseColorBg termbox.Attribute = termbox.ColorCyan

var ChoiceColorFg termbox.Attribute = termbox.ColorDefault
var ChoiceColorBg termbox.Attribute = termbox.ColorDefault

var NextLineStr string = "\n"
var NextLineColorFg termbox.Attribute = termbox.ColorDefault
var NextLineColorBg termbox.Attribute = termbox.ColorDefault

/*
 * Functions
 */

func contains(position int, chosePositions []int) bool {
	for _, chosePosition := range chosePositions {
		if position == chosePosition {
			return true
		}
	}
	return false
}

func New(choiceStrs []string) *Checkbox {
	checkbox := new(Checkbox)
	checkbox.ChoiceStrs = choiceStrs
	checkbox.CursorPosition = 0
	return checkbox
}

func (checkbox Checkbox) Print(x int, y int) (int, int) {
	// For each choice
	for index, choiceStr := range checkbox.ChoiceStrs {
		if index == checkbox.CursorPosition {
			// If the choice is selected, print cursor
			x, y = common.PrintString(
				CursorStr, CursorColorFg, CursorColorBg, x, y)
		} else {
			// If the choice is not selected, print no cursor
			x, y = common.PrintString(
				NoCursorStr, NoCursorColorFg, NoCursorColorBg, x, y)
		}
		// Print chose or no chose
		if contains(index, checkbox.ChosePositions) {
			x, y = common.PrintString(
				ChoseStr, ChoseColorFg, ChoseColorBg, x, y)
		} else {
			x, y = common.PrintString(
				NoChoseStr, NoChoseColorFg, NoChoseColorBg, x, y)
		}
		// Print choice
		x, y = common.PrintString(
			choiceStr, ChoiceColorFg, ChoiceColorBg, x, y)

		// If the choice is not last one, print next line
		if index != len(checkbox.ChoiceStrs)-1 {
			x, y = common.GoNextLine(x, y)
		}
	}
	return x, y
}

func (checkbox *Checkbox) Inquire(qStr string, x int, y int) ([]string, int, int) {
	// Save start point
	startX := x
	startY := y
	// Print question
	q := question.New(qStr)
	x, y = q.PrintQuestion(x, y)
	x, y = common.GoNextLine(x, y)

	// Print checkbox
	x, y = checkbox.Print(x, y)
	x, y = common.GoNextLine(x, y)

	// While typed Enter by keyboard
mainloop:
	for {
		// Get a key input
		key := common.GetKey()
		switch key {
		case termbox.KeyArrowUp:
			if checkbox.CursorPosition > 0 {
				checkbox.CursorPosition--
			}
		case termbox.KeyArrowDown:
			if checkbox.CursorPosition < len(checkbox.ChoiceStrs)-1 {
				checkbox.CursorPosition++
			}
		case termbox.KeySpace:
			checkbox.ChosePositions = append(checkbox.ChosePositions, checkbox.CursorPosition)
		case termbox.KeyEnter:
			break mainloop
		}
		// Go back to start point
		x, y = startX, startY
		// Reprint question
		x, y = q.PrintQuestion(x, y)
		x, y = common.GoNextLine(x, y)
		// Reprint checkbox
		x, y = checkbox.Print(x, y)
		x, y = common.GoNextLine(x, y)
	}
	x, y = common.GoNextLine(x, y)
	choseChoiceStrs := checkbox.GetChoseChoiceStrs()
	return choseChoiceStrs, x, y
}

func (checkbox Checkbox) GetChoseChoiceStrs() (choseChoiceStrs []string) {
	for _, chosePosition := range checkbox.ChosePositions {
		choseChoiceStrs = append(choseChoiceStrs, checkbox.ChoiceStrs[chosePosition])
	}
	return choseChoiceStrs
}

func (checkbox *Checkbox) AddChoice(choiceStr string) {
	checkbox.ChoiceStrs = append(checkbox.ChoiceStrs, choiceStr)
}

func (checkbox *Checkbox) RemoveChoice(rmChoiceStr string) {
	var newChoiceStrs []string
	for _, choiceStr := range checkbox.ChoiceStrs {
		if choiceStr == rmChoiceStr {
			continue
		}
		newChoiceStrs = append(newChoiceStrs, choiceStr)
	}
	checkbox.ChoiceStrs = newChoiceStrs
}
