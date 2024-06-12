package function

import (
	"github.com/nsf/termbox-go"
	"github.com/tianyueli8227/terminal-chat/graphic/shape"
)

func DrawButton(x, y int, width int, height int, text string, selected bool) {
	fg := termbox.ColorGreen
	bg := termbox.ColorBlack
	if selected {
		fg = termbox.ColorBlack
		bg = termbox.ColorLightGreen
	}
	if width < len(text)+2 {
		width = len(text) + 2
	}
	shape.DrawBox(x-1, y-1, width, height, termbox.ColorGreen, termbox.ColorBlack)
	DrawText(x, y, text, fg, bg)
}
