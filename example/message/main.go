package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/message"
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

	msgStr := "This is test message."

	msgDefault := message.New(msgStr)
	clitool.Print(msgDefault)

	msgWithColor := message.New(msgStr)
	msgWithColor.SetMessageColor(color.Red, color.Green)
	clitool.Print(msgWithColor)

	msgWithPrefix := message.New(msgStr)
	msgWithPrefix.SetPrefix(color.Green, color.Default, "[prefix]")
	clitool.Print(msgWithPrefix)

	msgWithSuffix := message.New(msgStr)
	msgWithSuffix.SetSuffix(color.Green, color.Default, "[suffix]")
	clitool.Print(msgWithSuffix)

	msgWithPrefixSuffix := message.New(msgStr)
	msgWithPrefixSuffix.SetPrefix(color.Green, color.Default, "[prefix]")
	msgWithPrefixSuffix.SetSuffix(color.Green, color.Default, "[suffix]")
	clitool.Print(msgWithPrefixSuffix)

	time.Sleep(2 * time.Second)

	msgDefault.SetSuffix(color.Cyan, color.Default, "[This suffix is reprinted]")
	clitool.RePrintAll()

	time.Sleep(2 * time.Second)

	clitool.Close()
}
