package clitool

/*
 * Module Dependencies
 */

import (
	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/checkbox"
	"github.com/mozzzzy/clitool/common"
	"github.com/mozzzzy/clitool/confirm"
	"github.com/mozzzzy/clitool/input"
	"github.com/mozzzzy/clitool/list"
	"github.com/mozzzzy/clitool/message"
	"github.com/mozzzzy/clitool/password"
	"github.com/mozzzzy/clitool/progressbar"
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

func (terminal *Terminal) Checkbox(question string, choiceStrs []string) (choseStrs []string) {
	chkbox := checkbox.New(choiceStrs)
	choseStrs, terminal.x, terminal.y = chkbox.Inquire(question, terminal.x, terminal.y)
	return choseStrs
}

func (terminal *Terminal) Confirm(question string) (answerBool bool) {
	answerBool, terminal.x, terminal.y = confirm.Inquire(question, terminal.x, terminal.y)
	return answerBool
}

func (terminal *Terminal) Input(question string) (answerStr string) {
	answerStr, terminal.x, terminal.y = input.Inquire(question, terminal.x, terminal.y)
	return answerStr
}

func (terminal *Terminal) List(question string, choiceStrs []string) (answerStr string) {
	lst := list.New(choiceStrs)
	answerStr, terminal.x, terminal.y = lst.Inquire(question, terminal.x, terminal.y)
	return answerStr
}

func (terminal *Terminal) Message(messageStr string) {
	msg := message.New(messageStr)
	terminal.x, terminal.y = msg.Print(terminal.x, terminal.y)
}

func (terminal *Terminal) Password(question string) (answerStr string) {
	answerStr, terminal.x, terminal.y = password.Inquire(question, terminal.x, terminal.y)
	return answerStr
}

func (terminal *Terminal) Progressbar(
	messageStr string, min float64, state *float64, max float64,
) {
	pbar := progressbar.New(messageStr, min, state, max)
	go pbar.Show(terminal.x, terminal.y)
	terminal.x, terminal.y = common.GoNextLine(terminal.x, terminal.y)
}

func (terminal *Terminal) Spinner(messageStr string, finished *bool) {
	spnr := spinner.New(messageStr)
	go spnr.Spin(terminal.x, terminal.y, finished)
	terminal.x, terminal.y = common.GoNextLine(terminal.x, terminal.y)
}
