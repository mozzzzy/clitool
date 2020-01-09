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
	message.PrefixStr = "! "
	message.PrefixColorFg = color.Cyan

	message.SuffixStr = " !"
	message.SuffixColorFg = color.Cyan

	clitool.Init()

	terminal := clitool.New()

	terminal.Message("This is test message.")
	time.Sleep(2 * time.Second)

	clitool.Close()
}
