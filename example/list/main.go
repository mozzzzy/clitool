package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
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

	terminal := clitool.New()

	terminal.List("Which language do you like?", []string{"go", "javascript", "c++", "java"})
	time.Sleep(2 * time.Second)

	clitool.Close()
}
