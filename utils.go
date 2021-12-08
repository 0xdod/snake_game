package main

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func fillX(x, y, w int, cell termbox.Cell) {
	for dx := 0; dx < w; dx++ {
		termbox.SetCell(x+dx, y, cell.Ch, cell.Fg, cell.Bg)
	}
}

func printString(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}
