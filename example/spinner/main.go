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

	// success case
	spinner1 := clitool.Spinner("Please wait a bit...")
	time.Sleep(5 * time.Second)
	spinner1.Resolve(true)

	// failure case
	spinner2 := clitool.Spinner("Please wait a bit...")
	time.Sleep(5 * time.Second)
	spinner2.Resolve(false)

	time.Sleep(2 * time.Second)

	clitool.Close()
}
