package main

//package clitool

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/nsf/termbox-go"

	"./list"
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
const PREFIX_QUESTION_RUNE rune = '?'

/*
 * Functions
 */

func getKey() termbox.Key {
	var returnKey termbox.Key

	// Poll event
	switch ev := termbox.PollEvent(); ev.Type {
	// Keyboard is typed
	case termbox.EventKey:
		returnKey = ev.Key
	// Terminal is resized
	case termbox.EventResize:
		// TODO implement if needed
	}
	return returnKey
}

func Close() {
	termbox.Close()
}

func Init() error {
	err := termbox.Init()
	return err
}

func (terminal *Terminal) goNextLine() {
	terminal.x = 0
	terminal.y++
}

func (terminal *Terminal) printQuestion(question string) {
	// Print prefix
	colorForeGround := termbox.ColorGreen
	colorBackGround := termbox.ColorDefault
	termbox.SetCell(
		terminal.x, terminal.y, PREFIX_QUESTION_RUNE, colorForeGround, colorBackGround,
	)
	terminal.x++

	// Print space between prefix and question
	termbox.SetCell(
		terminal.x, terminal.y, ' ', colorForeGround, colorBackGround,
	)
	terminal.x++

	// Print question
	questionRunes := []rune(question)
	colorForeGround = termbox.ColorDefault
	colorBackGround = termbox.ColorDefault
	for _, char := range questionRunes {
		termbox.SetCell(terminal.x, terminal.y, rune(char), colorForeGround, colorBackGround)
		terminal.x++
	}
}

func (terminal *Terminal) printAnswer(question string, answer string) {
	// Print question
	terminal.printQuestion(question)

	// Print space between question and answer
	colorForeGround := termbox.ColorCyan
	colorBackGround := termbox.ColorDefault
	termbox.SetCell(terminal.x, terminal.y, ' ', colorForeGround, colorBackGround)
	terminal.x++

	// Print answer
	answerRunes := []rune(answer)
	for _, char := range answerRunes {
		termbox.SetCell(terminal.x, terminal.y, rune(char), colorForeGround, colorBackGround)
		terminal.x++
	}
}

func (terminal *Terminal) printList(list *list.List) {
	// Print list
	listRunes := []rune(list.String())
	colorForeGround := termbox.ColorDefault
	colorBackGround := termbox.ColorDefault
	for _, char := range listRunes {
		switch char {
		case '❯':
			colorForeGroundOrg := colorForeGround
			colorForeGround = termbox.ColorGreen
			termbox.SetCell(terminal.x, terminal.y, rune(char), colorForeGround, colorBackGround)
			colorForeGround = colorForeGroundOrg
			terminal.x++
		case '\n':
			terminal.goNextLine()
		default:
			termbox.SetCell(terminal.x, terminal.y, rune(char), colorForeGround, colorBackGround)
			terminal.x++
		}
	}
}

func (terminal *Terminal) List(question string, choices []string) string {
	// Save start point
	zeroX := terminal.x
	zeroY := terminal.y

	// Print question
	terminal.printQuestion(question)
	// Move x,y to the start of next line
	terminal.goNextLine()

	// Create list
	list := list.New(choices)
	// Print list
	terminal.printList(list)

	// Flush
	termbox.Flush()

mainloop:
	// While typed Enter by keyboard
	for {
		// Get a key input
		key := getKey()
		switch key {
		case termbox.KeyArrowUp:
			if list.CursorPosition > 0 {
				list.CursorPosition--
			}
		case termbox.KeyArrowDown:
			if list.CursorPosition < len(list.Choices)-1 {
				list.CursorPosition++
			}
		case termbox.KeyEnter:
			break mainloop
		}

		// Go back to start point
		terminal.x, terminal.y = zeroX, zeroY

		// Reprint question
		terminal.printQuestion(question)
		// Move x,y to the start of next line
		terminal.goNextLine()

		// Reprint list
		terminal.printList(list)

		// Flush
		termbox.Flush()
	}

	// Save end point
	endX, endY := terminal.x, terminal.y
	// Go back to start point
	terminal.x, terminal.y = zeroX, zeroY
	// Print answer
	terminal.printAnswer(question, list.Choices[list.CursorPosition])
	// Flush
	termbox.Flush()

	// Move x,y to the end of list
	terminal.x, terminal.y = endX, endY
	// Move x,y to the start of next line
	terminal.goNextLine()

	return list.Choices[list.CursorPosition]
}

func main() {
	terminal := Terminal{
		x: 0,
		y: 0,
	}
	Init()
	terminal.List("What language do you like?", []string{"go", "javascript", "c++"})
	terminal.List("What language do you like?", []string{"go", "javascript", "c++"})
	terminal.List("What language do you like?", []string{"go", "javascript", "c++"})
	time.Sleep(2 * time.Second)
	Close()
}
