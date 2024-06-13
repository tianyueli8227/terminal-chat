package main

import (
	"github.com/nsf/termbox-go"
	// "github.com/tianyueli8227/terminal-chat/graphic/anime"
	// "github.com/tianyueli8227/terminal-chat/graphic/shape"
	"github.com/tianyueli8227/terminal-chat/page/front"
)

func displayMenu() {
	startMenu := front.NewStartMenu()
	startMenu.Display()
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	displayMenu()
}
