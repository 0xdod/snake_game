package main

import (
	"github.com/nsf/termbox-go"
)

type Action int

const (
	Move Action = iota + 1
	Quit
	Restart
)

type KeyPress struct {
	Action
	direction direction
}

func listenToKeyboard(ch chan KeyPress) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				ch <- KeyPress{Action: Move, direction: Left}
			case termbox.KeyArrowDown:
				ch <- KeyPress{Action: Move, direction: Down}
			case termbox.KeyArrowRight:
				ch <- KeyPress{Action: Move, direction: Right}
			case termbox.KeyArrowUp:
				ch <- KeyPress{Action: Move, direction: Up}
			case termbox.KeyEsc:
				ch <- KeyPress{Action: Quit}
			default:
				if ev.Ch == 'r' {
					ch <- KeyPress{Action: Restart}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
