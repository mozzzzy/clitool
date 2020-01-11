package clitool

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/clitool/common"
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
