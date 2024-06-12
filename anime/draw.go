package dra

import (
	"github.com/nsf/termbox-go"
)

func drawText(x, y int, text string, fg, bg termbox.Attribute) {
	for i, c := range text {
		termbox.SetCell(x+i, y, c, fg, bg)
	}
}
