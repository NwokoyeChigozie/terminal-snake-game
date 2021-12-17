package controllers

import (
	"github.com/gregoflash05/terminal-snake-game/models"
	"github.com/nsf/termbox-go"
)

const (
	Move models.Action = iota + 1
	Quit
	Restart
	Pause
)

func ListenToKeyboardAction(ch chan models.KeyPressData) {
	termbox.SetInputMode(termbox.InputEsc)

	for {
		//nolint:exhaustive //Greg: skip
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyArrowLeft:
				ch <- models.KeyPressData{Action: Move, Direction: Left}
			case termbox.KeyArrowDown:
				ch <- models.KeyPressData{Action: Move, Direction: Down}
			case termbox.KeyArrowRight:
				ch <- models.KeyPressData{Action: Move, Direction: Right}
			case termbox.KeyArrowUp:
				ch <- models.KeyPressData{Action: Move, Direction: Up}
			case termbox.KeyEsc:
				ch <- models.KeyPressData{Action: Quit}
			case termbox.KeySpace:
				ch <- models.KeyPressData{Action: Pause}
			default:
				if ev.Ch == 'r' {
					ch <- models.KeyPressData{Action: Restart}
				}
			}
		case termbox.EventError:
			panic(ev.Err)
		}
	}
}
