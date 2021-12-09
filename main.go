package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	width, height := getBoardDimension()
	g := NewGame(width, height)
	g.Start()

	fmt.Println("Snake died, game has ended")
}

func getBoardDimension() (width, height int) {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter gameboard height: ")
	sc.Scan()
	h := sc.Text()
	fmt.Println("Enter gameboard width: ")
	sc.Scan()
	w := sc.Text()
	height, _ = strconv.Atoi(h)
	width, _ = strconv.Atoi(w)

	return
}
