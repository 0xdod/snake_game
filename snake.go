package main

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
		length: 1,
	}
}

func (s snake) Move(d direction) {

	switch d {
	case Left: // -1 0
		s.body[0][0]--
	case Right: // 1 0
		s.body[0][0]++
	case Up: // 0 -1
		s.body[0][1]--
	case Down: // 0 1
		s.body[0][1]++
	default:
	}
}

func (s snake) Head() [2]int {
	return s.body[0]
}

func (s snake) Tail() [2]int {
	return s.body[len(s.body)-1]
}

func (s snake) Forward() {
	for _, row := range s.body {
		row[0] += 1
	}
}
