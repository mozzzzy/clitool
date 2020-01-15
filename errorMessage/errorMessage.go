package errorMessage

/*
 * Module Dependencies
 */

import (
	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/message"
)

/*
 * Types
 */

type ErrorMessage struct {
	Message message.Message
}

/*
 * Constants and Package Scope Variables
 */

var (
	DEFAULT_ERROR_COLOR_BG color.Color = color.Default
	DEFAULT_ERROR_COLOR_FG color.Color = color.Red

	DEFAULT_PREFIX_COLOR_BG color.Color = color.Default
	DEFAULT_PREFIX_COLOR_FG color.Color = color.Red
	DEFAULT_PREFIX_STR      string      = "✖"

	DEFAULT_SUFFIX_COLOR_BG color.Color = color.Default
	DEFAULT_SUFFIX_COLOR_FG color.Color = color.Red
	DEFAULT_SUFFIX_STR      string      = ""

	DEFAULT_PADDING_ERROR_AND_SUFFIX_STR string = " "
	DEFAULT_PADDING_ERROR_AND_PREFIX_STR string = " "
)

/*
 * Functions
 */

func New(errorStr string) *ErrorMessage {
	errorMessage := new(ErrorMessage)
	errorMessage.SetErrorMessage(DEFAULT_ERROR_COLOR_FG, DEFAULT_ERROR_COLOR_BG, errorStr)
	errorMessage.SetPrefix(DEFAULT_PREFIX_COLOR_FG, DEFAULT_PREFIX_COLOR_BG, DEFAULT_PREFIX_STR)
	errorMessage.SetSuffix(DEFAULT_SUFFIX_COLOR_FG, DEFAULT_SUFFIX_COLOR_BG, DEFAULT_SUFFIX_STR)
	return errorMessage
}

func (errorMessage *ErrorMessage) GetMinX() int {
	return errorMessage.Message.GetMinX()
}

func (errorMessage *ErrorMessage) GetMaxX() int {
	return errorMessage.Message.GetMaxX()
}

func (errorMessage *ErrorMessage) GetMinY() int {
	return errorMessage.Message.GetMinY()
}

func (errorMessage *ErrorMessage) GetMaxY() int {
	return errorMessage.Message.GetMaxY()
}

func (errorMessage *ErrorMessage) Print() {
	errorMessage.Message.Print()
}

func (errorMessage *ErrorMessage) SetErrorMessage(
	errorMessageColorFg color.Color, errorMessageColorBg color.Color, msg string) {
	errorMessage.Message.SetMessage(
		errorMessageColorFg,
		errorMessageColorBg,
		DEFAULT_PADDING_ERROR_AND_PREFIX_STR+msg+DEFAULT_PADDING_ERROR_AND_SUFFIX_STR)
}

func (errorMessage *ErrorMessage) SetErrorMessageStr(msg string) {
	errorMessage.Message.SetMessage(DEFAULT_ERROR_COLOR_FG, DEFAULT_ERROR_COLOR_BG, msg)
}

func (errorMessage *ErrorMessage) SetErrorMessageColor(
	errorMessageColorFg color.Color, errorMessageColorBg color.Color) {
	errorMessage.Message.SetMessageColor(errorMessageColorFg, errorMessageColorBg)
}

func (errorMessage *ErrorMessage) SetMinX(minX int) {
	errorMessage.Message.SetMinX(minX)
}

func (errorMessage *ErrorMessage) SetMinY(minY int) {
	errorMessage.Message.SetMinY(minY)
}

func (errorMessage *ErrorMessage) SetPrefix(
	prefixColorFg color.Color, prefixColorBg color.Color, prefix string) {
	errorMessage.Message.SetPrefix(prefixColorFg, prefixColorBg, prefix)
}

func (errorMessage *ErrorMessage) SetPrefixStr(prefix string) {
	errorMessage.Message.SetPrefixStr(prefix)
}

func (errorMessage *ErrorMessage) SetPrefixColor(
	prefixColorFg color.Color, prefixColorBg color.Color) {
	errorMessage.Message.SetPrefixColor(prefixColorFg, prefixColorBg)
}

func (errorMessage *ErrorMessage) SetSuffix(
	suffixColorFg color.Color, suffixColorBg color.Color, suffix string) {
	errorMessage.Message.SetSuffix(suffixColorFg, suffixColorBg, suffix)
}

func (errorMessage *ErrorMessage) SetSuffixStr(suffix string) {
	errorMessage.Message.SetSuffixStr(suffix)
}

func (errorMessage *ErrorMessage) SetSuffixColor(
	suffixColorFg color.Color, suffixColorBg color.Color) {
	errorMessage.Message.SetSuffixColor(suffixColorFg, suffixColorBg)
}

func (errorMessage *ErrorMessage) Resolve(answer string) {
	errorMessage.Message.SetSuffix(
		errorMessage.Message.SuffixColorFg,
		errorMessage.Message.SuffixColorBg,
		answer,
	)
}
