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

	min := 0.0
	max := 200.0

	state1 := new(float64)
	clitool.Progressbar("Now downloading", min, state1, max)
	for *state1 < max {
		time.Sleep(50 * time.Millisecond)
		*state1 += 1.0
	}

	state2 := new(float64)
	clitool.Progressbar("Now downloading", min, state2, max)
	for *state2 < max {
		time.Sleep(50 * time.Millisecond)
		*state2 += 1.0
	}

	time.Sleep(2 * time.Second)

	clitool.Close()
}
