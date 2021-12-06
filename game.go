package main

import (
	"fmt"

	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
)

// what is in the array doesn't matter for now all i need is an index to locate position
// position is located by boards[row][col] where row and col are slice indices.
type Game struct {
	board board
	snake snake
	score int64
}

func NewGame(width, height int) Game {
	g := Game{}
	g.board = newBoard(width, height)
	g.snake = newSnake()

	return g
}

func (g Game) DrawBoard(top, bottom, left int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left, i, '│', defaultColor, bgColor)
		termbox.SetCell(left+g.board.width, i, '│', defaultColor, bgColor)
	}

	termbox.SetCell(left, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+g.board.width, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+g.board.width, bottom, '┘', defaultColor, bgColor)

	fillX(left+1, top, g.board.width-1, termbox.Cell{Ch: '─'})
	fillX(left+1, bottom, g.board.width-1, termbox.Cell{Ch: '─'})
}

func (g Game) DrawSnake() {
	g.board.cells[g.snake.Head()[0]][g.snake.Head()[1]] = 1
	g.board.cells[g.snake.body[1][0]][g.snake.body[1][1]] = 1
	g.board.cells[g.snake.Tail()[0]][g.snake.Tail()[1]] = 1
}

func (g Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer termbox.Close()

	ch := make(chan KeyPress)
	go listenToKeyboard(ch)

	g.Render()

	// for {
	// 	select {
	// 	case evt := <-ch:
	// 		for i, r := range evt.Message {
	// 			termbox.SetCell(0, i, r, defaultColor, bgColor)
	// 		}

	// 	}
	// 	g.Render()
	// }
	for {
	}
}

func (g Game) Render() {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h = termbox.Size()
		midY = h / 2
		midX = w / 2
		left = midX - (g.board.width / 2)
		//right  = midX + (g.board.width / 2)
		top    = midY - (g.board.height / 2)
		bottom = midY + (g.board.height / 2)
	)

	printString(left, top-2, termbox.ColorBlue, defaultColor, fmt.Sprintf("Score: %d", g.score))
	// g.DrawSnake()
	g.DrawBoard(top, bottom, left)
	// fmt.Println(w, h, midY, left, right, top, bottom)
	termbox.Flush()
}

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
