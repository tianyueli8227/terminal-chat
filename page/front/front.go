package front

import (
	"fmt"
	"log"
	"strings"

	"github.com/nsf/termbox-go"
	"github.com/tianyueli8227/terminal-chat/graphic/anime"
	"github.com/tianyueli8227/terminal-chat/graphic/function"
	"github.com/tianyueli8227/terminal-chat/graphic/shape"
	"github.com/tianyueli8227/terminal-chat/page"
)

func (startMenu *StartMenu) FullFrame() {
	shape.DrawLine(15, 2, 65, "horizontal", termbox.ColorGreen, termbox.ColorDefault)
	shape.DrawLine(15, 39, 59, "horizontal", termbox.ColorGreen, termbox.ColorDefault)
	shape.DrawLine(73, 0, 48, "vertical", termbox.ColorGreen, termbox.ColorDefault)
	shape.DrawBox(15, 0, 65, 48, termbox.ColorGreen, termbox.ColorDefault)
}

type StartMenu struct {
	Options       []string
	Selected      int
	SideBar       []string
	FrontBar      string
	Emoji         string
	FrontBarIndex int
}

func NewStartMenu() *StartMenu {
	return &StartMenu{
		Options:       []string{"Register", "Log In", "Exit"},
		Selected:      0,
		SideBar:       []string{"H", "I", "E"},
		FrontBar:      "Welcome to Terminal Chat! I am happy you are here!",
		Emoji:         " >_< ",
		FrontBarIndex: 0,
	}
}

func (startMenu *StartMenu) DisplayFrontBar(text string, start int) {
	var displayText string
	var window string

	if len(text) > page.FrontBarWindowSize {
		log.Println("Front bar text is too long")
		return
	} else {
		displayText = text + strings.Repeat(" ", page.FrontBarWindowSize-len(text))
	}

	// displayText = displayText + strings.Repeat(" ", 20)

	textLength := len(displayText)
	if start+page.FrontBarWindowSize <= textLength {
		window = displayText[start : start+page.FrontBarWindowSize]
	} else {
		window = displayText[start:]
		window += displayText[:page.FrontBarWindowSize-len(window)]
	}

	function.DrawText(16, 1, window, termbox.ColorGreen, termbox.ColorDefault)
}

func (startMenu *StartMenu) DisplayEmoji(emoji string) error {
	if len(emoji) != page.EmojiSize {
		return fmt.Errorf("invalid emoji size")
	}
	function.DrawText(74, 1, emoji, termbox.ColorGreen, termbox.ColorDefault)
	return nil
}

func (startMenu *StartMenu) DisplaySideBar(items []string) {
	for i, item := range items {
		function.DrawText(76, 5+3*i, item, termbox.ColorGreen, termbox.ColorBlack)
	}
}

func (startMenu *StartMenu) DisplayMenu() {
	anime.DrawTextFromFile("./graphic/anime/logo", 20, 5, termbox.ColorGreen, termbox.ColorDefault)
	function.DrawButton(20, 34+len(startMenu.Options)*3, 12, 3, " Register ", startMenu.Selected == 0)
	function.DrawButton(40, 34+len(startMenu.Options)*3, 12, 3, "  Log In  ", startMenu.Selected == 1)
	function.DrawButton(59, 34+len(startMenu.Options)*3, 12, 3, "   Exit   ", startMenu.Selected == 2)
}

func (startMenu *StartMenu) Display() {
	for {
		termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
		startMenu.FullFrame()
		startMenu.DisplayFrontBar(startMenu.FrontBar, startMenu.FrontBarIndex)
		startMenu.DisplayEmoji(startMenu.Emoji)
		startMenu.DisplaySideBar(startMenu.SideBar)
		startMenu.DisplayMenu()
		termbox.Flush()

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				return
			} else if ev.Key == termbox.KeyArrowLeft {
				startMenu.Selected = (startMenu.Selected - 1+ len(startMenu.Options)) % len(startMenu.Options)
			} else if ev.Key == termbox.KeyArrowRight {
				startMenu.Selected = (startMenu.Selected + 1) % len(startMenu.Options)
			} else if ev.Key == termbox.KeyEnter {
				startMenu.handleOption(startMenu.Selected)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		startMenu.FrontBarIndex = (startMenu.FrontBarIndex + 1) % page.FrontBarWindowSize
	}
}

func (startMenu *StartMenu) handleOption(option int) {
	switch option {
	case 0:
		function.HandleInput()
	case 1:
		function.HandleInput()
	case 2:
		termbox.Close()
		return
	}
}
