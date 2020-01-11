package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/input"
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

	input := input.New("Please type messages.")
	answer := clitool.Inquire(input).(string)

	msg := message.New(answer)
	clitool.Print(msg)

	time.Sleep(2 * time.Second)
	clitool.Close()
}
