package main

import (
	"github.com/nsf/termbox-go"
	"github.com/tianyueli8227/terminal-chat/graphic/anime"
	"github.com/tianyueli8227/terminal-chat/graphic/function"
	"github.com/tianyueli8227/terminal-chat/graphic/shape"
)

func handleInput() {
	input := ""
	inputStartX, inputStartY := 12, 7
	maxLines := 20
	displayLines := 10
	lines := []string{}
	currentOffset := 0

	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		shape.DrawBox(10, 5, 40, 30, termbox.ColorGreen, termbox.ColorDefault)

		function.DrawText(inputStartX, inputStartY, "Input: "+input, termbox.ColorGreen, termbox.ColorDefault)

		start := 0
		if len(lines) > maxLines {
			start = len(lines) - maxLines
		}

		visibleLines := lines[start+currentOffset:]
		if len(visibleLines) > displayLines {
			visibleLines = visibleLines[:displayLines]
		}

		for i, line := range visibleLines {
			function.DrawText(inputStartX, inputStartY+2+i, line, termbox.ColorBlue, termbox.ColorDefault)
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

func displayMenu() {
	options := []string{"Register", "Log In", "Exit"}
	selected := 0

	for {
		// time.Sleep(10 * time.Millisecond)
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		shape.DrawBox(8, 3, 65, 20, termbox.ColorGreen, termbox.ColorDefault)
		function.DrawText(14, 3, "Menu", termbox.ColorGreen, termbox.ColorDefault)
		shape.DrawBox(8, 3, 65, 30, termbox.ColorGreen, termbox.ColorDefault)
		anime.DrawTextFromFile("./graphic/anime/logo", 14, 5, termbox.ColorGreen, termbox.ColorDefault)
		function.DrawButton(14, 18+len(options)*3, 12, 3, " Register ", selected == 0)
		function.DrawButton(35, 18+len(options)*3, 12, 3, "  Log In  ", selected == 1)
		function.DrawButton(55, 18+len(options)*3, 12, 3, "   Exit   ", selected == 2)

		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			} else if ev.Key == termbox.KeyArrowLeft {
				selected = (selected - 1) % len(options)
			} else if ev.Key == termbox.KeyArrowRight {
				selected = (selected + 1) % len(options)
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
		handleInput()
	case 2:
		termbox.Close()
		return
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
