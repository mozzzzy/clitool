package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/errorMessage"
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

	eStr := "Is this test error"

	eDefault := errorMessage.New(eStr)
	clitool.Print(eDefault)

	time.Sleep(2 * time.Second)

	clitool.Close()
}
