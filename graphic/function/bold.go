package function

import (
	"github.com/nsf/termbox-go"
)

func DrawBold(x, y int, text string, fg, bg termbox.Attribute) {
	for _, ch := range text {
		termbox.SetCell(x, y, ch, fg, bg)
		termbox.SetCell(x+1, y, ch, fg, bg)
		x += 2
	}
}
