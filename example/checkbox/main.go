package main

/*
 * Module Dependencies
 */

import (
	"strings"
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

	qStr := "Please select all languages you like."
	choices := []string{"c", "c++", "go", "java", "javascript", "php", "python"}
	answers := clitool.Checkbox(qStr, choices)

	msgStr := "answers are (" + strings.Join(answers, ",") + ")"
	clitool.Message(msgStr)

	time.Sleep(2 * time.Second)

	clitool.Close()
}
