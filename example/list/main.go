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

	qStr := "What is your favorite language?"
	choices := []string{"c", "c++", "go", "java", "javascript", "php", "python"}
	answer := clitool.List(qStr, choices)

	clitool.Message("answer is " + answer)

	time.Sleep(2 * time.Second)

	clitool.Close()
}
