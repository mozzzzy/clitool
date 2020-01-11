package main

/*
 * Module Dependencies
 */

import (
	"strings"
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/checkbox"
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

	qStr := "Please select all languages you like."
	choices := []string{"c", "c++", "go", "java", "javascript", "php", "python"}
	chkboxDefault := checkbox.New(qStr, choices)
	answers := clitool.Inquire(chkboxDefault)

	msgStr := "answers are (" + strings.Join(answers.([]string), ",") + ")"
	msg := message.New(msgStr)
	clitool.Print(msg)

	time.Sleep(2 * time.Second)

	clitool.Close()
}
