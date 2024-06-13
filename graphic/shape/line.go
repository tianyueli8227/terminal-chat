package shape

import (
	"github.com/nsf/termbox-go"
)

func DrawLine(x, y, l int, direction string, fg termbox.Attribute, bg termbox.Attribute) {
	if direction == "horizontal" {
		for i := 0; i < l; i++ {
			termbox.SetCell(x+i, y, '─', fg, bg)
		}
	} else if direction == "vertical" {
		for i := 0; i < l; i++ {
			termbox.SetCell(x, y+i, '│', fg, bg)
		}
	}
}
