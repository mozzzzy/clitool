package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/question"
)

/*
 * Types
 */

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func main() {
	clitool.Init()

	qStr := "Is this test question?"

	qDefault := clitool.Question(qStr)

	qWithColor := question.New(qStr)
	qWithColor.SetQuestionColor(color.Red, color.Green)
	clitool.Print(qWithColor)

	qWithPrefix := question.New(qStr)
	qWithPrefix.SetPrefix(color.Green, color.Default, "??? ")
	clitool.Print(qWithPrefix)

	qWithSuffix := question.New(qStr)
	qWithSuffix.SetSuffix(color.Cyan, color.Default, "(default xxx)")
	clitool.Print(qWithSuffix)

	qWithPrefixSuffix := question.New(qStr)
	qWithPrefixSuffix.SetPrefix(color.Green, color.Default, "???")
	qWithPrefixSuffix.SetSuffix(color.Cyan, color.Default, "(default xxx)")
	clitool.Print(qWithPrefixSuffix)

	time.Sleep(2 * time.Second)

	qDefault.Resolve("Yes!!")
	qWithColor.Resolve("Yes!!")
	qWithPrefix.Resolve("Yes!!")
	qWithSuffix.Resolve("Yes!!")
	qWithPrefixSuffix.Resolve("Yes!!")
	clitool.RePrintAll()

	time.Sleep(2 * time.Second)

	clitool.Close()
}
