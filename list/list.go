package list

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

type List struct {
	q               question.Question
	choiceColorBg   color.Color
	choiceColorFg   color.Color
	choiceStrs      []string
	cursorColorBg   color.Color
	cursorColorFg   color.Color
	cursorStr       string
	noCursorColorBg color.Color
	noCursorColorFg color.Color
	noCursorStr     string
	cursorPosition  int
	minX            int
	maxX            int
	minY            int
	maxY            int
	resolved        bool
}

/*
 * Constants and Package Scope Variables
 */

var (
	DEFAULT_CHOICE_COLOR_BG color.Color = color.Default
	DEFAULT_CHOICE_COLOR_FG color.Color = color.Default

	DEFAULT_CURSOR_COLOR_BG color.Color = color.Default
	DEFAULT_CURSOR_COLOR_FG color.Color = color.Cyan
	DEFAULT_CURSOR_STR      string      = "❯ "

	DEFAULT_NO_CURSOR_COLOR_BG color.Color = color.Default
	DEFAULT_NO_CURSOR_COLOR_FG color.Color = color.Cyan
	DEFAULT_NO_CURSOR_STR      string      = "  "
)

/*
 * Functions
 */

func New(questionStr string, choices []string) *List {
	list := new(List)
	list.SetQuestion(questionStr)

	list.choiceColorBg = DEFAULT_CHOICE_COLOR_BG
	list.choiceColorFg = DEFAULT_CHOICE_COLOR_FG
	list.choiceStrs = choices

	list.cursorColorBg = DEFAULT_CURSOR_COLOR_BG
	list.cursorColorFg = DEFAULT_CURSOR_COLOR_FG
	list.cursorStr = DEFAULT_CURSOR_STR

	list.noCursorColorBg = DEFAULT_NO_CURSOR_COLOR_BG
	list.noCursorColorFg = DEFAULT_NO_CURSOR_COLOR_FG
	list.noCursorStr = DEFAULT_NO_CURSOR_STR

	list.cursorPosition = 0
	return list
}

func (list *List) GetMinX() int {
	return list.minX
}

func (list *List) GetMaxX() int {
	return list.maxX
}

func (list *List) GetMinY() int {
	return list.minY
}

func (list *List) GetMaxY() int {
	return list.maxY
}

func (list *List) SetMinX(minX int) {
	// Set minX to inner question for q.Print() method.
	list.q.SetMinX(minX)
	// Set minX
	list.minX = minX
}

func (list *List) SetMinY(minY int) {
	// Set minY to inner question for q.Print() method.
	list.q.SetMinY(minY)
	// Set minY
	list.minY = minY
}

func (list *List) SetQuestion(qStr string) {
	q := question.New(qStr)
	list.q = *q
}

func (list *List) Print() {
	// NOTE if this list has alread printed, its maxX and maxY also has already set.
	//      But the case that its length is changed exists.
	//      So I clear old maxX and maxY here.
	//      These values are calculated again in following logic.
	list.maxX = 0
	list.maxY = 0

	// Print question
	list.q.Print()
	list.updateMinMax(list.q.GetMaxX(), list.q.GetMaxY())

	if !list.resolved {
		// Get next line of question
		x, y := common.GoNextLine(0, list.q.GetMaxY())
		list.updateMinMax(x, y)

		// For each choice
		for index, choiceStr := range list.choiceStrs {
			if index == list.cursorPosition {
				// If the choice is selected, print cursor
				x, y = common.PrintString(
					list.cursorStr, list.cursorColorFg, list.cursorColorBg, x, y)
				list.updateMinMax(x, y)
			} else {
				// If the choice is not selected, print no cursor
				x, y = common.PrintString(
					list.noCursorStr, list.noCursorColorFg, list.noCursorColorBg, x, y)
				list.updateMinMax(x, y)
			}
			// Print choice
			x, y = common.PrintString(
				choiceStr, list.choiceColorFg, list.choiceColorBg, x, y)
			list.updateMinMax(x, y)

			// If the choice is not last one, print next line
			if index != len(list.choiceStrs)-1 {
				x, y = common.GoNextLine(x, y)
				list.updateMinMax(x, y)
			}
		}
	}
}

func (list *List) Inquire() interface{} {
	list.Print()

	// While typed Enter by keyboard
mainloop:
	for {
		// Get a key input
		event := common.GetEventKey()
		switch event.Key {
		case termbox.KeyArrowUp:
			if list.cursorPosition > 0 {
				list.cursorPosition--
			}
		case termbox.KeyArrowDown:
			if list.cursorPosition < len(list.choiceStrs)-1 {
				list.cursorPosition++
			}
		case termbox.KeyEnter:
			list.q.Resolve(list.choiceStrs[list.cursorPosition])
			list.resolved = true
			break mainloop
		}
		switch event.Ch {
		case 'k':
			if list.cursorPosition > 0 {
				list.cursorPosition--
			}
		case 'j':
			if list.cursorPosition < len(list.choiceStrs)-1 {
				list.cursorPosition++
			}
		}
		list.Print()
	}
	list.Print()
	return list.choiceStrs[list.cursorPosition]
}

func (list *List) AddChoice(choiceStr string) {
	list.choiceStrs = append(list.choiceStrs, choiceStr)
}

func (list *List) RemoveChoice(rmChoiceStr string) {
	var newChoiceStrs []string
	for _, choiceStr := range list.choiceStrs {
		if choiceStr == rmChoiceStr {
			continue
		}
		newChoiceStrs = append(newChoiceStrs, choiceStr)
	}
	list.choiceStrs = newChoiceStrs
}

func (list *List) updateMinMax(x int, y int) {
	// Update maxX
	if list.maxX < x {
		list.maxX = x
	}
	// Update maxY
	if list.maxY < y {
		list.maxY = y
	}
	// Update minX
	if list.minX > x {
		list.minX = x
	}
	// Update minY
	if list.minY > y {
		list.minY = y
	}
}
