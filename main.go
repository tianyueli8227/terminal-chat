package main

import (
	"github.com/nsf/termbox-go"
)

func drawText(x, y int, text string, fg, bg termbox.Attribute) {
	for i, c := range text {
		termbox.SetCell(x+i, y, c, fg, bg)
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.Flush()

	input := ""
	inputStartX, inputStartY := 10, 5
	maxLines := 10
	lines := []string{}

	for {
		// Draw input box and the current input text
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		drawText(inputStartX, inputStartY, "Input: "+input, termbox.ColorGreen, termbox.ColorDefault)

		// Draw the last maxLines lines of output
		start := 0
		if len(lines) > maxLines {
			start = len(lines) - maxLines
		}
		for i, line := range lines[start:] {
			drawText(inputStartX, inputStartY+2+i, line, termbox.ColorBlue, termbox.ColorDefault)
		}

		termbox.Flush()

		// Wait for an event
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				// Exit on ESC key
				return
			} else if ev.Key == termbox.KeyEnter {
				// Add input text to lines and clear input
				lines = append(lines, input)
				input = ""
			} else if ev.Key == termbox.KeyBackspace || ev.Key == termbox.KeyBackspace2 {
				if len(input) > 0 {
					input = input[:len(input)-1]
				}
			} else {
				input += string(ev.Ch)
			}
		case termbox.EventError:
			panic(ev.Err)
		}

		// Remove old lines if necessary
		if len(lines) > maxLines {
			lines = lines[len(lines)-maxLines:]
		}

		// Make sure the terminal is updated
		termbox.Flush()
	}
}
