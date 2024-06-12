package shape

import (
	"github.com/nsf/termbox-go"
)

func DrawBox(x, y, w, h int,fg termbox.Attribute, bg termbox.Attribute) {
	horizontal := '─'
	vertical := '│'
	topLeft := '┌'
	topRight := '┐'
	bottomLeft := '└'
	bottomRight := '┘'

	termbox.SetCell(x, y, topLeft, fg, bg)
	termbox.SetCell(x+w-1, y, topRight, fg, bg)
	termbox.SetCell(x, y+h-1, bottomLeft, fg, bg)
	termbox.SetCell(x+w-1, y+h-1, bottomRight, fg, bg)

	for i := 1; i < w-1; i++ {
		termbox.SetCell(x+i, y, horizontal, fg, bg)
		termbox.SetCell(x+i, y+h-1, horizontal, fg, bg)
	}

	for i := 1; i < h-1; i++ {
		termbox.SetCell(x, y+i, vertical, fg, bg)
		termbox.SetCell(x+w-1, y+i, vertical, fg, bg)
	}
}
