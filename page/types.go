package page

var FrontBarWindowSize = 57
var EmojiSize = 5
type Page interface {
	DisplayFrontBar(text string, start int)
	DisplayEmoji(emoji string) error
	DisplaySideBar()
	DisplayMenu()
	Display()
}
