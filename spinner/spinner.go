package spinner

/*
 * Module Dependencies
 */

import (
	"fmt"
)

/*
 * Types
 */

type Spinner struct {
	Message string
	State   int
}

/*
 * Constants and Package Scope Variables
 */

const BRAILLE_1_4_5 rune = '⠙'
const BRAILLE_4_5_6 rune = '⠸'
const BRAILLE_5_6_3 rune = '⠴'
const BRAILLE_6_3_2 rune = '⠦'
const BRAILLE_1_2_3 rune = '⠇'
const BRAILLE_2_1_4 rune = '⠋'

/*
 * Functions
 */

func getBraille(state int) rune {
	var returnBraille rune
	switch state {
	case 0:
		returnBraille = BRAILLE_1_4_5
	case 1:
		returnBraille = BRAILLE_4_5_6
	case 2:
		returnBraille = BRAILLE_5_6_3
	case 3:
		returnBraille = BRAILLE_6_3_2
	case 4:
		returnBraille = BRAILLE_1_2_3
	case 5:
		returnBraille = BRAILLE_2_1_4
	}
	return returnBraille
}

func New(message string) *Spinner {
	spinner := new(Spinner)
	spinner.Message = message
	spinner.State = 0
	return spinner
}

func (spinner *Spinner) String() string {
	braille := getBraille(spinner.State)
	spinner.State = (spinner.State + 1) % 6
	return fmt.Sprintf("%v %v", string(braille), spinner.Message)
}
