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

	answer := clitool.Input("Please type messages.")

	clitool.Message(answer)

	time.Sleep(2 * time.Second)
	clitool.Close()
}
