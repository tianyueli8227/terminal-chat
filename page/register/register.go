package register

import (
	"fmt"
	"log"
	"strings"

	"github.com/nsf/termbox-go"
	"github.com/tianyueli8227/terminal-chat/graphic/function"
	"github.com/tianyueli8227/terminal-chat/graphic/shape"
	"github.com/tianyueli8227/terminal-chat/page"
)

func (startMenu *RegisterMenu) FullFrame() {
	shape.DrawLine(15, 2, 65, "horizontal", termbox.ColorGreen, termbox.ColorDefault)
	shape.DrawLine(15, 39, 59, "horizontal", termbox.ColorGreen, termbox.ColorDefault)
	shape.DrawLine(73, 0, 48, "vertical", termbox.ColorGreen, termbox.ColorDefault)
	shape.DrawBox(15, 0, 65, 48, termbox.ColorGreen, termbox.ColorDefault)
}

type RegisterMenu struct {
	Selected      int
	SideBar       []string
	FrontBar      string
	Emoji         string
	FrontBarIndex int
}

func NewStartMenu() *RegisterMenu {
	return &RegisterMenu{
		Selected:      0,
		SideBar:       []string{"H", "I", "E"},
		FrontBar:      "Please Register With you email, and enter Refer Code",
		Emoji:         "(ʘ‿ʘ)",
		FrontBarIndex: 0,
	}
}

func (startMenu *RegisterMenu) DisplayFrontBar(text string, start int) {
	var displayText string
	var window string

	if len(text) > page.FrontBarWindowSize {
		log.Println("Front bar text is too long")
		return
	} else {
		displayText = text + strings.Repeat(" ", page.FrontBarWindowSize-len(text))
	}
	textLength := len(displayText)
	if start+page.FrontBarWindowSize <= textLength {
		window = displayText[start : start+page.FrontBarWindowSize]
	} else {
		window = displayText[start:]
		window += displayText[:page.FrontBarWindowSize-len(window)]
	}
	function.DrawText(16, 1, window, termbox.ColorGreen, termbox.ColorDefault)
}

func (startMenu *RegisterMenu) DisplayEmoji(emoji string) error {
	if len(emoji) != page.EmojiSize {
		return fmt.Errorf("invalid emoji size")
	}
	function.DrawText(74, 1, emoji, termbox.ColorGreen, termbox.ColorDefault)
	return nil
}

func (startMenu *RegisterMenu) DisplaySideBar(items []string) {
	for i, item := range items {
		function.DrawText(76, 5+3*i, item, termbox.ColorGreen, termbox.ColorBlack)
	}
}

func (startMenu *RegisterMenu) DisplayMenu() {

}

func (startMenu *RegisterMenu) Display() {
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

			} else if ev.Key == termbox.KeyArrowRight {

			} else if ev.Key == termbox.KeyEnter {
				startMenu.handleOption(startMenu.Selected)
			}
		case termbox.EventError:
			panic(ev.Err)
		}
		startMenu.FrontBarIndex = (startMenu.FrontBarIndex + 1) % page.FrontBarWindowSize
	}
}

func (startMenu *RegisterMenu) handleOption(option int) {
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
