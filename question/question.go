package question

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

type Question struct {
	Message message.Message
}

/*
 * Constants and Package Scope Variables
 */

var (
	DEFAULT_QUESTION_COLOR_BG color.Color = color.Default
	DEFAULT_QUESTION_COLOR_FG color.Color = color.Default

	DEFAULT_PREFIX_COLOR_BG color.Color = color.Default
	DEFAULT_PREFIX_COLOR_FG color.Color = color.Green
	DEFAULT_PREFIX_STR      string      = "?"

	DEFAULT_SUFFIX_COLOR_BG color.Color = color.Default
	DEFAULT_SUFFIX_COLOR_FG color.Color = color.Cyan
	DEFAULT_SUFFIX_STR      string      = ""

	DEFAULT_PADDING_QUESTION_AND_ANSWER_STR string = " "
	DEFAULT_PADDING_QUESTION_AND_PREFIX_STR string = " "
)

/*
 * Functions
 */

func New(questionStr string) *Question {
	question := new(Question)
	question.SetQuestion(DEFAULT_QUESTION_COLOR_FG, DEFAULT_QUESTION_COLOR_BG, questionStr)
	question.SetPrefix(DEFAULT_PREFIX_COLOR_FG, DEFAULT_PREFIX_COLOR_BG, DEFAULT_PREFIX_STR)
	question.SetSuffix(DEFAULT_SUFFIX_COLOR_FG, DEFAULT_SUFFIX_COLOR_BG, DEFAULT_SUFFIX_STR)
	return question
}

func (question *Question) GetMinX() int {
	return question.Message.GetMinX()
}

func (question *Question) GetMaxX() int {
	return question.Message.GetMaxX()
}

func (question *Question) GetMinY() int {
	return question.Message.GetMinY()
}

func (question *Question) GetMaxY() int {
	return question.Message.GetMaxY()
}

func (question *Question) Print() {
	question.Message.Print()
}

func (question *Question) SetQuestion(
	questionColorFg color.Color, questionColorBg color.Color, msg string) {
	question.Message.SetMessage(
		questionColorFg,
		questionColorBg,
		DEFAULT_PADDING_QUESTION_AND_PREFIX_STR+msg+DEFAULT_PADDING_QUESTION_AND_ANSWER_STR)
}

func (question *Question) SetQuestionStr(msg string) {
	question.Message.SetMessage(DEFAULT_QUESTION_COLOR_FG, DEFAULT_QUESTION_COLOR_BG, msg)
}

func (question *Question) SetQuestionColor(
	questionColorFg color.Color, questionColorBg color.Color) {
	question.Message.SetMessageColor(questionColorFg, questionColorBg)
}

func (question *Question) SetMinX(minX int) {
	question.Message.SetMinX(minX)
}

func (question *Question) SetMinY(minY int) {
	question.Message.SetMinY(minY)
}

func (question *Question) SetPrefix(
	prefixColorFg color.Color, prefixColorBg color.Color, prefix string) {
	question.Message.SetPrefix(prefixColorFg, prefixColorBg, prefix)
}

func (question *Question) SetPrefixStr(prefix string) {
	question.Message.SetPrefixStr(prefix)
}

func (question *Question) SetPrefixColor(
	prefixColorFg color.Color, prefixColorBg color.Color) {
	question.Message.SetPrefixColor(prefixColorFg, prefixColorBg)
}

func (question *Question) SetSuffix(
	suffixColorFg color.Color, suffixColorBg color.Color, suffix string) {
	question.Message.SetSuffix(suffixColorFg, suffixColorBg, suffix)
}

func (question *Question) SetSuffixStr(suffix string) {
	question.Message.SetSuffixStr(suffix)
}

func (question *Question) SetSuffixColor(
	suffixColorFg color.Color, suffixColorBg color.Color) {
	question.Message.SetSuffixColor(suffixColorFg, suffixColorBg)
}

func (question *Question) Resolve(answer string) {
	question.Message.SetSuffix(
		question.Message.SuffixColorFg,
		question.Message.SuffixColorBg,
		answer,
	)
}
