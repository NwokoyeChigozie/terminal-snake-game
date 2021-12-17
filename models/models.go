package models

type Food [2]int
type Direction int
type Action int

type Game struct {
	Board  Board
	Snake  Snake
	Food   Food
	Score  int64
	Paused bool
	Round  int64
}

type Board struct {
	Cells  [][]int
	Width  int
	Height int
}

type Snake struct {
	Body      [][2]int
	Length    int
	Direction Direction
}

type KeyPressData struct {
	Action
	Direction Direction
}
