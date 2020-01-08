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

	fin := false
	terminal.Spinner("Please wait a bit...", &fin)
	time.Sleep(5 * time.Second)
	fin = true

	time.Sleep(2 * time.Second)

	clitool.Close()
}
