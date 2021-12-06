package main

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
