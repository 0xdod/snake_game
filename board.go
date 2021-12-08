package main

import "github.com/nsf/termbox-go"

type board struct {
	cells  [][]int
	width  int
	height int
}

func newBoard(width, height int) board {
	cells := make([][]int, height)

	for i := range cells {
		cells[i] = make([]int, width)
	}

	return board{
		cells:  cells,
		width:  width,
		height: height,
	}
}

func (b board) Draw(left, top, bottom int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left, i, '│', defaultColor, bgColor)
		termbox.SetCell(left+b.width, i, '│', defaultColor, bgColor)
	}

	termbox.SetCell(left, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+b.width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+b.width, bottom, '┘', defaultColor, bgColor)

	fillX(left+1, top, b.width-1, termbox.Cell{Ch: '─'})
	fillX(left+1, bottom, b.width-1, termbox.Cell{Ch: '─'})
}
