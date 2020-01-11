package confirm

/*
 * Module Dependencies
 */

import (
	"fmt"

	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/common"
	"github.com/mozzzzy/clitool/question"
)

/*
 * Types
 */

type Confirm struct {
	q                   question.Question
	parenthesisLeftStr  string
	parenthesisRightStr string
	separatorStr        string
	candidateColorBg    color.Color
	candidateColorFg    color.Color
	answerColorBg       color.Color
	answerColorFg       color.Color
	yesRune             rune
	noRune              rune
	minX                int
	maxX                int
	minY                int
	maxY                int
	resolved            bool
}

/*
 * Constants and Package Scope Variables
 */

var (
	DEFAULT_PARENTHESIS_LEFT_STR  string      = "("
	DEFAULT_PARENTHESIS_RIGHT_STR string      = ")"
	DEFAULT_SEPARATOR_STR         string      = "/"
	DEFAULT_YES_RUNE              rune        = 'y'
	DEFAULT_NO_RUNE               rune        = 'N'
	DEFAULT_CANDIDATE_COLOR_BG    color.Color = color.Default
	DEFAULT_CANDIDATE_COLOR_FG    color.Color = color.Default
	DEFAULT_ANSWER_COLOR_BG       color.Color = color.Default
	DEFAULT_ANSWER_COLOR_FG       color.Color = color.Cyan
)

/*
 * Functions
 */

func New(questionStr string) *Confirm {
	confirm := new(Confirm)
	confirm.SetQuestion(questionStr)
	confirm.parenthesisLeftStr = DEFAULT_PARENTHESIS_LEFT_STR
	confirm.parenthesisRightStr = DEFAULT_PARENTHESIS_RIGHT_STR
	confirm.separatorStr = DEFAULT_SEPARATOR_STR
	confirm.yesRune = DEFAULT_YES_RUNE
	confirm.noRune = DEFAULT_NO_RUNE
	confirm.candidateColorBg = DEFAULT_CANDIDATE_COLOR_BG
	confirm.candidateColorFg = DEFAULT_CANDIDATE_COLOR_FG
	confirm.answerColorBg = DEFAULT_ANSWER_COLOR_BG
	confirm.answerColorFg = DEFAULT_ANSWER_COLOR_FG
	return confirm
}

func (confirm *Confirm) GetMinX() int {
	return confirm.minX
}

func (confirm *Confirm) GetMaxX() int {
	return confirm.maxX
}

func (confirm *Confirm) GetMinY() int {
	return confirm.minY
}

func (confirm *Confirm) GetMaxY() int {
	return confirm.maxY
}

func (confirm *Confirm) SetMinX(minX int) {
	// Set minX to inner question for q.Print() method.
	confirm.q.SetMinX(minX)
	// Set minX
	confirm.minX = minX
}

func (confirm *Confirm) SetMinY(minY int) {
	// Set minY to inner question for q.Print() method.
	confirm.q.SetMinY(minY)
	// Set minY
	confirm.minY = minY
}

func (confirm *Confirm) SetQuestion(qStr string) {
	q := question.New(qStr)
	confirm.q = *q
}

func (confirm *Confirm) Print() {
	// NOTE if this confirm has alread printed, its maxX and maxY also has already set.
	//      But the case that its length is changed exists.
	//      So I clear old maxX and maxY here.
	//      These values are calculated again in following logic.
	confirm.maxX = 0
	confirm.maxY = 0

	// Set parenthesis, separator, candidates
	if !confirm.resolved {
		candidates := fmt.Sprintf("%v%v%v%v%v",
			confirm.parenthesisLeftStr,
			string(confirm.yesRune),
			confirm.separatorStr,
			string(confirm.noRune),
			confirm.parenthesisRightStr,
		)
		confirm.q.SetSuffix(confirm.candidateColorFg, confirm.candidateColorBg, candidates)
	} else {
		confirm.q.SetSuffixColor(confirm.answerColorFg, confirm.answerColorBg)
	}
	// Print question
	confirm.q.Print()
	confirm.updateMinMax(confirm.q.GetMaxX(), confirm.q.GetMaxY())
}

func (confirm *Confirm) Inquire() interface{} {
	confirm.Print()

	var answerBool bool
	// While typed YesRune or NoRune by keyboard
mainloop:
	for {
		// Get an event input
		event := common.GetEventKey()
		switch event.Ch {
		case confirm.yesRune:
			answerBool = true
			confirm.resolved = true
			confirm.q.Resolve(string(confirm.yesRune))
			break mainloop
		case confirm.noRune:
			answerBool = false
			confirm.resolved = true
			confirm.q.Resolve(string(confirm.noRune))
			break mainloop
		}
	}

	for yCursor := confirm.minY; yCursor <= confirm.maxY; yCursor++ {
		for xCursor := confirm.minX; xCursor <= confirm.maxX; xCursor++ {
			common.PrintString(" ", color.Default, color.Default, xCursor, yCursor)
		}
	}
	confirm.Print()

	return answerBool
}

func (confirm *Confirm) updateMinMax(x int, y int) {
	// Update maxX
	if confirm.maxX < x {
		confirm.maxX = x
	}
	// Update maxY
	if confirm.maxY < y {
		confirm.maxY = y
	}
	// Update minX
	if confirm.minX > x {
		confirm.minX = x
	}
	// Update minY
	if confirm.minY > y {
		confirm.minY = y
	}
}
