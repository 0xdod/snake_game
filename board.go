package main

import "github.com/nsf/termbox-go"

type board struct {
	cells  [][]rune
	width  int
	height int
}

func newBoard(width, height int) board {
	cells := make([][]rune, height)

	for i := range cells {
		cells[i] = make([]rune, width)
	}

	return board{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (b board) Draw(left, top, bottom int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '│', defaultColor, bgColor)
		termbox.SetCell(left+b.width+1, i, '│', defaultColor, bgColor)
	}

	termbox.SetCell(left-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+b.width+1, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+b.width+1, bottom, '┘', defaultColor, bgColor)

	fillX(left, top, b.width+1, termbox.Cell{Ch: '─'})
	fillX(left, bottom, b.width+1, termbox.Cell{Ch: '─'})
}
