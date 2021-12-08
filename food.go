package main

import (
	"math/rand"
	"time"

	"github.com/nsf/termbox-go"
)

type food [2]int

func (g *Game) RespawnFood() {
	g.board.cells[g.food[0]][g.food[1]] = '0'

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	row := r.Intn(g.board.height - 1)
	col := r.Intn(g.board.width - 1)

	if g.board.cells[row][col] == 's' {
		g.RespawnFood()
		return
	}

	g.food = food{row, col}
}

func (f food) Draw(b board, left, top int) {
	b.cells[f[0]][f[1]] = 'f'

	for i, row := range b.cells {
		for j, col := range row {
			if col == 'f' {
				termbox.SetCell(left+j+1, top+i+1, 'o', defaultColor, defaultColor)
			}
		}
	}
}
