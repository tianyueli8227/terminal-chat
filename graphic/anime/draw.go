package anime

import (
	"bufio"
	"os"

	"github.com/nsf/termbox-go"
)

func drawMultiLineText(x, y int, lines []string, fg, bg termbox.Attribute) {
	for i, line := range lines {
		for j, c := range line {
			termbox.SetCell(x+j, y+i, c, fg, bg)
		}
	}
}

func DrawTextFromFile(filePath string, x, y int, fg, bg termbox.Attribute) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	drawMultiLineText(x, y, lines, fg, bg)
	termbox.Flush()
	return nil
}
