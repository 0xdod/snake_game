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

func (g Game) Start() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	defer termbox.Close()

	ch := make(chan KeyPress)
	go listenToKeyboard(ch)

	if err := g.Render(); err != nil {
		return err
	}

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
			return err
		}

		if err := g.Render(); err != nil {
			return err
		}

		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func (g Game) Render() error {
	err := termbox.Clear(defaultColor, defaultColor)

	if err != nil {
		return err
	}

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

	return termbox.Flush()
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
