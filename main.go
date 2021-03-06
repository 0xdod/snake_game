package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	width, height := getBoardDimension()
	g := NewGame(width, height)

	if err := g.Start(); err != nil {
		log.Fatal(err)
	}
}

func getBoardDimension() (width, height int) {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter gameboard height: ")
	sc.Scan()
	h := sc.Text()
	fmt.Println("Enter gameboard width: ")
	sc.Scan()
	w := sc.Text()

	var err error
	height, err = strconv.Atoi(h)

	if err != nil {
		height = 12 // default value
	}

	width, err = strconv.Atoi(w)

	if err != nil {
		width = 80
	}

	return
}
