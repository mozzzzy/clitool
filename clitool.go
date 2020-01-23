package clitool

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/clitool/common"
	"github.com/mozzzzy/clitool/confirm"
	"github.com/mozzzzy/clitool/checkbox"
	"github.com/mozzzzy/clitool/errorMessage"
	"github.com/mozzzzy/clitool/input"
	"github.com/mozzzzy/clitool/message"
	"github.com/mozzzzy/clitool/list"
	"github.com/mozzzzy/clitool/password"
	"github.com/mozzzzy/clitool/progressbar"
	"github.com/mozzzzy/clitool/spinner"
	"github.com/mozzzzy/clitool/question"
	"github.com/nsf/termbox-go"
)

/*
 * Types
 */

type printable interface {
	Print()
	GetMinX() int
	GetMaxX() int
	GetMinY() int
	GetMaxY() int
	SetMinX(int)
	SetMinY(int)
}

type inquirable interface {
	Inquire() interface{}
	Print()
	GetMinX() int
	GetMaxX() int
	GetMinY() int
	GetMaxY() int
	SetMinX(int)
	SetMinY(int)
}

type runnable interface {
	Run()
	Print()
	GetMinX() int
	GetMaxX() int
	GetMinY() int
	GetMaxY() int
	SetMinX(int)
	SetMinY(int)
}

/*
 * Constants and Package Scope Variables
 */

var printables []printable

/*
 * Functions
 */

func Close() {
	termbox.Close()
}

func Init() error {
	go common.ExitByCtlC()
	err := termbox.Init()
	return err
}

func Print(p printable) {
	x, y := 0, GetMaxY(printables)
	if len(printables) != 0 {
		x, y = common.GoNextLine(x, y)
	}

	p.SetMinX(x)
	p.SetMinY(y)
	p.Print()
	printables = append(printables, p)
}

func RePrintAll() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	minX := 0
	minY := 0
	for index, p := range printables {
		if index == 0 {
			p.SetMinX(minX)
			p.SetMinY(minY)
		} else {
			minX, minY = common.GoNextLine(minX, GetMaxY(printables[:index]))
			p.SetMinX(minX)
			p.SetMinY(minY)
		}
		p.Print()
	}
}

func Inquire(i inquirable) interface{} {
	Print(i)
	answer := i.Inquire()
	RePrintAll()
	return answer
}

func Run(r runnable) {
	Print(r)
	go r.Run()
}

func GetMaxY(pSlice []printable) (y int) {
	for _, p := range pSlice {
		maxY := p.GetMaxY()
		if y < maxY {
			y = maxY
		}
	}
	return y
}

func WaitEsc() {
	common.GetEsc()
}

// Following functions are shorthand of Print and Inquire call of each types.
func Error(msg string) {
	errorMessage := errorMessage.New(msg)
	Print(errorMessage)
}

func Checkbox(question string, choices []string, choses []string) []string {
	chkbox := checkbox.New(question, choices)
	chkbox.Check(choses)
	answers := Inquire(chkbox)
	return answers.([]string)
}

func Confirm(question string) bool {
	cfm := confirm.New(question)
	answer := Inquire(cfm)
	return answer.(bool)
}

func Input(question string) string {
	input := input.New(question)
	answer := Inquire(input)
	return answer.(string)
}

func Message(msgStr string) {
	msg := message.New(msgStr)
	Print(msg)
}

func List(question string, choices []string) string {
	lst := list.New(question, choices)
	answer := Inquire(lst)
	return answer.(string)
}

func Password(question string) string {
	passwd := password.New(question)
	answer := Inquire(passwd)
	return answer.(string)
}

func Progressbar(
	message string, min float64, state *float64, max float64,
) *progressbar.Progressbar {
	pbar := progressbar.New(message, min, state, max)
	Run(pbar)
	return pbar
}

func Spinner(message string) *spinner.Spinner {
	spnr := spinner.New(message)
	Run(spnr)
	return spnr
}

func Question(message string) *question.Question {
	q := question.New(message)
	Print(q)
	return q
}
