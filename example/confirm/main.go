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

	answer := clitool.Confirm("Will you marry me?")
	if answer {
		clitool.Message("Accepted.")
	} else {
		clitool.Message("Rejected.")
	}

	time.Sleep(2 * time.Second)
	clitool.Close()
}
