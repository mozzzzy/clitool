package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/message"
	"github.com/mozzzzy/clitool/password"
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

	password := password.New("Please type password.")
	answer := clitool.Inquire(password).(string)

	msg := message.New(answer)
	clitool.Print(msg)

	time.Sleep(2 * time.Second)
	clitool.Close()
}
