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
	s.body = newBody
}

func (s *snake) ChangeDirection(to direction) {
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

// func (s snake) Move(d direction) {

// 	switch d {
// 	case Left: // -1 0
// 		s.body[0][0]--
// 	case Right: // 1 0
// 		s.body[0][0]++
// 	case Up: // 0 -1
// 		s.body[0][1]--
// 	case Down: // 0 1
// 		s.body[0][1]++
// 	default:

// 	}
// }

func (s snake) Head() [2]int {
	return s.body[0]
}

func (s snake) Tail() [2]int {
	return s.body[len(s.body)-1]
}

func (s snake) Draw(b board, left, top int) {
	for i, row := range b.cells {
		for j, col := range row {
			if col == 1 {
				b.cells[i][j] = 0
			}
		}
	}

	for i, v := range s.body {
		if i == 0 {
			b.cells[v[0]][v[1]] = 3
		} else {
			b.cells[v[0]][v[1]] = 1

		}
	}

	for i, row := range b.cells {
		for j, col := range row {
			if col == 3 {
				termbox.SetCell(left+j+1, top+i+1, 'h', snakeColor, defaultColor)
			}
			if col == 1 {
				termbox.SetCell(left+j+1, top+i+1, 's', snakeColor, defaultColor)
			}
		}
	}
}
