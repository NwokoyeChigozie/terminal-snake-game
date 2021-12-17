package controllers

import "github.com/gregoflash05/terminal-snake-game/models"

const (
	Left models.Direction = iota + 1
	Up
	Right
	Down
)

func GetnewSnake() (snake models.Snake) {
	snake = models.Snake{
		Body: [][2]int{
			{2, 3}, // head
			{2, 2}, // body
			{2, 1}, // tail
		},
		Length:    1,
		Direction: Right,
	}

	return
}
