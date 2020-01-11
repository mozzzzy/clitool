package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/confirm"
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

	confirm := confirm.New("Will you marry me?")
	answer := clitool.Inquire(confirm).(bool)

	if answer {
		msg := message.New("accepted.")
		clitool.Print(msg)
	} else {
		msg := message.New("rejected.")
		clitool.Print(msg)
	}

	time.Sleep(2 * time.Second)
	clitool.Close()
}
