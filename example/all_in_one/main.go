package main

/*
 * Module Dependencies
 */

import (
	"strings"
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

	// list
	var favoriteLanguage string = terminal.List(
		"Which language do you like?",
		[]string{"go", "javascript", "c++", "java"})
	terminal.Message(">> Chose language is " + favoriteLanguage + ".")

	// checkbox
	var languages []string = terminal.Checkbox(
		"Please choose all languages you like.",
		[]string{"go", "javascript", "c++", "java"})
	var langsStr string = strings.Join(languages, ", ")
	terminal.Message(">> Chose languages are " + langsStr + ".")

	// confirm
	var yesNo bool = terminal.Confirm("Do you know golang?")
	if yesNo {
		terminal.Message(">> yes")
	} else {
		terminal.Message(">> no")
	}

	// input
	var name string = terminal.Input("What is your name.")
	terminal.Message(">> " + name)

	// password
	var pass string = terminal.Password("password:")
	terminal.Message(">> " + pass)

	// spinner
	var fin bool = false
	terminal.Spinner("Please wait a bit...", &fin)
	time.Sleep(5 * time.Second)
	fin = true

	// progress bar
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
