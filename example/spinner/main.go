package main

/*
 * Module Dependencies
 */

import (
	"time"

	"github.com/mozzzzy/clitool"
	"github.com/mozzzzy/clitool/spinner"
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
	spinner1 := spinner.New("Please wait a bit...")
	clitool.Run(spinner1)
	time.Sleep(5 * time.Second)
	spinner1.Resolve(true)

	// failure case
	spinner2 := spinner.New("Please wait a bit...")
	clitool.Run(spinner2)
	time.Sleep(5 * time.Second)
	spinner2.Resolve(false)

	time.Sleep(2 * time.Second)

	clitool.Close()
}
