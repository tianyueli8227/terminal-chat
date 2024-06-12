package function

import "github.com/nsf/termbox-go"

func DrawText(x, y int, text string, fg, bg termbox.Attribute) {
	for i, c := range text {
		termbox.SetCell(x+i, y, c, fg, bg)
	}
}
