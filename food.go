package main

import "github.com/nsf/termbox-go"

type food [2]int

func (f food) SetPosition(x, y int) {

}

func (f food) Draw(b board, left, top int) {
	b.cells[f[0]][f[1]] = 2
	for i, row := range b.cells {
		for j, col := range row {
			if col == 2 {
				termbox.SetCell(left+j+1, top+i+1, 'o', defaultColor, defaultColor)
			}
		}
	}
}
