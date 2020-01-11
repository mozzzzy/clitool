package checkbox

/*
 * Module Dependencies
 */

import (
	"strings"

	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/common"
	"github.com/mozzzzy/clitool/question"
)

/*
 * Types
 */

type Checkbox struct {
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
	choseColorBg    color.Color
	choseColorFg    color.Color
	choseStr        string
	noChoseColorBg  color.Color
	noChoseColorFg  color.Color
	noChoseStr      string
	minX            int
	maxX            int
	minY            int
	maxY            int
	chosePositions  []int
	cursorPosition  int
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

	DEFAULT_CHOSE_COLOR_BG color.Color = color.Default
	DEFAULT_CHOSE_COLOR_FG color.Color = color.Green
	DEFAULT_CHOSE_STR      string      = "⬢ "

	DEFAULT_NO_CHOSE_COLOR_BG color.Color = color.Default
	DEFAULT_NO_CHOSE_COLOR_FG color.Color = color.Default
	DEFAULT_NO_CHOSE_STR      string      = "⬡ "
)

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

func New(questionStr string, choices []string) *Checkbox {
	checkbox := new(Checkbox)
	checkbox.SetQuestion(questionStr)

	checkbox.choiceColorBg = DEFAULT_CHOICE_COLOR_BG
	checkbox.choiceColorFg = DEFAULT_CHOICE_COLOR_FG
	checkbox.choiceStrs = choices

	checkbox.cursorColorBg = DEFAULT_CURSOR_COLOR_BG
	checkbox.cursorColorFg = DEFAULT_CURSOR_COLOR_FG
	checkbox.cursorStr = DEFAULT_CURSOR_STR

	checkbox.noCursorColorBg = DEFAULT_NO_CURSOR_COLOR_BG
	checkbox.noCursorColorFg = DEFAULT_NO_CURSOR_COLOR_FG
	checkbox.noCursorStr = DEFAULT_NO_CURSOR_STR

	checkbox.choseColorBg = DEFAULT_CHOSE_COLOR_BG
	checkbox.choseColorFg = DEFAULT_CHOSE_COLOR_FG
	checkbox.choseStr = DEFAULT_CHOSE_STR

	checkbox.noChoseColorBg = DEFAULT_NO_CHOSE_COLOR_BG
	checkbox.noChoseColorFg = DEFAULT_NO_CHOSE_COLOR_FG
	checkbox.noChoseStr = DEFAULT_NO_CHOSE_STR

	checkbox.cursorPosition = 0
	return checkbox
}

func (checkbox *Checkbox) GetMinX() int {
	return checkbox.minX
}

func (checkbox *Checkbox) GetMaxX() int {
	return checkbox.maxX
}

func (checkbox *Checkbox) GetMinY() int {
	return checkbox.minY
}

func (checkbox *Checkbox) GetMaxY() int {
	return checkbox.maxY
}

func (checkbox *Checkbox) SetMinX(minX int) {
	// Set minX to inner question for q.Print() method.
	checkbox.q.SetMinX(minX)
	// Set minX
	checkbox.minX = minX
}

func (checkbox *Checkbox) SetMinY(minY int) {
	// Set minY to inner question for q.Print() method.
	checkbox.q.SetMinY(minY)
	// Set minY
	checkbox.minY = minY
}

func (checkbox *Checkbox) SetQuestion(qStr string) {
	q := question.New(qStr)
	checkbox.q = *q
}

func (checkbox *Checkbox) Print() {
	// NOTE if this checkbox has alread printed, its maxX and maxY also has already set.
	//      But the case that its length is changed exists.
	//      So I clear old maxX and maxY here.
	//      These values are calculated again in following logic.
	checkbox.maxX = 0
	checkbox.maxY = 0

	// Print question
	checkbox.q.Print()
	checkbox.updateMinMax(checkbox.q.GetMaxX(), checkbox.q.GetMaxY())

	if !checkbox.resolved {
		// Get next line of question
		x, y := common.GoNextLine(0, checkbox.q.GetMaxY())
		checkbox.updateMinMax(x, y)

		// For each choice
		for index, choiceStr := range checkbox.choiceStrs {
			if index == checkbox.cursorPosition {
				// If the choice is selected, print cursor
				x, y = common.PrintString(
					checkbox.cursorStr, checkbox.cursorColorFg, checkbox.cursorColorBg, x, y)
				checkbox.updateMinMax(x, y)
			} else {
				// If the choice is not selected, print no cursor
				x, y = common.PrintString(
					checkbox.noCursorStr, checkbox.noCursorColorFg, checkbox.noCursorColorBg, x, y)
				checkbox.updateMinMax(x, y)
			}
			// Print chose or no chose
			if contains(index, checkbox.chosePositions) {
				x, y = common.PrintString(
					checkbox.choseStr, checkbox.choseColorFg, checkbox.choseColorBg, x, y)
				checkbox.updateMinMax(x, y)
			} else {
				x, y = common.PrintString(
					checkbox.noChoseStr, checkbox.noChoseColorFg, checkbox.noChoseColorBg, x, y)
				checkbox.updateMinMax(x, y)
			}
			// Print choice
			x, y = common.PrintString(
				choiceStr, checkbox.choiceColorFg, checkbox.choiceColorBg, x, y)
			checkbox.updateMinMax(x, y)

			// If the choice is not last one, print next line
			if index != len(checkbox.choiceStrs)-1 {
				x, y = common.GoNextLine(x, y)
				checkbox.updateMinMax(x, y)
			}
		}
	}
}

func (checkbox *Checkbox) Inquire() interface{} {
	checkbox.Print()

	// While typed Enter by keyboard
mainloop:
	for {
		// Get a key input
		event := common.GetEventKey()
		switch event.Key {
		case termbox.KeyArrowUp:
			if checkbox.cursorPosition > 0 {
				checkbox.cursorPosition--
			}
		case termbox.KeyArrowDown:
			if checkbox.cursorPosition < len(checkbox.choiceStrs)-1 {
				checkbox.cursorPosition++
			}
		case termbox.KeySpace:
			if contains(checkbox.cursorPosition, checkbox.chosePositions) {
				checkbox.RemoveChose(checkbox.cursorPosition)
			} else {
				checkbox.chosePositions = append(checkbox.chosePositions, checkbox.cursorPosition)
			}
		case termbox.KeyEnter:
			if len(checkbox.chosePositions) > 0 {
				checkbox.q.Resolve(strings.Join(checkbox.GetChoseChoiceStrs(), ","))
				checkbox.resolved = true
				break mainloop
			}
		}
		switch event.Ch {
		case 'k':
			if checkbox.cursorPosition > 0 {
				checkbox.cursorPosition--
			}
		case 'j':
			if checkbox.cursorPosition < len(checkbox.choiceStrs)-1 {
				checkbox.cursorPosition++
			}
		}
		checkbox.Print()
	}
	checkbox.Print()
	choseChoiceStrs := checkbox.GetChoseChoiceStrs()
	return choseChoiceStrs
}

func (checkbox Checkbox) GetChoseChoiceStrs() (choseChoiceStrs []string) {
	for _, chosePosition := range checkbox.chosePositions {
		choseChoiceStrs = append(choseChoiceStrs, checkbox.choiceStrs[chosePosition])
	}
	return choseChoiceStrs
}

func (checkbox *Checkbox) AddChoice(choiceStr string) {
	checkbox.choiceStrs = append(checkbox.choiceStrs, choiceStr)
}

func (checkbox *Checkbox) RemoveChoice(rmChoiceStr string) {
	var newChoiceStrs []string
	for _, choiceStr := range checkbox.choiceStrs {
		if choiceStr == rmChoiceStr {
			continue
		}
		newChoiceStrs = append(newChoiceStrs, choiceStr)
	}
	checkbox.choiceStrs = newChoiceStrs
}

func (checkbox *Checkbox) RemoveChose(rmChosePosition int) {
	var newChosePositions []int
	for _, chosePosition := range checkbox.chosePositions {
		if chosePosition == rmChosePosition {
			continue
		}
		newChosePositions = append(newChosePositions, chosePosition)
	}
	checkbox.chosePositions = newChosePositions
}

func (checkbox *Checkbox) updateMinMax(x int, y int) {
	// Update maxX
	if checkbox.maxX < x {
		checkbox.maxX = x
	}
	// Update maxY
	if checkbox.maxY < y {
		checkbox.maxY = y
	}
	// Update minX
	if checkbox.minX > x {
		checkbox.minX = x
	}
	// Update minY
	if checkbox.minY > y {
		checkbox.minY = y
	}
}
