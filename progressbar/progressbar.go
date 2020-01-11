package progressbar

/*
 * Module Dependencies
 */

import (
	"strconv"
	"time"

	"github.com/mozzzzy/clitool/color"
	"github.com/mozzzzy/clitool/common"
)

/*
 * Types
 */

type Progressbar struct {
	messageStr                             string
	messageColorBg                         color.Color
	messageColorFg                         color.Color
	paddingMessageAndProgressbarStr        string
	paddingMessageAndProgressbarColorBg    color.Color
	paddingMessageAndProgressbarColorFg    color.Color
	progressbarPrefixStr                   string
	progressbarPrefixColorBg               color.Color
	progressbarPrefixColorFg               color.Color
	progressbarWidth                       int
	progressbarElementRune                 rune
	progressbarElementColorBg              color.Color
	progressbarElementColorFg              color.Color
	progressbarNoElementRune               rune
	progressbarNoElementColorBg            color.Color
	progressbarNoElementColorFg            color.Color
	progressbarSuffixStr                   string
	progressbarSuffixColorBg               color.Color
	progressbarSuffixColorFg               color.Color
	paddingProgressbarAndPercentageStr     string
	paddingProgressbarAndPercentageColorBg color.Color
	paddingProgressbarAndPercentageColorFg color.Color
	percentageColorBg                      color.Color
	percentageColorFg                      color.Color
	min                                    float64
	state                                  *float64
	max                                    float64
	minX                                   int
	maxX                                   int
	minY                                   int
	maxY                                   int
}

/*
 * Constants and Package Scope Variables
 */

var (
	DEFAULT_MESSAGE_COLOR_BG color.Color = color.Default
	DEFAULT_MESSAGE_COLOR_FG color.Color = color.Default

	DEFAULT_PADDING_MESSAGE_AND_PROGRESS_BAR_STR      string      = " "
	DEFAULT_PADDING_MESSAGE_AND_PROGRESS_BAR_COLOR_BG color.Color = color.Default
	DEFAULT_PADDING_MESSAGE_AND_PROGRESS_BAR_COLOR_FG color.Color = color.Default

	DEFAULT_PROGRESS_BAR_PREFIX_STR      string      = "|"
	DEFAULT_PROGRESS_BAR_PREFIX_COLOR_BG color.Color = color.Default
	DEFAULT_PROGRESS_BAR_PREFIX_COLOR_FG color.Color = color.Default

	DEFAULT_PROGRESS_BAR_WIDTH int = 50

	DEFAULT_PROGRESS_BAR_ELEMENT_RUNE     rune        = '█'
	DEFAULT_PROGRESS_BAR_ELEMENT_COLOR_BG color.Color = color.Default
	DEFAULT_PROGRESS_BAR_ELEMENT_COLOR_FG color.Color = color.Cyan

	DEFAULT_PROGRESS_BAR_NO_ELEMENT_RUNE     rune        = ' '
	DEFAULT_PROGRESS_BAR_NO_ELEMENT_COLOR_BG color.Color = color.Default
	DEFAULT_PROGRESS_BAR_NO_ELEMENT_COLOR_FG color.Color = color.Default

	DEFAULT_PROGRESS_BAR_SUFFIX_STR      string      = "|"
	DEFAULT_PROGRESS_BAR_SUFFIX_COLOR_BG color.Color = color.Default
	DEFAULT_PROGRESS_BAR_SUFFIX_COLOR_FG color.Color = color.Default

	DEFAULT_PADDING_PROGRESS_BAR_AND_PARCENTAGE_STR      string      = " "
	DEFAULT_PADDING_PROGRESS_BAR_AND_PARCENTAGE_COLOR_BG color.Color = color.Default
	DEFAULT_PADDING_PROGRESS_BAR_AND_PARCENTAGE_COLOR_FG color.Color = color.Default

	DEFAULT_PARCENTAGE_COLOR_BG color.Color = color.Default
	DEFAULT_PARCENTAGE_COLOR_FG color.Color = color.Default
)

/*
 * Functions
 */

func New(messageStr string, min float64, state *float64, max float64) *Progressbar {
	progressbar := new(Progressbar)

	progressbar.messageStr = messageStr
	progressbar.messageColorBg = DEFAULT_MESSAGE_COLOR_BG
	progressbar.messageColorFg = DEFAULT_MESSAGE_COLOR_FG

	progressbar.paddingMessageAndProgressbarStr =
		DEFAULT_PADDING_MESSAGE_AND_PROGRESS_BAR_STR
	progressbar.paddingMessageAndProgressbarColorBg =
		DEFAULT_PADDING_MESSAGE_AND_PROGRESS_BAR_COLOR_BG
	progressbar.paddingMessageAndProgressbarColorFg =
		DEFAULT_PADDING_MESSAGE_AND_PROGRESS_BAR_COLOR_FG

	progressbar.progressbarPrefixStr = DEFAULT_PROGRESS_BAR_PREFIX_STR
	progressbar.progressbarPrefixColorBg = DEFAULT_PROGRESS_BAR_PREFIX_COLOR_BG
	progressbar.progressbarPrefixColorFg = DEFAULT_PROGRESS_BAR_PREFIX_COLOR_FG

	progressbar.progressbarWidth = DEFAULT_PROGRESS_BAR_WIDTH

	progressbar.progressbarElementRune = DEFAULT_PROGRESS_BAR_ELEMENT_RUNE
	progressbar.progressbarElementColorBg = DEFAULT_PROGRESS_BAR_ELEMENT_COLOR_BG
	progressbar.progressbarElementColorFg = DEFAULT_PROGRESS_BAR_ELEMENT_COLOR_FG

	progressbar.progressbarNoElementRune = DEFAULT_PROGRESS_BAR_NO_ELEMENT_RUNE
	progressbar.progressbarNoElementColorBg = DEFAULT_PROGRESS_BAR_NO_ELEMENT_COLOR_BG
	progressbar.progressbarNoElementColorFg = DEFAULT_PROGRESS_BAR_NO_ELEMENT_COLOR_FG

	progressbar.progressbarSuffixStr = DEFAULT_PROGRESS_BAR_SUFFIX_STR
	progressbar.progressbarSuffixColorBg = DEFAULT_PROGRESS_BAR_SUFFIX_COLOR_BG
	progressbar.progressbarSuffixColorFg = DEFAULT_PROGRESS_BAR_SUFFIX_COLOR_FG

	progressbar.paddingProgressbarAndPercentageStr =
		DEFAULT_PADDING_PROGRESS_BAR_AND_PARCENTAGE_STR
	progressbar.paddingProgressbarAndPercentageColorBg =
		DEFAULT_PADDING_PROGRESS_BAR_AND_PARCENTAGE_COLOR_BG
	progressbar.paddingProgressbarAndPercentageColorFg =
		DEFAULT_PADDING_PROGRESS_BAR_AND_PARCENTAGE_COLOR_FG

	progressbar.percentageColorBg = DEFAULT_PARCENTAGE_COLOR_BG
	progressbar.percentageColorFg = DEFAULT_PARCENTAGE_COLOR_FG
	progressbar.min = min
	progressbar.state = state
	progressbar.max = max

	return progressbar
}

func (progressbar *Progressbar) GetMinX() int {
	return progressbar.minX
}

func (progressbar *Progressbar) GetMaxX() int {
	return progressbar.maxX
}

func (progressbar *Progressbar) GetMinY() int {
	return progressbar.minY
}

func (progressbar *Progressbar) GetMaxY() int {
	return progressbar.maxY
}

func (progressbar *Progressbar) SetMinX(minX int) {
	// Set minX
	progressbar.minX = minX
}

func (progressbar *Progressbar) SetMinY(minY int) {
	// Set minY
	progressbar.minY = minY
}

func (progressbar *Progressbar) Print() {
	// NOTE if this progress bar has alread printed, its maxX and maxY also has already set.
	//      But the case that its length is changed exists.
	//      So I clear old maxX and maxY here.
	//      These values are calculated again in following logic.
	progressbar.maxX = 0
	progressbar.maxY = 0

	// Print message
	x, y := common.PrintString(
		progressbar.messageStr,
		progressbar.messageColorFg, progressbar.messageColorBg,
		progressbar.minX, progressbar.minY)
	progressbar.updateMinMax(x, y)

	// Print padding between message and progress bar
	x, y = common.PrintString(
		progressbar.paddingMessageAndProgressbarStr,
		progressbar.paddingMessageAndProgressbarColorFg,
		progressbar.paddingMessageAndProgressbarColorBg,
		x, y)
	progressbar.updateMinMax(x, y)

	// Print prefix
	x, y = common.PrintString(
		progressbar.progressbarPrefixStr,
		progressbar.progressbarPrefixColorFg,
		progressbar.progressbarPrefixColorBg,
		x, y)
	progressbar.updateMinMax(x, y)

	// Print progress bar elements
	var width float64 = progressbar.max - progressbar.min
	var oneElementSize float64 = width / float64(progressbar.progressbarWidth)
	var elementLen int = int(*(progressbar.state) / oneElementSize)
	for count := 0; count < elementLen; count++ {
		x, y = common.PrintString(
			string(progressbar.progressbarElementRune),
			progressbar.progressbarElementColorFg,
			progressbar.progressbarElementColorBg,
			x, y)
		progressbar.updateMinMax(x, y)
	}
	// Print progress bar no elements
	for count := elementLen; count < progressbar.progressbarWidth; count++ {
		x, y = common.PrintString(
			string(progressbar.progressbarNoElementRune),
			progressbar.progressbarNoElementColorFg,
			progressbar.progressbarNoElementColorBg,
			x, y)
		progressbar.updateMinMax(x, y)
	}
	// Print suffix
	x, y = common.PrintString(
		progressbar.progressbarSuffixStr,
		progressbar.progressbarSuffixColorFg,
		progressbar.progressbarSuffixColorBg,
		x, y)
	progressbar.updateMinMax(x, y)

	// Print padding between suffix and percentage
	x, y = common.PrintString(
		progressbar.paddingProgressbarAndPercentageStr,
		progressbar.paddingProgressbarAndPercentageColorFg,
		progressbar.paddingProgressbarAndPercentageColorBg,
		x, y)
	progressbar.updateMinMax(x, y)

	// Print Percentage
	x, y = common.PrintString(
		strconv.Itoa(int(*(progressbar.state)/(progressbar.max-progressbar.min)*100.0))+"%",
		progressbar.percentageColorFg,
		progressbar.percentageColorBg,
		x, y)

	progressbar.updateMinMax(x, y)
}

func (progressbar *Progressbar) Run() {
	for *(progressbar.state) <= progressbar.max {
		progressbar.Print()
		time.Sleep(100 * time.Millisecond)
	}
}

func (progressbar *Progressbar) updateMinMax(x int, y int) {
	// Update maxX
	if progressbar.maxX < x {
		progressbar.maxX = x
	}
	// Update maxY
	if progressbar.maxY < y {
		progressbar.maxY = y
	}
	// Update minX
	if progressbar.minX > x {
		progressbar.minX = x
	}
	// Update minY
	if progressbar.minY > y {
		progressbar.minY = y
	}
}
