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

func (terminal *Terminal) Spinner(messageStr string, finished *bool) {
	spnr := spinner.New(messageStr)
	go spnr.Spin(terminal.x, terminal.y, finished)
	terminal.x, terminal.y = common.GoNextLine(terminal.x, terminal.y)
}

func (terminal *Terminal) List(question string, choiceStrs []string) (answerStr string) {
	lst := list.New(choiceStrs)
	answerStr, terminal.x, terminal.y = lst.Inquire(question, terminal.x, terminal.y)
	return answerStr
}

/*
func main() {
	Init()

	terminal := Terminal{
		x: 0,
		y: 0,
	}

	lst := list.New([]string{"go", "javascript", "c++"})
	_, terminal.x, terminal.y = lst.Inquire("What is your favorite language?", terminal.x, terminal.y)

	fin := new(bool)
	*fin = false
	spnr := spinner.New("Please wait for 5 seconds")
	go spnr.Spin(terminal.x, terminal.y, fin)

	time.Sleep(5 * time.Second)
	*fin = true

	time.Sleep(2 * time.Second)
	Close()
}
*/
