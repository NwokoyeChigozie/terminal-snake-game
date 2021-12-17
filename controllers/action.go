package controllers

import (
	"errors"
	"math/rand"
	"time"

	"github.com/gregoflash05/terminal-snake-game/models"
)

func ChangeDirection(s *models.Snake, to models.Direction, paused bool) models.Snake {
	if paused {
		ts := *s
		return ts
	}

	if s.Direction == to {
		ts := *s
		return ts
	}

	opposites := map[models.Direction]models.Direction{
		Left:  Right,
		Right: Left,
		Up:    Down,
		Down:  Up,
	}

	if oppDir, ok := opposites[to]; !ok || s.Direction == oppDir {
		ts := *s
		return ts
	}

	s.Direction = to

	return *s
}

func MoveSnake(game *models.Game) (*models.Game, error) {
	if game.Paused {
		return game, nil
	}

	MoveS(&game.Snake)
	game.Round++

	if hasHitFood(game.Snake, game.Food) {
		game.Score++
		game.Snake.Length++
		RespawnFood(game)
	}

	if hasHitSelf(game.Snake) {
		return game, errors.New("snake has hit self")
	}

	if hasHitWall(game.Snake, game.Board) {
		return game, errors.New("snake has hit wall")
	}

	return game, nil
}

func RespawnFood(g *models.Game) *models.Game {
	g.Board.Cells[g.Food[0]][g.Food[1]] = '0'

	//nolint:gosec //Greg: skip
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	row := r.Intn(g.Board.Height - 1)
	col := r.Intn(g.Board.Width - 1)

	if g.Board.Cells[row][col] == 's' {
		RespawnFood(g)
		return g
	}

	g.Food = models.Food{row, col}

	return g
}

func HeadS(s models.Snake) [2]int {
	return s.Body[0]
}
func TailS(s models.Snake) [2]int {
	return s.Body[len(s.Body)-1]
}

func MoveS(s *models.Snake) *models.Snake {
	head := HeadS(*s)

	switch s.Direction {
	case Right:
		head[1]++
	case Left:
		head[1]--
	case Up:
		head[0]--
	case Down:
		head[0]++
	}

	newBody := make([][2]int, 1)
	newBody[0] = head
	newBody = append(newBody, s.Body[:len(s.Body)-1]...)

	if s.Length+2 > len(s.Body) {
		tail := TailS(*s)
		tail[1]--
		newBody = append(newBody, tail)
	}

	s.Body = newBody

	return s
}

func hasHitFood(s models.Snake, f models.Food) bool {
	h := HeadS(s)
	return h[0] == f[0] && h[1] == f[1]
}

func hasHitWall(s models.Snake, b models.Board) bool {
	head := HeadS(s)

	//nolint:gocritic //Greg: skip
	return head[0] < 0 || head[1] < 0 || head[0] > b.Height-1 || head[1] > b.Width-1
}

func hasHitSelf(s models.Snake) bool {
	head := HeadS(s)

	for _, v := range s.Body[1:] {
		if head[0] == v[0] && head[1] == v[1] {
			return true
		}
	}

	return false
}
