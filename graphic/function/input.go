package function

import (
	"github.com/nsf/termbox-go"
	"github.com/tianyueli8227/terminal-chat/graphic/shape"
)

func HandleInput() {
	input := ""
	inputStartX, inputStartY := 12, 7
	maxLines := 20
	displayLines := 10
	lines := []string{}
	currentOffset := 0

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		shape.DrawBox(10, 5, 40, 30, termbox.ColorGreen, termbox.ColorDefault)

		DrawText(inputStartX, inputStartY, "Input: "+input, termbox.ColorGreen, termbox.ColorDefault)

		start := 0
		if len(lines) > maxLines {
			start = len(lines) - maxLines
		}

		visibleLines := lines[start+currentOffset:]
		if len(visibleLines) > displayLines {
			visibleLines = visibleLines[:displayLines]
		}

		for i, line := range visibleLines {
			DrawText(inputStartX, inputStartY+2+i, line, termbox.ColorBlue, termbox.ColorDefault)
		}

		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			} else if ev.Key == termbox.KeyEnter {
				if len(lines) < maxLines {
					lines = append(lines, input)
				} else {
					lines = append(lines[1:], input)
				}
				input = ""
			} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			} else if ev.Key == termbox.KeyArrowUp {
				if currentOffset > 0 {
					currentOffset--
				}
			} else if ev.Key == termbox.KeyArrowDown {
				if currentOffset < len(lines)-displayLines {
					currentOffset++
				}
			} else {
				input += string(ev.Ch)
			}
		case termbox.EventError:
			panic(ev.Err)
		}

		termbox.Flush()
	}
}
