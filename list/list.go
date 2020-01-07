package list

/*
 * Module Dependencies
 */

import (
	"fmt"
)

/*
 * Types
 */

type List struct {
	Choices        []string
	CursorPosition int
}

/*
 * Constants and Package Scope Variables
 */

/*
 * Functions
 */

func New(choices []string) *List {
	list := new(List)
	list.Choices = choices
	list.CursorPosition = 0

	return list
}

func (list *List) String() string {
	var listStr string
	for index, choice := range list.Choices {
		if index == list.CursorPosition {
			listStr += fmt.Sprintf("❯ ")
		} else {
			listStr += fmt.Sprintf("  ")
		}
		listStr += fmt.Sprintf("%v", choice)
		if index != len(list.Choices)-1 {
			listStr += "\n"
		}
	}
	return listStr
}

func (list *List) AddChoice(choice string) {
	list.Choices = append(list.Choices, choice)
}

func (list *List) RemoveChoice(removedChoice string) {
	var newChoices []string
	for _, choice := range list.Choices {
		if choice == removedChoice {
			continue
		}
		newChoices = append(newChoices, choice)
	}
	list.Choices = newChoices
}
