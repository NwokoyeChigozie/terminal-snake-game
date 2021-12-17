package controllers

import (
	"log"
	"time"

	"github.com/gregoflash05/terminal-snake-game/models"
	"github.com/nsf/termbox-go"
)

const (
	defaultColor = termbox.ColorDefault
	bgColor      = termbox.ColorDefault
	snakeColor   = termbox.ColorGreen
)

var (
	TwoHundred time.Duration = 200
)

func GenerateGameInfo(boardHeight, boardWidth int) (game models.Game) {
	game.Board = GetnewBoard(boardHeight, boardWidth)
	game.Snake = GetnewSnake()
	game.Food = models.Food{5, 5}

	return
}

func StartSnakeGame() {
	height, width := RequestBoardInfo()
	snakeGame := GenerateGameInfo(height, width)

	if err := Start(&snakeGame); err != nil {
		log.Fatal(err)
	}
}

func Start(game *models.Game) error {
	if err := termbox.Init(); err != nil {
		return err
	}

	defer termbox.Close()

	keyboardChannel := make(chan models.KeyPressData)
	go ListenToKeyboardAction(keyboardChannel)

	if err := Render(game); err != nil {
		return err
	}

Main:
	for {
		select {
		case kp := <-keyboardChannel:
			switch kp.Action {
			case Move:
				game.Snake = ChangeDirection(&game.Snake, kp.Direction, game.Paused)
			case Quit:
				break Main
			case Pause:
				game.Paused = !game.Paused // toggle paused state
			}
		default:
			// do nothing and move on instead of blocking select
		}
		var err error

		if game, err = MoveSnake(game); err != nil {
			return err
		}

		if err := Render(game); err != nil {
			return err
		}

		time.Sleep(TwoHundred * time.Millisecond)
	}

	return nil
}
