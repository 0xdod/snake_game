package main

import (
	"github.com/nsf/termbox-go"
)

type direction int

const (
	Left direction = iota + 1
	Up
	Right
	Down
)

type snake struct {
	body      [][2]int
	length    int
	direction direction
}

func newSnake() snake {
	return snake{
		body: [][2]int{
			{2, 3}, // head
			{2, 2}, // body
			{2, 1}, // tail
		},
		length:    1,
		direction: Right,
	}
}

func (s *snake) Move() {
	head := s.Head()

	switch s.direction {
	case Right:
		head[1] += 1
	case Left:
		head[1] -= 1
	case Up:
		head[0] -= 1
	case Down:
		head[0] += 1
	}

	newBody := make([][2]int, 1)
	newBody[0] = head
	newBody = append(newBody, s.body[:len(s.body)-1]...)

	if s.length+2 > len(s.body) {
		tail := s.Tail()
		tail[1] -= 1
		newBody = append(newBody, tail)
	}

	s.body = newBody
}

func (s *snake) ChangeDirection(to direction, paused bool) {
	if paused {
		return
	}

	if s.direction == to {
		return
	}

	opposites := map[direction]direction{
		Left:  Right,
		Right: Left,
		Up:    Down,
		Down:  Up,
	}

	if oppDir, ok := opposites[to]; !ok || s.direction == oppDir {
		return
	}

	s.direction = to
}

func (s snake) Head() [2]int {
	return s.body[0]
}

func (s snake) Tail() [2]int {
	return s.body[len(s.body)-1]
}

func (s *snake) Draw(b board, left, top int) {
	for i, row := range b.cells {
		for j, col := range row {
			if col == 'h' || col == 's' {
				b.cells[i][j] = '0'
			}
		}
	}

	for i, v := range s.body {
		if i == 0 {
			b.cells[v[0]][v[1]] = 'h'
		} else {
			b.cells[v[0]][v[1]] = 's'

		}
	}

	for i, row := range b.cells {
		for j, col := range row {
			if col == 'h' {
				termbox.SetCell(left+j+1, top+i+1, 's', termbox.ColorRed, defaultColor)
			} else if col == 's' {
				termbox.SetCell(left+j+1, top+i+1, 's', snakeColor, defaultColor)
			}
		}
	}
}

func (s snake) hasHitFood(f food) bool {
	h := s.Head()
	return h[0] == f[0] && h[1] == f[1]
}

func (s snake) hasHitWall(b board) bool {
	head := s.Head()

	return head[0] < 0 || head[1] < 0 || head[0] > b.height-1 || head[1] > b.width-1
}

func (s snake) hasHitSelf() bool {
	head := s.Head()

	for _, v := range s.body[1:] {
		if head[0] == v[0] && head[1] == v[1] {
			return true
		}
	}

	return false
}
