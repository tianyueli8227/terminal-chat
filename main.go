package main

import (
	"github.com/nsf/termbox-go"
	"github.com/tianyueli8227/terminal-chat/anime"
)

func handleInput() {
	input := ""
	inputStartX, inputStartY := 10, 5
	maxLines := 10
	lines := []string{}

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		anime.drawText(inputStartX, inputStartY, "Input: "+input, termbox.ColorGreen, termbox.ColorDefault)

		start := 0
		if len(lines) > maxLines {
			start = len(lines) - maxLines
		}
		for i, line := range lines[start:] {
			anime.drawText(inputStartX, inputStartY+2+i, line, termbox.ColorBlue, termbox.ColorDefault)
		}

		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			} else if ev.Key == termbox.KeyEnter {
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

		if len(lines) > maxLines {
			lines = lines[len(lines)-maxLines:]
		}

		termbox.Flush()
	}
}

func displayMenu() {
	options := []string{"Option 1", "Option 2", "Option 3"}
	selected := 0

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

		for i, option := range options {
			fg := termbox.ColorWhite
			if i == selected {
				fg = termbox.ColorRed
			}
			anime.drawText(10, 5+i, option, fg, termbox.ColorDefault)
		}

		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			} else if ev.Key == termbox.KeyArrowUp {
				if selected > 0 {
					selected--
				}
			} else if ev.Key == termbox.KeyArrowDown {
				if selected < len(options)-1 {
					selected++
				}
			} else if ev.Key == termbox.KeyEnter {
				handleOption(selected)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}

func handleOption(option int) {
	switch option {
	case 0:
		handleInput()
	case 1:
		// Add more functionality here for Option 2
		handleInput()
	case 2:
		// Add more functionality here for Option 3
		handleInput()
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	displayMenu()
}
