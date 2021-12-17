package controllers

import (
	"fmt"

	"github.com/gregoflash05/terminal-snake-game/models"
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

var (
	Fifiteen = 15
	Two      = 2
)

func fillX(x, y, w int, cell termbox.Cell) {
	for dx := 0; dx < w; dx++ {
		termbox.SetCell(x+dx, y, cell.Ch, cell.Fg, cell.Bg)
	}
}

func printString(x, y int, fg, bg termbox.Attribute, msg string) {
	for _, c := range msg {
		termbox.SetCell(x, y, c, fg, bg)
		x += runewidth.RuneWidth(c)
	}
}

func Render(game *models.Game) error {
	err := termbox.Clear(defaultColor, defaultColor)

	if err != nil {
		return err
	}

	var (
		w, h = termbox.Size()
		midY = h / 2
		midX = w / 2
		left = midX - (game.Board.Width / 2)
		//nolint:gocritic //Greg: skip
		// right  = midX + (game.Board.Width / 2)
		top    = midY - (game.Board.Height / 2)
		bottom = midY + (game.Board.Height / 2) + 1
	)

	printString(left, top-1, termbox.ColorBlue, defaultColor, fmt.Sprintf("Score: %d", game.Score))
	printString(left+Fifiteen, top-1, termbox.ColorBlue, defaultColor, fmt.Sprintf("Round: %d", game.Round))
	printString(left, bottom+Two, termbox.ColorBlue, defaultColor, "Press SPACE to pause and ESC to quit")
	DrawBoard(game.Board, left, top, bottom)
	DrawSnake(&game.Snake, game.Board, left, top)
	DrawFood(game.Food, game.Board, left, top)

	return termbox.Flush()
}

func DrawBoard(b models.Board, left, top, bottom int) {
	for i := top; i < bottom; i++ {
		termbox.SetCell(left-1, i, '│', defaultColor, bgColor)
		termbox.SetCell(left+b.Width+1, i, '│', defaultColor, bgColor)
	}

	termbox.SetCell(left-1, top, '┌', defaultColor, bgColor)
	termbox.SetCell(left-1, bottom, '└', defaultColor, bgColor)
	termbox.SetCell(left+b.Width+1, top, '┐', defaultColor, bgColor)
	termbox.SetCell(left+b.Width+1, bottom, '┘', defaultColor, bgColor)

	fillX(left, top, b.Width+1, termbox.Cell{Ch: '─'})
	fillX(left, bottom, b.Width+1, termbox.Cell{Ch: '─'})
}

func DrawSnake(s *models.Snake, b models.Board, left, top int) {
	for i, row := range b.Cells {
		for j, col := range row {
			if col == 'h' || col == 's' {
				b.Cells[i][j] = '0'
			}
		}
	}

	for i, v := range s.Body {
		if i == 0 {
			b.Cells[v[0]][v[1]] = 'h'
		} else {
			b.Cells[v[0]][v[1]] = 's'
		}
	}

	for i, row := range b.Cells {
		for j, col := range row {
			if col == 'h' {
				termbox.SetCell(left+j+1, top+i+1, '=', termbox.ColorRed, defaultColor)
			} else if col == 's' {
				termbox.SetCell(left+j+1, top+i+1, '=', snakeColor, defaultColor)
			}
		}
	}
}

func DrawFood(f models.Food, b models.Board, left, top int) {
	b.Cells[f[0]][f[1]] = 'f'

	for i, row := range b.Cells {
		for j, col := range row {
			if col == 'f' {
				termbox.SetCell(left+j+1, top+i+1, 'o', defaultColor, defaultColor)
			}
		}
	}
}
