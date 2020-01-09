package message

/*
 * Module Dependencies
 */

import (
	"github.com/nsf/termbox-go"

	"github.com/mozzzzy/clitool/common"
)

/*
 * Types
 */

type Message struct {
	MessageStr string
}

/*
 * Constants and Package Scope Variables
 */

var MessageColorFg termbox.Attribute = termbox.ColorDefault
var MessageColorBg termbox.Attribute = termbox.ColorDefault

var PrefixStr string = ""
var PrefixColorFg termbox.Attribute = termbox.ColorDefault
var PrefixColorBg termbox.Attribute = termbox.ColorDefault

var SuffixStr string = ""
var SuffixColorFg termbox.Attribute = termbox.ColorDefault
var SuffixColorBg termbox.Attribute = termbox.ColorDefault

/*
 * Functions
 */

func New(messageStr string) *Message {
	message := new(Message)
	message.MessageStr = messageStr
	return message
}

func (message Message) Print(x int, y int) (int, int) {
	// Print prefix
	x, y = common.PrintString(PrefixStr, PrefixColorFg, PrefixColorBg, x, y)
	// Print message
	x, y = common.PrintString(
		message.MessageStr, MessageColorFg, MessageColorBg, x, y)
	// Print suffix
	x, y = common.PrintString(SuffixStr, SuffixColorFg, SuffixColorBg, x, y)
	return x, y
}
