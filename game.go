package main

import (
	"errors"
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
	board  board
	snake  snake
	food   food
	score  int64
	paused bool
	round  int64
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
				g.snake.ChangeDirection(kp.direction, g.paused)
			case Quit:
				break Main
			case Pause:
				g.paused = !g.paused // toggle paused state
			}
		default:
			// do nothing and move on instead of blocking select
		}

		if err := g.MoveSnake(); err != nil {
			break Main
		}

		g.Render()
		time.Sleep(200 * time.Millisecond)
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
		bottom = midY + (g.board.height / 2) + 1
	)

	printString(left, top-1, termbox.ColorBlue, defaultColor, fmt.Sprintf("Score: %d", g.score))
	printString(left+15, top-1, termbox.ColorBlue, defaultColor, fmt.Sprintf("Round: %d", g.round))
	printString(left, bottom+2, termbox.ColorBlue, defaultColor, "Press SPACE to pause and ESC to quit")
	g.board.Draw(left, top, bottom)
	g.snake.Draw(g.board, left, top)
	g.food.Draw(g.board, left, top)
	termbox.Flush()
}

func (g *Game) MoveSnake() error {
	if g.paused {
		return nil
	}

	g.snake.Move()
	g.round += 1

	if g.snake.hasHitFood(g.food) {
		g.score++
		g.snake.length++
		g.RespawnFood()
	}

	if g.snake.hasHitSelf() {
		return errors.New("snake has hit self")
	}

	if g.snake.hasHitWall(g.board) {
		return errors.New("snake has hit wall")
	}

	return nil
}
