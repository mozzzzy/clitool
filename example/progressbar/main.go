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

	terminal := clitool.New()

	state := new(float64)
	min := 0.0
	max := 200.0

	terminal.Progressbar("Now downloading.", min, state, max)
	for *state < max {
		time.Sleep(50 * time.Millisecond)
		*state += 1.0
	}
	time.Sleep(2 * time.Second)

	clitool.Close()
}
