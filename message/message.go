package message

/*
 * Module Dependencies
 */

import (
	"strings"

	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/common"
)

/*
 * Types
 */

type Message struct {
	Message        string
	MessageColorBg color.Color
	MessageColorFg color.Color
	Prefix         string
	PrefixColorBg  color.Color
	PrefixColorFg  color.Color
	Suffix         string
	SuffixColorBg  color.Color
	SuffixColorFg  color.Color
	maxX           int
	minX           int
	maxY           int
	minY           int
}

/*
 * Constants and Package Scope Variables
 */

var (
	DEFAULT_MESSAGE_COLOR_BG color.Color = color.Default
	DEFAULT_MESSAGE_COLOR_FG color.Color = color.Default

	DEFAULT_PREFIX_COLOR_BG color.Color = color.Default
	DEFAULT_PREFIX_COLOR_FG color.Color = color.Default
	DEFAULT_PREFIX_STR      string      = ""

	DEFAULT_SUFFIX_COLOR_BG color.Color = color.Default
	DEFAULT_SUFFIX_COLOR_FG color.Color = color.Default
	DEFAULT_SUFFIX_STR      string      = ""
)

/*
 * Functions
 */

func New(msg string) *Message {
	message := new(Message)

	message.MessageColorBg = DEFAULT_MESSAGE_COLOR_BG
	message.MessageColorFg = DEFAULT_MESSAGE_COLOR_FG
	message.Message = msg

	message.PrefixColorBg = DEFAULT_PREFIX_COLOR_BG
	message.PrefixColorFg = DEFAULT_PREFIX_COLOR_FG
	message.Prefix = DEFAULT_PREFIX_STR

	message.SuffixColorBg = DEFAULT_SUFFIX_COLOR_BG
	message.SuffixColorFg = DEFAULT_SUFFIX_COLOR_FG
	message.Suffix = DEFAULT_SUFFIX_STR

	return message
}

func (message *Message) GetMinX() int {
	return message.minX
}

func (message *Message) GetMaxX() int {
	return message.maxX
}

func (message *Message) GetMinY() int {
	return message.minY
}

func (message *Message) GetMaxY() int {
	return message.maxY
}

func (message *Message) Print() {
	// NOTE if this message has alread printed, its maxX and maxY also has already set.
	//      But the case that its length is changed exists.
	//      So I clear old maxX and maxY here.
	//      These values are calculated again in following logic.
	message.maxX = 0
	message.maxY = 0

	x, y := message.minX, message.minY

	// Split prefix, message, and suffix strings by "\n"
	prefixLines := strings.Split(message.Prefix, "\n")
	messageLines := strings.Split(message.Message, "\n")
	suffixLines := strings.Split(message.Suffix, "\n")

	// Print prefix
	x, y = message.printMultiLines(
		prefixLines, message.PrefixColorFg, message.PrefixColorBg, x, y)

	// Print message
	x, y = message.printMultiLines(
		messageLines, message.MessageColorFg, message.MessageColorBg, x, y)

	// Print suffix
	x, y = message.printMultiLines(
		suffixLines, message.SuffixColorFg, message.SuffixColorBg, x, y)
}

func (message *Message) SetMessage(
	messageColorFg color.Color, messageColorBg color.Color, msg string) {
	message.SetMessageColor(messageColorFg, messageColorBg)
	message.Message = msg
}

func (message *Message) SetMessageColor(
	messageColorFg color.Color, messageColorBg color.Color) {
	message.MessageColorFg = messageColorFg
	message.MessageColorBg = messageColorBg
}

func (message *Message) SetMinX(minX int) {
	message.minX = minX
}

func (message *Message) SetMinY(minY int) {
	message.minY = minY
}

func (message *Message) SetPrefix(
	prefixColorFg color.Color, prefixColorBg color.Color, prefix string) {
	message.PrefixColorFg = prefixColorFg
	message.PrefixColorBg = prefixColorBg
	message.Prefix = prefix
}

func (message *Message) SetPrefixStr(prefix string) {
	message.Prefix = prefix
}

func (message *Message) SetPrefixColor(
	prefixColorFg color.Color, prefixColorBg color.Color) {
	message.PrefixColorFg = prefixColorFg
	message.PrefixColorBg = prefixColorBg
}

func (message *Message) SetSuffix(
	suffixColorFg color.Color, suffixColorBg color.Color, suffix string) {
	message.SuffixColorFg = suffixColorFg
	message.SuffixColorBg = suffixColorBg
	message.Suffix = suffix
}

func (message *Message) SetSuffixStr(suffix string) {
	message.Suffix = suffix
}

func (message *Message) SetSuffixColor(
	suffixColorFg color.Color, suffixColorBg color.Color) {
	message.SuffixColorFg = suffixColorFg
	message.SuffixColorBg = suffixColorBg
}

func (message *Message) printMultiLines(
	lines []string, colorFg color.Color, colorBg color.Color, x int, y int,
) (int, int) {
	for index, line := range lines {
		// Print one line
		x, y = common.PrintString(line, colorFg, colorBg, x, y)
		message.updateMinMax(x, y)

		// If this line is not last line, go to next line
		if index < len(lines)-1 {
			x, y = common.GoNextLine(x, y)
			message.updateMinMax(x, y)
		}
	}
	return x, y
}

func (message *Message) updateMinMax(x int, y int) {
	// Update maxX
	if message.maxX < x {
		message.maxX = x
	}
	// Update maxY
	if message.maxY < y {
		message.maxY = y
	}
	// Update minX
	if message.minX > x {
		message.minX = x
	}
	// Update minY
	if message.minY > y {
		message.minY = y
	}
}
