package main

import (
	"os"

	"github.com/nsf/termbox-go"
)

type KeyPress struct {
	Message string
}

func listenToKeyboard(ch chan KeyPress) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:

			case termbox.KeyArrowDown:

			case termbox.KeyArrowRight:
				ch <- KeyPress{"Right arrow"}
			case termbox.KeyArrowUp:

			case termbox.KeyEsc:

			default:
				if ev.Ch == 'q' {
					os.Exit(0)
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
