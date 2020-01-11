package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/list"
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

	qStr := "What is your favorite language?"
	choices := []string{"c", "c++", "go", "java", "javascript", "php", "python"}
	lstDefault := list.New(qStr, choices)
	answer := clitool.Inquire(lstDefault)

	msg := message.New("answer is " + answer.(string))
	clitool.Print(msg)

	time.Sleep(2 * time.Second)

	clitool.Close()
}
