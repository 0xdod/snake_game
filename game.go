package main

import (
	"fmt"
	"time"

	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
)

type Game struct {
	board board
	snake snake
	food  food
	score int64
}

func NewGame(width, height int) Game {
	g := Game{}
	g.board = newBoard(width, height)
	g.snake = newSnake()
	g.food = food{5, 6}

	return g
}

func (g Game) Start() {
	if err := termbox.Init(); err != nil {
		panic(err)
	}

	defer termbox.Close()

	ch := make(chan KeyPress)
	go listenToKeyboard(ch)

	g.Render()

Main:
	for {
		select {
		case kp := <-ch:
			switch kp.Action {
			case Move:
				g.snake.ChangeDirection(kp.direction)
			case Quit:
				break Main
			}
		default:
			// do nothing and move on instead of waiting
		}
		g.snake.Move()
		g.Render()
		time.Sleep(500 * time.Millisecond)
	}
}

func (g Game) Render() {
	termbox.Clear(defaultColor, defaultColor)

	var (
		w, h = termbox.Size()
		midY = h / 2
		midX = w / 2
		left = midX - (g.board.width / 2)
		//right = midX + (g.board.width / 2)
		top    = midY - (g.board.height / 2)
		bottom = midY + (g.board.height / 2)
	)

	printString(left, top-1, termbox.ColorBlue, defaultColor, fmt.Sprintf("Score: %d", g.score))
	g.board.Draw(left, top, bottom)
	g.snake.Draw(g.board, left, top)
	g.food.Draw(g.board, left, top)
	termbox.Flush()
}
