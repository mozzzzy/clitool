package clitool

/*
 * Module Dependencies
 */

import (
	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/common"
	"github.com/mozzzzy/clitool/list"
	"github.com/mozzzzy/clitool/spinner"
)

/*
 * Types
 */

type Terminal struct {
	x, y int
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func Close() {
	termbox.Close()
}

func Init() error {
	err := termbox.Init()
	return err
}

func New() *Terminal {
	terminal := new(Terminal)
	terminal.x = 0
	terminal.y = 0
	return terminal
}

func (terminal *Terminal) List(question string, choiceStrs []string) (answerStr string) {
	lst := list.New(choiceStrs)
	answerStr, terminal.x, terminal.y = lst.Inquire(question, terminal.x, terminal.y)
	return answerStr
}

func (terminal *Terminal) Spinner(messageStr string, finished *bool) {
	spnr := spinner.New(messageStr)
	go spnr.Spin(terminal.x, terminal.y, finished)
	terminal.x, terminal.y = common.GoNextLine(terminal.x, terminal.y)
}
